package parsing

import (
	"github.com/ANG13T/DroneXtract/helpers"
	"github.com/tkrajina/gpxgo/gpx"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
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

	gpxPath, _ := filepath.Abs(gpxFileArg)

	fmt.Print("File: ", gpxPath, "\n")

	fmt.Println(gpxFile.GetGpxInfo())
}

func formatXML(content []byte) ([]byte, error) {
	// Create a buffer to store the formatted XML
	formatted := &bytes.Buffer{}

	// Indent the XML content
	err := xml.Indent(formatted, content, "", "  ")
	if err != nil {
		return nil, err
	}

	return formatted.Bytes(), nil
}