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
  // "fmt"
)

type Coordinate struct {
	Latitude  float64
	Longitude float64
}

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

	coors := []Coordinate{}

	// Print each record
	for _, record := range records {

		lat_val := 0.0
		lon_val := 0.0

		for in, value := range record {
			val, _  := strconv.ParseFloat(value, 64)

			if (columns[in] == "latitude" && val != 0) {
				lat_val = val
			}

			if (columns[in] == "longitude" && val != 0) {
				lon_val = val
			}
		}

		if (lat_val != 0.0 && lon_val != 0.0) {
			coor_val := Coordinate{Latitude: lat_val, Longitude: lon_val}
			coors = append(coors, coor_val)
		}
	}

	GenerateMapOutput(coors, parser.outputPath)
}

func GenerateMapOutput(coors []Coordinate, outputPath string) {
	if len(outputPath) == 0 {
		outputPath = "flight-path-map.png"
	}

	if len(coors) > 10 {
		coors = downsampleCoordinates(coors, 10)
	}

	check := helpers.CheckFileFormat(outputPath, ".png")
	if check == false {
		helpers.PrintError("INVALID OUTPUT FILE FORMAT. MUST BE CSV FILE")
		return
	}
	ctx := sm.NewContext()
	ctx.SetSize(400, 300)

	for index, coor := range coors {
		ctx.AddObject(
			sm.NewMarker(
			  s2.LatLngFromDegrees(coor.Latitude, coor.Longitude),
			  color.RGBA{0xff, 0, 0, 0xff},
			  16.0,
			),
		)

		if (index < len(coors) - 1) {
			point1 := s2.LatLngFromDegrees(coors[index].Latitude, coors[index].Longitude)
			point2 := s2.LatLngFromDegrees(coors[index + 1].Latitude, coors[index + 1].Longitude)
	
			pos := []s2.LatLng{point1, point2}
	
			path := sm.NewPath(pos, color.RGBA{0xff, 0, 0, 0xff}, 6.0)
	
			ctx.AddObject(path)
		}
	} 

	PrintCoordinates(coors, len(coors) <= 10)
  
	img, err := ctx.Render()
	if err != nil {
	  panic(err)
	}
  
	if err := gg.SavePNG(outputPath, img); err != nil {
	  panic(err)
	}

	helpers.PrintLog("Created Flight Path Map at " + outputPath)
}

func PrintCoordinates(coordinates []Coordinate, downsampled bool) {
	if (downsampled) {
		helpers.GenTableHeader("Downsampled GPS Coordinates");
	} else {
		helpers.GenTableHeader("GPS Coordinates");
	}

	for in, coor := range coordinates {
		str_lat := strconv.FormatFloat(coor.Latitude, 'f', -1, 64)
		str_lon := strconv.FormatFloat(coor.Longitude, 'f', -1, 64)
		helpers.GenRowString("Coordinate " + strconv.Itoa(in + 1), "(" + str_lat + ", " + str_lon + ")")
	}

	helpers.GenTableFooter();
}

func downsampleCoordinates(coordinates []Coordinate, targetLength int) []Coordinate {
	length := len(coordinates)
	if targetLength >= length {
		return coordinates
	}

	ratio := float64(length) / float64(targetLength)
	result := make([]Coordinate, targetLength)
	resultIndex := 0

	for i := 0; i < targetLength; i++ {
		rangeStart := int(float64(i) * ratio)
		rangeEnd := int(float64(i+1) * ratio)

		// Calculate the average of latitude and longitude within the range
		sumLat := 0.0
		sumLon := 0.0
		for j := rangeStart; j < rangeEnd; j++ {
			sumLat += coordinates[j].Latitude
			sumLon += coordinates[j].Longitude
		}
		averageLat := sumLat / float64(rangeEnd-rangeStart)
		averageLon := sumLon / float64(rangeEnd-rangeStart)

		result[resultIndex] = Coordinate{
			Latitude:  averageLat,
			Longitude: averageLon,
		}
		resultIndex++
	}

	return result
}