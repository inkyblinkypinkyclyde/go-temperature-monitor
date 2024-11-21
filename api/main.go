package main

import (
	datacollector "main/data_collector"
	"main/models"
	"main/report"
	"time"

	"github.com/labstack/gommon/log"
)

func main() {
	probes := []models.Probe{
		{
			ProbeName:     "inside",
			ProbeaAddress: "http://192.168.4.94",
		},
	}

	fileName := "testReport.xlsx"
	err := report.GenerateEmptyReport(fileName, probes)
	if err != nil {
		panic(err)
	}

	for {
		collectedProbereports, datacollectorError := datacollector.CollectAllData(probes, time.Now(), datacollector.CollectDatum)
		if datacollectorError != nil {
			log.Info(datacollectorError)
		}
		nextEmptyRow, nextEmptyRowError := report.GetNextEmptyRow(fileName)
		if nextEmptyRowError != nil {
			log.Info(nextEmptyRowError)
			continue
		}
		if err = report.LogCollectedProbeReports(collectedProbereports, nextEmptyRow, fileName); err != nil {
			log.Info(err)
		}
		time.Sleep(1 * time.Second)
	}
}
