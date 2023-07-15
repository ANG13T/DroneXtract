package telemetry

import (
	"encoding/csv"
	"github.com/ANG13T/DroneXtract/helpers"
	"os"
	"fmt"
)

// Disploay flight path GPS coordinates and corresponding map

type DJI_Flight_Path_Map struct {
	fileName        string
}

func NewDJI_Flight_Path_Map(fileName string) *DJI_Flight_Path_Map {
	parser := DJI_Flight_Path_Map{
		fileName: fileName,
	}
	return &parser
}

func (parser *DJI_Flight_Path_Map) ExecuteFlightPathAnalysis() {
	parser.PrintGPSCoordinates()
}

func (parser *DJI_Flight_Path_Map) PrintGPSCoordinates() {
	file, err := os.Open(parser.fileName)
	if err != nil {
		helpers.PrintErrorLog("INVALID FILE. UNABLE TO OPEN", err)
		return
	}
	defer file.Close()

	// Create a new CSV reader
	reader := csv.NewReader(file)

	// Read all records from the CSV file
	records, err := reader.ReadAll()
	
	if err != nil {
		helpers.PrintErrorLog("INVALID FILE. UNABLE TO OPEN", err)
		return
	}

	columns := records[0]

	// Print each record
	for _, record := range records {
		for in, value := range record {
			fmt.Println(columns[in], value)
		}
	}
}

func GenerateMapOutput(outputPath string) {

}