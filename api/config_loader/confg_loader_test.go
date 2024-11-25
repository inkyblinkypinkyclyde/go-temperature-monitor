package configloader

import (
	"main/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ConfigLoaderTestSuite struct {
	suite.Suite
}

func TestConfigLoaderTestSuite(t *testing.T) {
	suite.Run(t, new(ConfigLoaderTestSuite))
}

func (s *ConfigLoaderTestSuite) TestLoadConfig() {
	config, err := LoadConfig("test_config_file.yaml")
	assert.Nil(s.T(), err)
	expectedConfig := models.Config{
		Probes: []models.Probe{
			{
				ProbeName:     "testProbe1",
				ProbeaAddress: "testProbeAddress1",
			},
			{
				ProbeName:     "testProbe2",
				ProbeaAddress: "testProbeAddress2",
			},
		},
		Interval:          10,
		NumberOfIntervals: 11,
		FileName:          "testFileName.xlsx",
	}
	assert.Equal(s.T(), &expectedConfig, config)
}
