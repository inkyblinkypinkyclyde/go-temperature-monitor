package main

import (
	"main/models"
	"main/report"
)

func main() {
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
	err := report.GenerateEmptyReport("testreport.xlsx", probes)
	if err != nil {
		panic(err)
	}
}
