package models

import "time"

type TemperatureReport struct {
	CollectorLocation string
	RcvdTime          time.Time
	Temperature       int
	Humidity          int
}
