package main

import (
	"fmt"
	configloader "main/config_loader"
	datacollector "main/data_collector"
	"main/report"
	"time"
)

func main() {
	config, err := configloader.LoadConfig("config.yaml")
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
			fmt.Println(datacollectorError)
		}
		fmt.Println(collectedProbereports)
		nextEmptyRow, nextEmptyRowError := report.GetNextEmptyRow(config.FileName)
		if nextEmptyRowError != nil {
			fmt.Println(nextEmptyRowError)
			continue
		}
		if err = report.LogCollectedProbeReports(collectedProbereports, nextEmptyRow, config.FileName); err != nil {
			fmt.Println(err)
		}
		fmt.Println(time.Now())
		time.Sleep(time.Duration(config.Interval) * time.Second)
	}
}
