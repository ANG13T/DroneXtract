package parsing

import (
	"github.com/ANG13T/DroneXtract/helpers"
	"github.com/tkrajina/gpxgo/gpx"
	"path/filepath"
	"fmt"
)

type DJI_GPX_Parser struct {
	fileName        string
}

func NewDJI_GPX_Parser(fileName string) *DJI_GPX_Parser {
	check := CheckFileFormat(fileName, ".gpx")
	if check == false {
		helpers.PrintError("INVALID FILE FORMAT. MUST BE GPX FILE")
		return nil
	}

	parser := DJI_GPX_Parser{
		fileName: fileName,
	}
	return &parser
}

func (parser *DJI_GPX_Parser) ParseContents() {
	gpxFile, err := gpx.ParseFile(parser.fileName)

	if err != nil {
		fmt.Println("Error opening gpx file: ", err)
		return
	}

	gpxPath, _ := filepath.Abs(parser.fileName)

	fmt.Print("File: ", gpxPath, "\n")

	fmt.Println(gpxFile.GetGpxInfo())
}
