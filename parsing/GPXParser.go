package parsing

import (
	"github.com/ANG13T/DroneXtract/helpers"
	"github.com/tkrajina/gpxgo/gpx"
	"path/filepath"
	"github.com/TwiN/go-color"
	"fmt"
)

type DJI_GPX_Parser struct {
	fileName        string
}

func NewDJI_GPX_Parser(fileName string) *DJI_GPX_Parser {
	check := helpers.CheckFileFormat(fileName, ".gpx")
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
		helpers.PrintErrorLog("INVALID FILE. ERROR READING CONTENTS", err)
		return
	}

	gpxPath, err2 := filepath.Abs(parser.fileName)

	if err2 != nil {
		helpers.PrintErrorLog("INVALID FILE. ERROR READING CONTENTS", err)
		return
	}

	fmt.Println(color.Ize(color.Blue, string("File: " +  gpxPath + "\n")))

	fmt.Println(color.Ize(color.Blue, gpxFile.GetGpxInfo()))
}
