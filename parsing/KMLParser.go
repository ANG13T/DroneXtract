package parsing

import (
	"github.com/ANG13T/DroneXtract/helpers"
	"encoding/xml"
	"io/ioutil"
	"strconv"
	"strings"
)

type DJI_KML_Parser struct {
	fileName        string
}

type KML struct {
	Placemarks []Placemark `xml:"Document>Placemark"`
}

type Placemark struct {
	Name        string       `xml:"name"`
	Point       Point        `xml:"Point"`
	LineString  LineString   `xml:"LineString"`
}

type ExtendedData struct {
	Data []Data `xml:"Data"`
}

type Data struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value"`
}

type Point struct {
	Coordinates string `xml:"coordinates"`
}

type LineString struct {
	Coordinates string `xml:"coordinates"`
}


func NewDJI_KML_Parser(fileName string) *DJI_KML_Parser {
	check := helpers.CheckFileFormat(fileName, ".kml")
	if check == false {
		helpers.PrintError("INVALID FILE FORMAT. MUST BE KML FILE")
		return nil
	}

	parser := DJI_KML_Parser{
		fileName: fileName,
	}
	return &parser
}

func (parser *DJI_KML_Parser) ParseContents() {
	content, err := ioutil.ReadFile(parser.fileName)
	if err != nil {
		helpers.PrintErrorLog("INVALID FILE. ERROR READING CONTENTS", err)
	}

	// Parse the KML data
	var kml KML
	err = xml.Unmarshal(content, &kml)
	if err != nil {
		helpers.PrintErrorLog("INVALID FILE. ERROR READING CONTENTS", err)
	}

	// Process the placemarks
	for _, placemark := range kml.Placemarks {

		if placemark.Point.Coordinates != "" {
			coorValues := strings.Split(placemark.Point.Coordinates, ",")
			GenTableHeader("Home Point Information", true)
			helpers.GenRowString("Coordinates", "(" + coorValues[0] + "," + coorValues[1] + ")")
			helpers.GenRowString("Altitude", coorValues[2] + " ft")
			helpers.GenTableFooter()
		}

		if placemark.LineString.Coordinates != "" {

			lines := strings.Split(placemark.LineString.Coordinates, "\n")

			for cIndex, coor := range lines {
				coorValues := strings.Split(coor, ",")
				if len(coor) > 0 && len(coorValues) > 1 {
					GenTableHeader("Coordinate Point " + strconv.Itoa(cIndex + 1), false)
					helpers.GenRowString("Coordinates", "(" + coorValues[0] + ", " + coorValues[1] + ")")
					helpers.GenRowString("Altitude", coorValues[2] + " ft")
					helpers.GenTableFooter()
				}
			}
		}
	}
}