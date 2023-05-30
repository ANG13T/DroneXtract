package steganographics

import (
	"testing"
	"github.com/TwiN/go-color"
	"fmt"
)

// SRT - 2 tests
// XMP - 2 tests
// EXIF - 2 tests
// DNG - 2 tests

func TestSRTAnalysis(t *testing.T) {
	// Testing Parsing
	filename := `C:\Users\AT\Desktop\DroneXtract\test-data\SRT-Files\mavic_air.SRT`

	suite := NewDJI_SRT_Parser(filename)

	suite.PrintFileMetadata()
	suite.GeneratePackets()
	suite.PrintAllPackets()

	test_packet := suite.packets[0]

	if test_packet.frame_count == "1" && test_packet.diff_time == "39ms" && test_packet.iso == "100" && test_packet.shutter == "1/240.0" && test_packet.fnum == "280" && test_packet.ev == "0" && test_packet.ct == "5627" && test_packet.date == "2018-02-19 08:04:51:265.847" && test_packet.time_stamp == "00:00:00" {
		t.Logf("[Steganographics #1] SRT Analysis Parsing Case - PASS")
		fmt.Println(color.Ize(color.Green, "[Steganographics #1] SRT Analysis Parsing Case - PASS"))
	} else {
		t.Errorf("[Steganographics #1] SRT Analysis Parsing Case - ERROR")
		fmt.Println(color.Ize(color.Red, "[Steganographics #1] SRT Analysis Parsing Case - ERROR"))
	}

	// Testing Conversion
}


func RunExifAnalysis() {
	// to text
	// parsing
}

func RunXMPAnalysis() {
	// to text
	// parsing
}

func RunDNGAnalysis() {
	// to text
	// parsing
	// to png
}
