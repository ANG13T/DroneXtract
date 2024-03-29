package telemetry

import (
	"github.com/guptarohit/asciigraph"
	"fmt"
	"github.com/ANG13T/DroneXtract/helpers"
	"encoding/csv"
	"github.com/TwiN/go-color"
	"os"
	"strconv"
)

var downsample_number_telemetry = helpers.GetEnvVariable("TELEMETRY_VIS_DOWNSAMPLE")

type DJI_TelemetryVisualizations struct {
	fileName        string
}

func NewDJI_TelemetryVisualizations(fileName string) *DJI_TelemetryVisualizations {
	check := helpers.CheckFileFormat(fileName, ".csv")
	if check == false {
		helpers.PrintError("INVALID FILE FORMAT. MUST BE CSV FILE")
		return nil
	}

	parser := DJI_TelemetryVisualizations{
		fileName: fileName,
	}
	return &parser
}

var indicators = []string{"height_above_takeoff(feet)", "height_above_ground_at_drone_location(feet)", "ground_elevation_at_drone_location(feet)", "ground_elevation_at_drone_location(feet)", "altitude_above_seaLevel(feet)", "height_sonar(feet)", "speed(mph)", "distance(feet)", "mileage(feet)", "satellites", "gpslevel", "voltage(v)", "max_altitude(feet)", "max_ascent(feet)", "max_speed(mph)", "max_distance(feet)", "xSpeed(mph)", "ySpeed(mph)", "zSpeed(mph)", "compass_heading(degrees)", "pitch(degrees)", "roll(degrees)", "rc_elevator", "rc_aileron", "rc_throttle", "rc_rudder", "rc_elevator(percent)", "rc_aileron(percent)", "rc_throttle(percent)", "rc_rudder(percent)", "gimbal_heading(degrees)", "gimbal_pitch(degrees)", "gimbal_roll(degrees)", "battery_percent", "voltageCell1", "voltageCell2", "voltageCell3", "voltageCell4", "voltageCell5", "voltageCell6", "current(A)", "battery_temperature(f)", "altitude(feet)", "ascent(feet)", "flycStateRaw"}
var print_indicators = []string{
	"Height Above Takeoff (feet)",
	"Height Above Ground at Drone Location (feet)",
	"Ground Elevation at Drone Location (feet)",
	"Ground Elevation at Drone Location (feet)",
	"Altitude Above Sea Level (feet)",
	"Height Sonar (feet)",
	"Speed (mph)",
	"Distance (feet)",
	"Mileage (feet)",
	"Satellites",
	"GPS Level",
	"Voltage (V)",
	"Max Altitude (feet)",
	"Max Ascent (feet)",
	"Max Speed (mph)",
	"Max Distance (feet)",
	"X Speed (mph)",
	"Y Speed (mph)",
	"Z Speed (mph)",
	"Compass Heading (degrees)",
	"Pitch (degrees)",
	"Roll (degrees)",
	"RC Elevator",
	"RC Aileron",
	"RC Throttle",
	"RC Rudder",
	"RC Elevator (percent)",
	"RC Aileron (percent)",
	"RC Throttle (percent)",
	"RC Rudder (percent)",
	"Gimbal Heading (degrees)",
	"Gimbal Pitch (degrees)",
	"Gimbal Roll (degrees)",
	"Battery Percent",
	"Voltage Cell 1",
	"Voltage Cell 2",
	"Voltage Cell 3",
	"Voltage Cell 4",
	"Voltage Cell 5",
	"Voltage Cell 6",
	"Current (A)",
	"Battery Temperature (F)",
	"Altitude (feet)",
	"Ascent (feet)",
	"Flyc State Raw",
}

func (parser *DJI_TelemetryVisualizations) ExecuteTelemetryVisualizations() {
	value := GenerateOptions()
	if (value == -1) {
		return
	}
	parser.GenerateGraph(value)
}

func (parser *DJI_TelemetryVisualizations) GenerateGraph(index int) {
	index = index - 1
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

	output := []float64{}

	// Print each record
	for _, record := range records {
		for in, value := range record {
			if (in == index) {
				val, _ :=  strconv.ParseFloat(value, 64)
				output = append(output, val)
			}
		}
	}

	if (len(output) > downsample_number_telemetry) {
		output = downsampleArray(output, downsample_number_telemetry)
	}

	graph := asciigraph.Plot(output, asciigraph.Height(10), asciigraph.Width(100), asciigraph.Caption(print_indicators[index]))

    fmt.Println(color.Ize(color.Cyan,graph))
}

func downsampleArray(data []float64, targetLength int) []float64 {
	length := len(data)
	if targetLength >= length {
		return data
	}

	ratio := float64(length) / float64(targetLength)
	result := make([]float64, targetLength)
	resultIndex := 0

	for i := 0; i < targetLength; i++ {
		rangeStart := int(float64(i) * ratio)
		rangeEnd := int(float64(i+1) * ratio)

		// Calculate the average within the range
		sum := 0.0
		for j := rangeStart; j < rangeEnd; j++ {
			sum += data[j]
		}
		average := sum / float64(rangeEnd-rangeStart)

		result[resultIndex] = average
		resultIndex++
	}

	result[0] = data[0]

	return result
}

func GenerateOptions() int {
	helpers.GenTableHeader("Select Value to Visualize")
	for index, record := range print_indicators {
		helpers.GenRowString(strconv.Itoa(index + 1), record)
	}
	helpers.GenRowString(strconv.Itoa(len(print_indicators) + 1), "Back to Main Menu")
	helpers.GenRowString("0", "Exit DroneXtract")
	helpers.GenTableFooter()
	return helpers.Option(0, len(print_indicators))
}