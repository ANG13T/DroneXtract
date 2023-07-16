package analysis

import (
	"testing"
	"github.com/TwiN/go-color"
	"fmt"
)

func TestDJIAnalysis(t *testing.T) {
	t.Logf("Running Analysis Test Unit - CSV File")
	fmt.Println(color.Ize(color.Cyan, "Running Telemetry Test Unit - CSV Files"))

	filename := `../test-data/Airdata-Files/AirdataCSV.csv`

	suite := NewDJI_Analysis(filename)

	suite.RunAnalysis()

	t.Logf("[Parsing #1] CSV Analysis Case - PASS")
}
