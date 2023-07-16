package analysis

import (
	"github.com/ANG13T/DroneXtract/helpers"
	"strconv"
	"os"
	"fmt"
	"encoding/csv"
)
var cooresponding_index = []int{8,9,10,11,14,18,19,20,22,23,24,27,28,29,30,31,32,33,34,35,36,37,38,39,40,41,42,43}
var indicators = []string{"height_sonar(feet)", "speed(mph)", "distance(feet)", "mileage(feet)", "voltage(v)", "xSpeed(mph)", "ySpeed(mph)", "zSpeed(mph)", "compass_heading(degrees)", "pitch(degrees)", "roll(degrees)", "rc_elevator", "rc_aileron", "rc_throttle", "rc_rudder", "rc_elevator(percent)", "rc_aileron(percent)", "rc_throttle(percent)", "rc_rudder(percent)", "gimbal_heading(degrees)", "gimbal_pitch(degrees)", "gimbal_roll(degrees)", "battery_percent", "current(A)", "battery_temperature(f)", "altitude(feet)", "ascent(feet)"}
var print_indicators = []string{
	"Height Sonar (feet)",
	"Speed (mph)",
	"Distance (feet)",
	"Mileage (feet)",
	"Voltage (V)",
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
	"Current (A)",
	"Battery Temperature (F)",
	"Altitude (feet)",
	"Ascent (feet)",
}

var max_variance = []float64{
		300,    // Height Sonar (feet)
		90,   // Speed (mph)
		100,  // Distance (feet)
		1000, // Mileage (feet)
		120,  // Voltage (V)
		100,   // X Speed (mph)
		30,   // Y Speed (mph)
		30,   // Z Speed (mph)
		360,  // Compass Heading (degrees)
		40,    // Pitch (degrees)
		25,    // Roll (degrees)
		1684,    // RC Elevator
		1684,    // RC Aileron
		1684,    // RC Throttle
		1684,    // RC Rudder
		1684,    // RC Elevator (percent)
		1684,    // RC Aileron (percent)
		1684,    // RC Throttle (percent)
		1684,    // RC Rudder (percent)
		360,    // Gimbal Heading (degrees)
		180,    // Gimbal Pitch (degrees)
		180,    // Gimbal Roll (degrees)
		100,   // Battery Percent
		20,    // Current (A)
		5,    // Battery Temperature (F)
		500,  // Altitude (feet)
		50,   // Ascent (feet)
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

	check := helpers.CheckFileFormat(fileName, ".csv")
	if check == false {
		helpers.PrintError("INVALID FILE FORMAT. MUST BE CSV FILE")
		return
	}

	suite := NewDJI_Analysis(fileName)
	suite.RunAnalysis()
}

func (parser *DJI_Analysis) RunAnalysis() {
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

	helpers.PrintLog("==== RUNNING FLIGHT ANALYSIS ====")

	invalid_count := 0
	valid_count := 0

	for index, _ := range indicators {
		variance := parser.GetVariance(index, records)
		valid := variance <= float64(max_variance[index])
		if (valid) {
			helpers.PrintValidLog("[√] " + print_indicators[index] + " > " +  fmt.Sprintf("%.2f", variance))
			valid_count++
		} else {
			helpers.PrintInvalidLog("[X] " + print_indicators[index] + " > " +  fmt.Sprintf("%.2f", variance) + " " + fmt.Sprintf("%.2f", max_variance[index]))
			invalid_count++
		}
	}

	helpers.PrintLog("==== FINISHED FLIGHT ANALYSIS ====")
	helpers.PrintValidLog("[√] " + strconv.Itoa(valid_count) + " VALID VARIANCE")
	helpers.PrintInvalidLog("[X] " + strconv.Itoa(invalid_count) + " INVALID VARIANCE")
	helpers.PrintLog("==================================")
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

func (parser *DJI_Analysis) GetVariance(value int, records [][]string) float64 {
	result := parser.GetCSVValues(value, records)

	max := GetMax(result)

	min := GetMin(result)

	variance := max - min

	return variance
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
			if (in == cooresponding_index[index]) {
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