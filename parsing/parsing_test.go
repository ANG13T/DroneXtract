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

	t.Logf("[Steganography #1] CSV Parsing Case - PASS")
}