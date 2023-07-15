package analysis

import (
	"github.com/ANG13T/DroneXtract/helpers"
	"strconv"
	"os"
	"encoding/csv"
)

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

var max_variance = []float64{
		50,   // Height Above Takeoff (feet)
		20,   // Height Above Ground at Drone Location (feet)
		100,  // Ground Elevation at Drone Location (feet)
		100,  // Ground Elevation at Drone Location (feet)
		500,  // Altitude Above Sea Level (feet)
		5,    // Height Sonar (feet)
		10,   // Speed (mph)
		100,  // Distance (feet)
		1000, // Mileage (feet)
		5,    // Satellites
		1,    // GPS Level
		0.5,  // Voltage (V)
		500,  // Max Altitude (feet)
		10,   // Max Ascent (feet)
		20,   // Max Speed (mph)
		5000, // Max Distance (feet)
		10,   // X Speed (mph)
		10,   // Y Speed (mph)
		10,   // Z Speed (mph)
		5,    // Compass Heading (degrees)
		5,    // Pitch (degrees)
		5,    // Roll (degrees)
		2,    // RC Elevator
		2,    // RC Aileron
		2,    // RC Throttle
		2,    // RC Rudder
		2,    // RC Elevator (percent)
		2,    // RC Aileron (percent)
		2,    // RC Throttle (percent)
		2,    // RC Rudder (percent)
		5,    // Gimbal Heading (degrees)
		5,    // Gimbal Pitch (degrees)
		5,    // Gimbal Roll (degrees)
		1,    // Battery Percent
		0.2,  // Voltage Cell 1
		0.2,  // Voltage Cell 2
		0.2,  // Voltage Cell 3
		0.2,  // Voltage Cell 4
		0.2,  // Voltage Cell 5
		0.2,  // Voltage Cell 6
		5,    // Current (A)
		5,    // Battery Temperature (F)
		500,  // Altitude (feet)
		50,   // Ascent (feet)
		10,   // Flyc State Raw
}

type DJI_Analysis struct {
	fileName        string
	outputPath		string
}

func NewDJI_Analysis(fileName string) *DJI_Analysis {
	parser := DJI_Analysis{
		fileName: fileName,
	}
	return &parser
}

func ExecuteAnalysis() {
	fileName := helpers.FileInputString()
	suite := NewDJI_Analysis(fileName)
	suite.RunAnalysis()
}

func (parser *DJI_Analysis) RunAnalysis() {
	value := GenerateOptions()
	if (value == -1) {
		return
	}

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

	for index, _ := range indicators {
		valid := parser.IsValidValue(index, records)
		if (valid == false) {
			helpers.PrintLog(print_indicators[index] + " IS VALID")
		} else {
			helpers.PrintLog(print_indicators[index] + " NOT VALID")
		}
	}
}

func GenerateOptions() int {
	helpers.GenTableHeader("Select Value to Analyze")
	for index, record := range print_indicators {
		helpers.GenRowString(strconv.Itoa(index + 1), record)
	}
	helpers.GenRowString(strconv.Itoa(len(print_indicators) + 1), "Back to Main Menu")
	helpers.GenRowString("0", "Exit DroneXtract")
	helpers.GenTableFooter()
	return helpers.Option(0, len(print_indicators))
}

func GetMax(numbers []float64) float64 {
	if len(numbers) == 0 {
		return 0.0
	}

	maxValue := numbers[0]
	for _, num := range numbers {
		if num > maxValue {
			maxValue = num
		}
	}

	return maxValue
}

func (parser *DJI_Analysis) IsValidValue(value int, records [][]string) bool {
	result := parser.GetCSVValues(value, records)

	max := GetMax(result)

	min := GetMin(result)

	variance := max - min

	return variance < max_variance[value]
}

func GetMin(numbers []float64) float64 {
	if len(numbers) == 0 {
		return 0.0
	}

	minValue := numbers[0]
	for _, num := range numbers {
		if num < minValue {
			minValue = num
		}
	}

	return minValue
}

func (parser *DJI_Analysis) GetCSVValues(index int, records [][]string) []float64 {
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

	if (len(output) > 40) {
		output = downsampleArray(output, 40)
	}

	return output
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