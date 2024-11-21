package datacollector

import (
	"main/models"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type DataCollectorTestSuite struct {
	suite.Suite
	DataCollector *gin.Engine
}

func TestDataCollectorTestSuite(t *testing.T) {
	suite.Run(t, new(DataCollectorTestSuite))
}

func (s *DataCollectorTestSuite) SetupSuite() {
}

func (s *DataCollectorTestSuite) TearDownSuite() {
}

func (s *DataCollectorTestSuite) TestGenerateFile() {
	res, err := CollectDatum("http://192.168.4.94/humidity")
	assert.Nil(s.T(), err)
	assert.NotEqual(s.T(), "", res)
	res, err = CollectDatum("http://192.168.4.94/temperature")
	assert.Nil(s.T(), err)
	assert.NotEqual(s.T(), "", res)
}

func (s *DataCollectorTestSuite) TestCollectAllData() {
	probes := []models.Probe{
		{
			ProbeName: "inside",
			ProbeIP:   "192.168.1.2",
		},
		{
			ProbeName: "outside",
			ProbeIP:   "192.168.1.3",
		},
	}
	collectedProbeReports, err := CollectAllData(probes, time.Date(2006, 11, 11, 11, 11, 1, 1, time.UTC), MockDatumCollector)
	assert.Nil(s.T(), err)
	expectedCollectedData := models.CollectedProbeReports{
		CollectedTime: time.Date(2006, time.November, 11, 11, 11, 1, 1, time.UTC),
		ProbeReports: []models.ProbeReport{
			{
				Probe:       probes[0],
				Temperature: "20.5",
				Humidity:    "20.5",
			},
			{
				Probe:       probes[1],
				Temperature: "20.5",
				Humidity:    "20.5",
			},
		},
	}
	assert.Equal(s.T(), expectedCollectedData, collectedProbeReports)
}
