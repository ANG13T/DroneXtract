package telemetry

import (
	"github.com/ANG13T/DroneXtract/helpers"
)

func ExecuteTelemetry(index int) {
	filePath := helpers.FileInputString()
	switch in := index; in {
		case 1:
			suite := NewDJI_TelemetryVisualizations(filePath)
			suite.ExecuteTelemetryVisualizations()
		case 2:
			// suite := NewDJI_KML_Parser(filePath)
			// suite.ParseContents()
	}
}
