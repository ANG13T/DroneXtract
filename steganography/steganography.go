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

func ExecuteSteganography(index int) {
	output := PrintParsingOptions(index)
	if (output == -1) {
		return
	}
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

