package telemetry

import (
	"testing"
	"github.com/TwiN/go-color"
	"fmt"
)

func TestTelemetryVisualizations(t *testing.T) {
	t.Logf("Running Telemetry Visualization Test Unit")
	fmt.Println(color.Ize(color.Cyan, "Running Telemetry Visualization Test Unit"))

	// Testing Parsing
	filename := `../test-data/Airdata-Files/AirdataCSV.csv`

	suite := NewDJI_TelemetryVisualizations(filename)

	suite.GenerateGraph(0)
}