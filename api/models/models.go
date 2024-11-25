package models

import "time"

type Probe struct {
	ProbeName     string `yaml:"ProbeName"`
	ProbeaAddress string `yaml:"ProbeaAddress"`
}

type ProbeReport struct {
	Probe       Probe
	Temperature float64
	Humidity    float64
}

type CollectedProbeReports struct {
	ProbeReports  []ProbeReport
	CollectedTime time.Time
}

type Config struct {
	Probes            []Probe `yaml:"probes"`
	Interval          int     `yaml:"interval"`
	NumberOfIntervals int     `yaml:"number_of_intervals"`
	FileName          string  `yaml:"filename"`
}
