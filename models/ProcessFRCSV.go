package models

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

type NotCSVFileError struct {
	InF string
}

func (e NotCSVFileError) Error() string {
	return "Ignoring non-DJI CSV file: " + e.InF
}

type ProcessFRCSV struct {
	CsvData map[string]map[string][]string
}

func NewProcessFRCSV(path string) *ProcessFRCSV {
	p := &ProcessFRCSV{
		CsvData: make(map[string]map[string][]string),
	}
	p.ProcessPath(path)
	return p
}

func (p *ProcessFRCSV) ProcessPath(path string) {
	fmt.Println("Flight record files found:")
	fileInfo, err := os.Stat(path)
	if err != nil {
		fmt.Println(path+" was not processed because it was not recognized as a file or directory")
		return
	}

	if fileInfo.IsDir() {
		err = filepath.Walk(path, func(filePath string, fileInfo os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !fileInfo.IsDir() && p.isFRFile(filePath) {
				fmt.Println(filePath)
				p.CsvData[filePath] = make(map[string][]string)
				p.getData(filePath)
			}

			return nil
		})
	} else if p.isFRFile(path) {
		fmt.Println(path)
		p.CsvData[path] = make(map[string][]string)
		p.getData(path)
	} else {
		fmt.Println(path + " was not processed because it was not recognized as a file or directory")
		return
	}
}

func (p *ProcessFRCSV) isFRFile(filePath string) bool {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	fields, err := csvReader.Read()
	if err != nil {
		fmt.Println(err)
		return false
	}

	if !contains(fields, "latitude") || !contains(fields, "longitude") || !contains(fields, "time(millisecond)") {
		err = NotCSVFileError{filePath}
		fmt.Println(err.Error())
		return false
	}

	return true
}

func contains(fields []string, field string) bool {
	for _, f := range fields {
		if f == field {
			return true
		}
	}
	return false
}

func (p *ProcessFRCSV) getData(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	fields, err := csvReader.Read()
	if err != nil {
		fmt.Println(err)
		return
	}

	timeIndex := findIndex(fields, "time(millisecond)")
	latIndex := findIndex(fields, "latitude")
	lonIndex := findIndex(fields, "longitude")
	altIndex := findIndex(fields, "altitude(feet)")
	satIndex := findIndex(fields, "satellites")
	voltIndex := findIndex(fields, "voltage(v)")
	flycIndex := findIndex(fields, "flycStateRaw")

	for {
		record, err := csvReader.Read()
		if err != nil {
			break
		}

		time := record[timeIndex]
		lat := record[latIndex]
		lon := record[lonIndex]
		alt := record[altIndex]
		sat := record[satIndex]
		volt := record[voltIndex]
		flyc := record[flycIndex]
		if lat != "" && lon != "" && time != "" {
			p.CsvData[filePath][time] = []string{lat, lon, alt, sat, volt, flyc}
		}
	}
}

func findIndex(fields []string, field string) int {
	for i, f := range fields {
		if f == field {
			return i
		}
	}
	return -1
}

func (p *ProcessFRCSV) sortOut(ftGPSDict map[string]map[string][]string) ([]string, [][]string) {
	ftList := make([]string, 0, len(ftGPSDict))
	gpsList := make([][]string, 0, len(ftGPSDict))

	for d := range ftGPSDict {
		ftList = append(ftList, d)
	}

	sort.Strings(ftList)

	for _, d := range ftList {
		gpsList = append(gpsList, ftGPSDict[d])
	}

	return ftList, gpsList
}

func (p *ProcessFRCSV) outToFiles() {
	for f := range p.CsvData {
		outputFilePath := f + "-output.txt"
		outputFile, err := os.Create(outputFilePath)
		if err != nil {
			fmt.Println(err)
			continue
		}

		ftList, gpsList := p.sortOut(p.CsvData[f])

		for d := range ftList {
			outputFile.WriteString(ftList[d] + "," + gpsList[d][0] + "," + gpsList[d][1] + "\r\n")
		}

		outputFile.Close()
	}
}
