package datacollector

import (
	"fmt"
	"io"
	"main/models"
	"net/http"
	"strconv"
	"time"
)

func CollectDatum(url string) (float64, error) { // untested
	res, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return 0, err
	}

	return strconv.ParseFloat(string(body), 64)
}

func MockDatumCollector(mockUrl string) (float64, error) {
	return 20.5, nil
}

func CollectAllData(probes []models.Probe, time time.Time, datumCollector func(string) (float64, error)) (models.CollectedProbeReports, error) {
	collectedProbeReports := models.CollectedProbeReports{CollectedTime: time}
	for _, probe := range probes {
		temperature, err := datumCollector(fmt.Sprintf("%s/temperature", probe.ProbeaAddress))
		if err != nil {
			return collectedProbeReports, err
		}
		humidity, err := datumCollector(fmt.Sprintf("%s/humidity", probe.ProbeaAddress))
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
