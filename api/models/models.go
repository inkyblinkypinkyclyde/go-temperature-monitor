package models

import "time"

type Probe struct {
	ProbeName     string
	ProbeaAddress string
}

type ProbeReport struct {
	Probe       Probe
	Temperature string
	Humidity    string
}

type CollectedProbeReports struct {
	ProbeReports  []ProbeReport
	CollectedTime time.Time
}
