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
  "fmt"
)

// Disploay flight path GPS coordinates and corresponding map

type DJI_Flight_Path_Map struct {
	fileName        string
	outputPath		string
}

func NewDJI_Flight_Path_Map(fileName string, outputPath string) *DJI_Flight_Path_Map {
	parser := DJI_Flight_Path_Map{
		fileName: fileName,
		outputPath: outputPath,
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
			val, _  := strconv.ParseFloat(value, 64)

			if (columns[in] == "latitude" && val != 0) {
				lats = append(lats, val)
			}

			if (columns[in] == "longitude" && val != 0) {
				longs = append(longs, val)
			}
		}
	}

	GenerateMapOutput(lats, longs, parser.outputPath)
}

func GenerateMapOutput(lats []float64, longs []float64, outputPath string) {
	if len(outputPath) == 0 {
		outputPath = "flight-path-map.png"
	}

	check := helpers.CheckFileFormat(outputPath, ".png")
	if check == false {
		helpers.PrintError("INVALID OUTPUT FILE FORMAT. MUST BE CSV FILE")
		return
	}
	ctx := sm.NewContext()
	ctx.SetSize(400, 300)
	ctx.AddObject(
	  sm.NewMarker(
		s2.LatLngFromDegrees(lats[0], longs[0]),
		color.RGBA{0xff, 0, 0, 0xff},
		16.0,
	  ),
	)

	ctx.AddObject(
		sm.NewMarker(
		  s2.LatLngFromDegrees(lats[1], longs[1]),
		  color.RGBA{0xff, 0, 0, 0xff},
		  16.0,
		),
	  )

	point1 := s2.LatLngFromDegrees(lats[0], longs[0])
	point2 := s2.LatLngFromDegrees(lats[1], longs[1])

	fmt.Println(lats[0], longs[0])

	pos := []s2.LatLng{point1, point2}

	path := sm.NewPath(pos, color.RGBA{0xff, 0, 0, 0xff}, 6.0)

	ctx.AddObject(path)
  
	img, err := ctx.Render()
	if err != nil {
	  panic(err)
	}
  
	if err := gg.SavePNG(outputPath, img); err != nil {
	  panic(err)
	}
}

// .\test-data\Airdata-Files\AirdataCSV.csv
// 
// 
//
