package report

import (
	"fmt"
	"main/models"

	"github.com/xuri/excelize/v2"
)

const (
	HEADER_ROW = 1
	FIRST_ROW  = 2
)

func GenerateEmptyReport(reportName string, probes []models.Probe) error {
	file := excelize.NewFile()
	currentColumn := 65
	file.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(currentColumn)), HEADER_ROW), "time")
	readings := []string{"temperature", "humidity"}
	for _, probe := range probes {
		for _, reading := range readings {
			currentColumn++
			file.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(currentColumn)), HEADER_ROW), fmt.Sprintf("%s %s", probe.ProbeName, reading))
		}
	}
	return file.SaveAs(reportName)
}

func LogCollectedProbeReports(collectedProbeReports models.CollectedProbeReports, row int, fileName string) error {
	file, _ := excelize.OpenFile(fileName)
	currentColumn := 65
	file.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(currentColumn)), row), collectedProbeReports.CollectedTime)
	for _, probeReport := range collectedProbeReports.ProbeReports {
		file.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(currentColumn+1)), row), probeReport.Temperature)
		file.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(currentColumn+2)), row), probeReport.Humidity)
		currentColumn = currentColumn + 2
	}
	file.Save()
	return nil
}

func GetNextEmptyRow(fileName string) (int, error) {
	report, err := getReport(fileName)
	if err != nil {
		return 0, err
	}
	return len(report) + 1, nil
}

func getReport(fileName string) ([][]string, error) {
	file, err := excelize.OpenFile(fileName)
	if err != nil {
		return [][]string{}, err
	}
	rows, err := file.GetRows("Sheet1")
	if err != nil {
		return [][]string{}, err
	}
	return rows, nil
}
