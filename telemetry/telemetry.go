package telemetry

import (
	"github.com/ANG13T/DroneXtract/helpers"
)

func ExecuteTelemetry(index int) {
	filePath := helpers.FileInputString()
	output := helpers.OutputPathString()
	switch in := index; in {
		case 1:
			suite := NewDJI_Flight_Path_Map(filePath, output)
			suite.ExecuteFlightPathAnalysis()
		case 2:
			suite := NewDJI_TelemetryVisualizations(filePath)
			suite.ExecuteTelemetryVisualizations()
	}
}
