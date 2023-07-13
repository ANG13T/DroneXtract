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

	// if test_packet.frame_count == "1" && test_packet.diff_time == "39ms" && test_packet.iso == "100" && test_packet.shutter == "1/240.0" && test_packet.fnum == "280" && test_packet.ev == "0" && test_packet.ct == "5627" && test_packet.date == "2018-02-19 08:04:51:265.847" && test_packet.time_stamp == "00:00:00" {
	// 	t.Logf("[Steganography #1] SRT Analysis Parsing Case - PASS")
	// 	fmt.Println(color.Ize(color.Green, "[Steganography  #1] SRT Analysis Parsing Case - PASS"))
	// } else {
	// 	t.Errorf("[Steganography #1] SRT Analysis Parsing Case - ERROR")
	// 	fmt.Println(color.Ize(color.Red, "[Steganography  #1] SRT Analysis Parsing Case - ERROR"))
	// }

	// // Testing Parsing Print
	// suite.ExportToGeoJSON(`../output/DJI_SRT_GEOJSON_OUTPUT.geojson`)
}