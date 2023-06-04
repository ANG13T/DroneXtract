package steganography

import (
	"testing"
	"github.com/TwiN/go-color"
	"fmt"
)

func TestSRTAnalysis(t *testing.T) {
	t.Logf("Running Steganography Test Unit - SRT Files")
	fmt.Println(color.Ize(color.Cyan, "Running Steganography  Test Unit - SRT Files"))

	// Testing Parsing
	filename := `..\test-data\SRT-Files\mavic_air.SRT`

	suite := NewDJI_SRT_Parser(filename)

	suite.PrintFileMetadata()
	suite.GeneratePackets()
	suite.PrintAllPackets()

	test_packet := suite.packets[0]

	if test_packet.frame_count == "1" && test_packet.diff_time == "39ms" && test_packet.iso == "100" && test_packet.shutter == "1/240.0" && test_packet.fnum == "280" && test_packet.ev == "0" && test_packet.ct == "5627" && test_packet.date == "2018-02-19 08:04:51:265.847" && test_packet.time_stamp == "00:00:00" {
		t.Logf("[Steganography #1] SRT Analysis Parsing Case - PASS")
		fmt.Println(color.Ize(color.Green, "[Steganography  #1] SRT Analysis Parsing Case - PASS"))
	} else {
		t.Errorf("[Steganography #1] SRT Analysis Parsing Case - ERROR")
		fmt.Println(color.Ize(color.Red, "[Steganography  #1] SRT Analysis Parsing Case - ERROR"))
	}

	// Testing Conversion
	suite.ExporttoGeoJSON(`..\output\DJI_SRT_GEOJSON_OUTPUT.geojson`)
}


// func TestEXIFAnalysis(t *testing.T) {
// 	t.Logf("Running Steganography  Test Unit - EXIF Files")
// 	fmt.Println(color.Ize(color.Cyan, "Running Steganography  Test Unit - SRT Files"))

// 	filename := `..\test-data\JPG-Files\DJI_0001.jpg`
// 	suite := NewDJI_EXIF_Parser(filename)

// 	// Parsing
// 	suite.Read()

// 	// Conversion to TXT
// 	outputTXTFileName := `..\output\DJI_JPG_EXIF_OUTPUT.txt`
// 	suite.ExportToTXT(outputTXTFileName)

// 	// Conversion to CSV
// 	outputCSVFileName := `..\output\DJI_JPG_EXIF_OUTPUT.csv`
// 	suite.ExportToCSV(outputCSVFileName)

// 	// Conversion to JSON
// 	outputJSONFileName := `..\output\DJI_JPG_EXIF_OUTPUT.json`
// 	suite.ExportToJSON(outputJSONFileName)
// }

// func RunXMPAnalysis(t *testing.T) {
// 	t.Logf("Running Steganography  Test Unit - XMP Files")
// 	fmt.Println(color.Ize(color.Cyan, "Running Steganography  Test Unit - XMP Files"))

// 	filename := `..\test-data\JPG-Files\DJI_0001.jpg`
// 	suite := NewDJI_XMP_Parser(filename)

// 	// Parsing
// 	suite.Read()

// 	// Conversion to TXT
// 	outputTXTFileName := `..\output\DJI_JPG_XMP_OUTPUT.txt`
// 	suite.ExportToTXT(outputTXTFileName)

// 	// Conversion to CSV
// 	outputCSVFileName := `..\output\DJI_JPG_XMP_OUTPUT.csv`
// 	suite.ExportToCSV(outputCSVFileName)

// 	// Conversion to JSON
// 	outputJSONFileName := `..\output\DJI_JPG_XMP_OUTPUT.json`
// 	suite.ExportToJSON(outputJSONFileName)
// }

// func RunDNGAnalysis(t *testing.T) {
// 	t.Logf("Running Steganography  Test Unit - DNG Files")
// 	fmt.Println(color.Ize(color.Cyan, "Running Steganography  Test Unit - DNG Files"))

// 	filename := `..\test-data\DNG-Files\DJI_0234.dng`
// 	suite := NewDJI_DNG_Parser(filename)

// 	// Parsing
// 	suite.Read()

// 	// Conversion to PNG
// 	output := `..\output\DJI_0234.png`
// 	suite.DNGtoPNG(output)

// 	// Conversion to TXT
// 	outputTXTFileName := `..\output\DJI_JPG_DNG_OUTPUT.txt`
// 	suite.ExportToTXT(outputTXTFileName)

// 	// Conversion to CSV
// 	outputCSVFileName := `..\output\DJI_JPG_DNG_OUTPUT.csv`
// 	suite.ExportToCSV(outputCSVFileName)

// 	// Conversion to JSON
// 	outputJSONFileName := `..\output\DJI_JPG_DNG_OUTPUT.json`
// 	suite.ExportToJSON(outputJSONFileName)
// }
