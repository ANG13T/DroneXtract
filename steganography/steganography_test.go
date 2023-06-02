package steganography

import (
	"testing"
	"github.com/TwiN/go-color"
	"fmt"
)

// SRT - 2 tests
// XMP - 2 tests
// EXIF - 2 tests
// DNG - 2 tests

// func TestSRTAnalysis(t *testing.T) {
// 	t.Logf("Running Steganography  Test Unit - SRT Files")
// 	fmt.Println(color.Ize(color.Cyan, "Running Steganography  Test Unit - SRT Files"))

// 	// Testing Parsing
// 	filename := `C:\Users\AT\Desktop\DroneXtract\test-data\SRT-Files\mavic_air.SRT`

// 	suite := NewDJI_SRT_Parser(filename)

// 	suite.PrintFileMetadata()
// 	suite.GeneratePackets()
// 	suite.PrintAllPackets()

// 	test_packet := suite.packets[0]

// 	if test_packet.frame_count == "1" && test_packet.diff_time == "39ms" && test_packet.iso == "100" && test_packet.shutter == "1/240.0" && test_packet.fnum == "280" && test_packet.ev == "0" && test_packet.ct == "5627" && test_packet.date == "2018-02-19 08:04:51:265.847" && test_packet.time_stamp == "00:00:00" {
// 		t.Logf("[Steganography  #1] SRT Analysis Parsing Case - PASS")
// 		fmt.Println(color.Ize(color.Green, "[Steganography  #1] SRT Analysis Parsing Case - PASS"))
// 	} else {
// 		t.Errorf("[Steganography  #1] SRT Analysis Parsing Case - ERROR")
// 		fmt.Println(color.Ize(color.Red, "[Steganography  #1] SRT Analysis Parsing Case - ERROR"))
// 	}

// 	// Testing Conversion
// }


func TestEXIFAnalysis(t *testing.T) {
	t.Logf("Running Steganography  Test Unit - EXIF Files")
	fmt.Println(color.Ize(color.Cyan, "Running Steganography  Test Unit - SRT Files"))

	filename := `C:\Users\AT\Desktop\DroneXtract\test-data\JPG-Files\DJI_0001.jpg`
	suite := NewDJI_EXIF_Parser(filename)

	// Parsing
	suite.Read()

	// Conversion to TXT
	outputFileName := `C:\Users\AT\Desktop\DroneXtract\output\DJI_JPG_EXIF_OUTPUT.txt`
	suite.ExportToTXT(outputFileName)

	// Conversion to CSV

	// Conversion to JSON
}

func RunXMPAnalysis(t *testing.T) {
	t.Logf("Running Steganography  Test Unit - XMP Files")
	fmt.Println(color.Ize(color.Cyan, "Running Steganography  Test Unit - XMP Files"))

	// TODO
	filename := `C:\Users\AT\Desktop\DroneXtract\test-data\JPG-Files\DJI_0001.jpg`
	suite := NewDJI_XMP_Parser(filename)

	// parsing
	suite.Read()

	// conversion to txt, csv, and json
}

func RunDNGAnalysis(t *testing.T) {
	t.Logf("Running Steganography  Test Unit - DNG Files")
	fmt.Println(color.Ize(color.Cyan, "Running Steganography  Test Unit - DNG Files"))

	filename := `C:\Users\AT\Desktop\DroneXtract\test-data\DNG-Files\DJI_0234.dng`
	suite := NewDJI_DNG_Parser(filename)

	// parsing
	suite.Read()

	// to png
	output := `C:\Users\AT\Desktop\DroneXtract\output\DJI_0234.png`
	suite.DNGtoPNG(output)

	// conversion to txt, csv, and json
}
