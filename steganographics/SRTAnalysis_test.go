package steganographics

import (
	// "github.com/ANG13T/DroneXtract/steganographics"
	"io/ioutil"
	"fmt"
	"testing"
)

func TestSRTToObject(t *testing.T) {
	suite := DJI_SRT_Parser{}

	filename := `C:\Users\AT\Desktop\DroneXtract\test-data\SRT-Files\broken_empty2.SRT`

	suite.GeneratePackets(filename)
	suite.PrintAllPackets()
}

func RunSteganographics(){
	RunSRTAnalysis()
	RunExifAnalysis()
	RunXMPAnalysis()
	RunDNGAnalysis()
}

func RunSRTAnalysis() {
	// parsing 
	filename := `C:\Users\AT\Desktop\DroneXtract\test-data\SRT-Files\m2zoom.SRT`

	content, _ := ioutil.ReadFile(filename)
	fmt.Println(content)

	// Test empty
	// broken_empty
	// Test incomplete
	// Test all formats
	// - m2zoom

	// conversion
}

// {
//   "TIMECODE":"00:00:01,000",
//   "HOME":[
//     "149.0251",
//     "-20.2532"
//   ],
//   "DATE":"2017.08.05 14:11:51",
//   "GPS":[
//     "149.0251",
//     "-20.2533",
//     "16"
//   ],
//   "BAROMETER":"1.9",
//   "ISO":"100",
//   "Shutter":"60",
//   "Fnum":"2.2"
// }

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