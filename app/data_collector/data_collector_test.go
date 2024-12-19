package datacollector

import (
	"flag"
	"main/models"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

var integration *bool = flag.Bool("integration", false, "run only integration tests")

type DataCollectorTestSuite struct {
	suite.Suite
	DataCollector *gin.Engine
}

func TestDataCollectorTestSuite(t *testing.T) {
	if !*integration {
		t.Skip("integration test only")
	}
	suite.Run(t, new(DataCollectorTestSuite))
}

func (s *DataCollectorTestSuite) TestGenerateFile() {
	res, err := CollectDatum("http://192.168.4.94/humidity")
	assert.Nil(s.T(), err)
	assert.NotEqual(s.T(), 0, res)
	res, err = CollectDatum("http://192.168.4.94/temperature")
	assert.Nil(s.T(), err)
	assert.NotEqual(s.T(), 0, res)
}

func (s *DataCollectorTestSuite) TestCollectAllData() {
	probes := []models.Probe{
		{
			ProbeName:     "inside",
			ProbeaAddress: "http://192.168.4.1",
		},
		{
			ProbeName:     "outside",
			ProbeaAddress: "http://192.168.4.1",
		},
	}
	collectedProbeReports, err := CollectAllData(probes, time.Date(2006, 11, 11, 11, 11, 1, 1, time.UTC), MockDatumCollector)
	assert.Nil(s.T(), err)
	expectedCollectedData := models.CollectedProbeReports{
		CollectedTime: time.Date(2006, time.November, 11, 11, 11, 1, 1, time.UTC),
		ProbeReports: []models.ProbeReport{
			{
				Probe:       probes[0],
				Temperature: 20.5,
				Humidity:    20.5,
			},
			{
				Probe:       probes[1],
				Temperature: 20.5,
				Humidity:    20.5,
			},
		},
	}
	assert.Equal(s.T(), expectedCollectedData, collectedProbeReports)
}
