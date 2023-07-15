package telemetry

import (
	"testing"
	"github.com/TwiN/go-color"
	"fmt"
)

func TestTelemetryVisualizations(t *testing.T) {
	t.Logf("Running Telemetry Visualization Test Unit")
	fmt.Println(color.Ize(color.Cyan, "Running Telemetry Visualization Test Unit"))

	filename := `../test-data/Airdata-Files/AirdataCSV.csv`

	suite := NewDJI_TelemetryVisualizations(filename)

	suite.GenerateGraph(0)

	t.Logf("[Telemetry Visualizations #1] Graph Display Case - PASS")
}

func TestFlightPathMap(t *testing.T) {
	t.Logf("Running Flight Path Test Unit")
	fmt.Println(color.Ize(color.Cyan, "Running Flight Path Map Test Unit"))
	
	filename := `../test-data/Airdata-Files/AirdataCSV.csv`

	suite := NewDJI_Flight_Path_Map(filename, "flight-map.png")

	suite.ExecuteFlightPathAnalysis()

	t.Logf("[Telemetry Visualizations #1] Flight Map Case - PASS")
}