package datacollector

import (
	"fmt"
	"io"
	"main/models"
	"net/http"
	"time"
)

func CollectDatum(url string) (string, error) { // untested
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func MockDatumCollector(mockUrl string) (string, error) {
	return "20.5", nil
}

func CollectAllData(probes []models.Probe, time time.Time, datumCollector func(string) (string, error)) (models.CollectedProbeReports, error) {
	collectedProbeReports := models.CollectedProbeReports{CollectedTime: time}
	for _, probe := range probes {
		temperature, err := datumCollector(fmt.Sprintf("%s/temperature", probe.ProbeIP))
		if err != nil {
			return collectedProbeReports, err
		}
		humidity, err := datumCollector(fmt.Sprintf("%s/humidity", probe.ProbeIP))
		if err != nil {
			return collectedProbeReports, err
		}
		probeReport := models.ProbeReport{
			Probe:       probe,
			Temperature: temperature,
			Humidity:    humidity,
		}
		collectedProbeReports.ProbeReports = append(collectedProbeReports.ProbeReports, probeReport)
	}
	return collectedProbeReports, nil
}
