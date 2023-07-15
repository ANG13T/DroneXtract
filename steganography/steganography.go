package steganography

import (
	"github.com/TwiN/go-color"
	"fmt"
	"strings"
	"log"
	"path/filepath"
	"github.com/ANG13T/DroneXtract/helpers"
	"io/ioutil"
)

// TODO: steg UI

func ExecuteSteganography(index int) {
	output := PrintParsingOptions(index)
	filePath := helpers.FileInputString()
	switch in := index; in {
		case 1:
			suite := NewDJI_EXIF_Parser(filePath)
			suite.ExecuteEXIFAnalysis(output)
		case 2:
			suite := NewDJI_DNG_Parser(filePath)
			suite.ExecuteDNGAnalysis(output)
		case 3:
			suite := NewDJI_SRT_Parser(filePath)
			suite.ExecuteSRTAnalysis(output)
		case 4:
			suite := NewDJI_XMP_Parser(filePath)
			suite.ExecuteXMPAnalysis(output)
	}
}

func PrintParsingOptions(index int) int {
	var parsing_banners = []string{"exif.txt", "kml.txt", "gpx.txt", "xmp.txt"}
	contents, _ := ioutil.ReadFile("txt/steganography/" + parsing_banners[index - 1])
	fmt.Println(color.Ize(color.Cyan, string(contents)))
	result := helpers.Option(0, 5)
	return result
}

func GenTableHeader(name string, containBreak bool) {
	if containBreak {
		fmt.Println(color.Ize(color.Blue, "\n    ╔══════════════════════════════════════════════════════════════════════════════╗"))
	} else {
		fmt.Println(color.Ize(color.Blue, "    ╔══════════════════════════════════════════════════════════════════════════════╗"))
	}
	var amount = (78 - len(name)) / 2
	var extraPadding = 1
	if len(name) % 2 == 0 {
		extraPadding = 0
	}
	fmt.Println(color.Ize(color.Blue, "    ║" +  strings.Repeat(" ", amount) + name + strings.Repeat(" ", amount + extraPadding) + "║"))
	fmt.Println(color.Ize(color.Blue, "    ╠══════════════════════════════════════════════════════════════════════════════╣"))
}

func GenTableHeaderModified(name string) {
	fmt.Println(color.Ize(color.Blue, "    ╠══════════════════════════════════════════════════════════════════════════════╣"))
	var amount = (78 - len(name)) / 2
	var extraPadding = 1
	if len(name) % 2 == 0 {
		extraPadding = 0
	}
	fmt.Println(color.Ize(color.Blue, "    ║" +  strings.Repeat(" ", amount) + name + strings.Repeat(" ", amount + extraPadding) + "║"))
	fmt.Println(color.Ize(color.Blue, "    ╠══════════════════════════════════════════════════════════════════════════════╣"))
}

func GenRowString(intro string, input string) {
	if input == "UNSPECIFIED" {
		return
	}
	var totalCount int = 4 + len(intro) + len(input) + 2
	var useCount = 80 - totalCount
	if useCount < 0 { 
		useCount = 0
	}
	var val = "    ║ " + intro + ": " + input + strings.Repeat(" ", useCount) + " ║"
	fmt.Println(color.Ize(color.Blue, val))
}

func GenTableFooter() {
	fmt.Println(color.Ize(color.Blue, "    ╚══════════════════════════════════════════════════════════════════════════════╝"))
}


func PrintError(message string) {
	fmt.Println(color.Ize(color.Red, "[ERROR] " + message))
}

func PrintErrorLog(message string, err error) {
	fmt.Println(color.Ize(color.Red, message))
	log.Println(color.Ize(color.Red, "[ERROR]"), err)
}

func CheckFileFormat(path string, exten string) bool {
	extension := strings.ToLower(filepath.Ext(path))
	return (extension == exten)
}