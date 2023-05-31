package main

// OSINT, DAT File parsing, other file format parsing, crystal web server, TUI and GUI 

import (
	"fmt"
	"io/ioutil"
	"github.com/iskaa02/qalam/gradient"
	"github.com/TwiN/go-color"
	"strconv"
	// "github.com/ANG13T/DroneXtract/forensics"
	"os"
	"github.com/ANG13T/DroneXtract/steganography"
	// "github.com/twpayne/go-kml/v3"
)

func Banner() {
	banner, _ := ioutil.ReadFile("txt/banner.txt")
	info, _ := ioutil.ReadFile("txt/info.txt")
	options, _ := ioutil.ReadFile("txt/options.txt")
	g,_:=gradient.NewGradient("#1179ef", "cyan")	
	g.Print("\n" + string(banner))
	fmt.Println(color.Ize(color.Cyan, string(info)))
	fmt.Println(color.Ize(color.Cyan, "\n" + string(options) + "\n"))
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
	} else if (x == 1) {
		Banner()
		// Option()
		steganography.ExampleExiftool_Read()
	} else if (x == 2) {
		Banner()
	
		steganography.DNGtoPNG()
	} else if (x == 3) {
		Banner()
		Option()
	}else if (x == 4) {
		Banner()
		Option()
	}
}

func main() {	
	Banner()
	Option()
}
