package tests

import (
	"github.com/ANG13T/DroneXtract/steganographics"
	"io/ioutil"
)

func RunSteganographics(){
	RunSRTAnalysis()
	RunExifAnalysis()
	RunXMPAnalysis()
	RunDNGAnalysis()
}

func RunSRTAnalysis() {
	// parsing 
	filename := "../"

	content, _ := ioutil.ReadFile(filename)
	// conversion
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