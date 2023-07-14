package parsing

import (
	"testing"
	"github.com/TwiN/go-color"
	"fmt"
)

func TestCSVAnalysis(t *testing.T) {
	t.Logf("Running Parsing Test Unit - CSV File")
	fmt.Println(color.Ize(color.Cyan, "Running Parsing Test Unit - CSV Files"))

	// Testing Parsing
	filename := `../test-data/Airdata-Files/AirdataCSV.csv`

	suite := NewDJI_CSV_Parser(filename)

	suite.ParseContents()

	t.Logf("[Parsing #1] CSV Parsing Case - PASS")
}

func TestGPXAnalysis(t *testing.T) {
	t.Logf("Running Parsing Test Unit - GPX File")
	fmt.Println(color.Ize(color.Cyan, "Running Parsing Test Unit - GPX Files"))

	// Testing Parsing
	filename := `../test-data/Airdata-Files/AirdataGPX.gpx`

	suite := NewDJI_GPX_Parser(filename)

	suite.ParseContents()

	t.Logf("[Parsing #2] GPX Parsing Case - PASS")
}


func TestKMLAnalysis(t *testing.T) {
	t.Logf("Running Parsing Test Unit - KML File")
	fmt.Println(color.Ize(color.Cyan, "Running Parsing Test Unit - KML Files"))

	// Testing Parsing
	filename := `../test-data/Airdata-Files/AirdataKML.kml`

	suite := NewDJI_KML_Parser(filename)

	suite.ParseContents()

	t.Logf("[Parsing #2] KML Parsing Case - PASS")
}