package telemetry

import (
	"encoding/csv"
	"github.com/ANG13T/DroneXtract/helpers"
	"os"
	"strconv"

	"image/color"

  sm "github.com/flopp/go-staticmaps"
  "github.com/fogleman/gg"
  "github.com/golang/geo/s2"
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

	lats := []float64{}
	longs := []float64{}

	// Print each record
	for _, record := range records {
		for in, value := range record {
			if (columns[in] == "latitude") {
				lat_val, _  := strconv.ParseFloat(value, 64)
				lats = append(lats, lat_val)
			}

			if (columns[in] == "longitude") {
				lon_val, _  := strconv.ParseFloat(value, 64)
				longs = append(longs, lon_val)
			}
		}
	}
}

func GenerateMapOutput(lats []float64, longs []float64, outputPath string) {
	ctx := sm.NewContext()
	ctx.SetSize(400, 300)
	ctx.AddObject(
	  sm.NewMarker(
		s2.LatLngFromDegrees(52.514536, 13.350151),
		color.RGBA{0xff, 0, 0, 0xff},
		16.0,
	  ),
	)
  
	img, err := ctx.Render()
	if err != nil {
	  panic(err)
	}
  
	if err := gg.SavePNG("my-map.png", img); err != nil {
	  panic(err)
	}
}