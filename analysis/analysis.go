package analysis

import (
	"github.com/ANG13T/DroneXtract/helpers"
	"strconv"
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

var warning_indicators = []float64{}

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

func (parser *DJI_Analysis) ExecuteAnalysis() {
	value := GenerateOptions()
	if (value == -1) {
		return
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