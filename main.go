package main

// OSINT, DAT File parsing, other file format parsing, crystal web server, TUI and GUI 

// 5 categories
// 1- Drone File Parsing
// 2- Telemetry Toolkit
// 3- Steganography Suite
// 4- Firmware and Binary Reader
// 5- Flight and Integrity Analysis

import (
	"fmt"
	"io/ioutil"
	"github.com/iskaa02/qalam/gradient"
	"github.com/TwiN/go-color"
	"strconv"
	"github.com/ANG13T/DroneXtract/helpers"
	"os"
	"github.com/ANG13T/DroneXtract/parsing"
	"github.com/ANG13T/DroneXtract/steganography"
	"github.com/ANG13T/DroneXtract/telemetry"
	"github.com/ANG13T/DroneXtract/analysis"
)

var category_banners = []string{"fileparsing.txt", "telemetrymapping.txt", "steganography.txt"}
var option_values = []int{5, 4, 6} 

func Banner() {
	banner, _ := ioutil.ReadFile("txt/banner.txt")
	info, _ := ioutil.ReadFile("txt/info.txt")
	options, _ := ioutil.ReadFile("txt/options.txt")
	g,_:=gradient.NewGradient("#1179ef", "cyan")	
	g.Print("\n" + string(banner))
	fmt.Println(color.Ize(color.Cyan, string(info)))
	fmt.Println(color.Ize(color.Cyan,  string(options) + "\n"))
}

func Option() {
	fmt.Print("\n ENTER INPUT > ")
	var selection string
	fmt.Scanln(&selection)
	num, err := strconv.Atoi(selection)
    if err != nil {
		fmt.Println(color.Ize(color.Red, "  [!] INVALID INPUT"))
		Option()
    } else {
        if (num >= 0  && num < 5) {
			DisplayOption(num)
		} else {
			fmt.Println(color.Ize(color.Red, "  [!] INVALID INPUT"))
			Option()
		}
    }
}

func DisplayOption(x int) {
	if (x == 0) {
		fmt.Println(color.Ize(color.Blue, " Farewell and fly high!"))
		os.Exit(1)
	} else if (x == 5) {
		Banner()
		Option()
	} else if (x > 0 && x < 5) {
		DisplayOptionInformation(x)
	} else {
		helpers.PrintError("INVALID INPUT")
		Option()
	}
}

func DisplayOptionInformation(option int) {
	returnVal := 0
	if (option == 4) {
		analysis.ExecuteAnalysis()
		Banner()
		Option()
	} else {
		contents, _ := ioutil.ReadFile("txt/" + category_banners[option - 1])
		fmt.Println(color.Ize(color.Cyan, string(contents)))
		returnVal = helpers.Option(0, option_values[option - 1])
	}

	if (returnVal == -1) {
		Banner()
		Option()
	} else {
		if (option == 1) {
			parsing.ExecuteParser(returnVal)
			Banner()
			Option()
		} else if (option == 2) {
			telemetry.ExecuteTelemetry(returnVal)
			Banner()
			Option()
		} else {
			steganography.ExecuteSteganography(returnVal)
			Banner()
			Option()
		}
	}
	
}

func main() {	
	Banner()
	Option()
}
