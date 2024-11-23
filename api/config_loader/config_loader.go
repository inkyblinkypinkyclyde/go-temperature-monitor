package configloader

import (
	"fmt"
	"main/models"
	"os"

	"gopkg.in/yaml.v2"
)

func LoadConfig(configFile string) (*models.Config, error) {
	file, err := os.Open(configFile)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %w", err)
	}
	defer file.Close()

	var config models.Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, fmt.Errorf("failed to decode YAML file: %w", err)
	}
	return &config, nil
}
