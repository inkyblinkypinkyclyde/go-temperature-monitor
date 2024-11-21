package report

import (
	"main/models"
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ReportTestSuite struct {
	suite.Suite
	Report *gin.Engine
}

const testFileName = "testfile.xlsx"

var collectedProbeReports = models.CollectedProbeReports{
	CollectedTime: time.Date(2006, 11, 11, 11, 11, 1, 1, time.UTC),
	ProbeReports: []models.ProbeReport{
		{
			Probe: models.Probe{
				ProbeName:     "inside",
				ProbeaAddress: "192.168.1.2",
			},
			Temperature: 20,
			Humidity:    50,
		},
		{
			Probe: models.Probe{
				ProbeName:     "outside",
				ProbeaAddress: "192.168.1.3",
			},
			Temperature: 10,
			Humidity:    40,
		},
	},
}

var probes = []models.Probe{
	{
		ProbeName:     "inside",
		ProbeaAddress: "192.168.1.2",
	},
	{
		ProbeName:     "outside",
		ProbeaAddress: "192.168.1.3",
	},
}

func TestReportTestSuite(t *testing.T) {
	suite.Run(t, new(ReportTestSuite))
}

func (s *ReportTestSuite) TearDownSuite() {
	os.Remove(testFileName)
}

func (s *ReportTestSuite) TestGenerateFile() {
	err := GenerateEmptyReport(testFileName, probes)
	assert.Equal(s.T(), nil, err)

	_, err = os.Stat(testFileName)
	assert.Equal(s.T(), nil, err)
	rows, _ := getReport(testFileName)
	assert.Equal(s.T(), "time", rows[0][0])
	assert.Equal(s.T(), "inside temperature", rows[0][1])
	assert.Equal(s.T(), "inside humidity", rows[0][2])
	assert.Equal(s.T(), "outside temperature", rows[0][3])
	assert.Equal(s.T(), "outside humidity", rows[0][4])
}

func (s *ReportTestSuite) TestLogReading() {
	GenerateEmptyReport(testFileName, probes)

	err := LogCollectedProbeReports(collectedProbeReports, 2, testFileName)

	assert.Nil(s.T(), err)

	rows, _ := getReport(testFileName)
	assert.Equal(s.T(), "11/11/06 11:11", rows[1][0])
	assert.Equal(s.T(), "20", rows[1][1])
	assert.Equal(s.T(), "50", rows[1][2])
	assert.Equal(s.T(), "10", rows[1][3])
	assert.Equal(s.T(), "40", rows[1][4])
}

func (s *ReportTestSuite) TestFindNextEmptyRow() {
	GenerateEmptyReport(testFileName, probes)
	LogCollectedProbeReports(collectedProbeReports, 2, testFileName)

	nextEmptyRow, err := GetNextEmptyRow(testFileName)
	assert.Nil(s.T(), err)

	assert.Equal(s.T(), 3, nextEmptyRow)
}
