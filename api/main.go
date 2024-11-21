package main

import (
	_ "embed"
	"fmt"
	datacollector "main/data_collector"
	"main/models"
	"main/report"
	"os"
	"time"

	"github.com/labstack/gommon/log"
	"gopkg.in/yaml.v2"
)

//go:embed config.yaml
var configFile []byte

type Config struct {
	Probes            []models.Probe `yaml:"probes"`
	Interval          int            `yaml:"interval"`
	NumberOfIntervals int            `yaml:"number_of_intervals"`
	FileName          string         `yaml:"filename"`
}

func main() {
	config, err := loadConfig()
	if err != nil {
		panic(err)
	}

	for _, probe := range config.Probes {
		fmt.Printf("Probe Name: %s, Address: %s\n", probe.ProbeName, probe.ProbeaAddress)
	}

	err = report.GenerateEmptyReport(config.FileName, config.Probes)
	if err != nil {
		panic(err)
	}

	for i := 0; i < config.NumberOfIntervals; i++ {
		collectedProbereports, datacollectorError := datacollector.CollectAllData(config.Probes, time.Now(), datacollector.CollectDatum)
		if datacollectorError != nil {
			log.Info(datacollectorError)
		}
		nextEmptyRow, nextEmptyRowError := report.GetNextEmptyRow(config.FileName)
		if nextEmptyRowError != nil {
			log.Info(nextEmptyRowError)
			continue
		}
		if err = report.LogCollectedProbeReports(collectedProbereports, nextEmptyRow, config.FileName); err != nil {
			log.Info(err)
		}
		fmt.Println(time.Now())
		time.Sleep(time.Duration(config.Interval) * time.Second)
	}
}

func loadConfig() (*Config, error) {
	file, err := os.Open("config.yaml")
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %w", err)
	}
	defer file.Close()

	var config Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, fmt.Errorf("failed to decode YAML file: %w", err)
	}
	return &config, nil
}
