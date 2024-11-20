package models

import "time"

type ProbeReport struct {
	ProbeName   string
	ProbeIP     string
	Temperature int
	Humidity    int
}

type CollectedProbeReports struct {
	ProbeReports  []ProbeReport
	CollectedTime time.Time
}
