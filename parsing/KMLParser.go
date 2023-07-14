package parsing

import (
	"github.com/ANG13T/DroneXtract/helpers"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
)

type DJI_KML_Parser struct {
	fileName        string
}

type KML struct {
	Placemarks []Placemark `xml:"Document>Placemark"`
}

type Placemark struct {
	Name        string       `xml:"name"`
	ExtendedData ExtendedData `xml:"ExtendedData"`
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
	check := CheckFileFormat(fileName, ".kml")
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
		log.Fatal(err)
	}

	// Parse the KML data
	var kml KML
	err = xml.Unmarshal(content, &kml)
	if err != nil {
		log.Fatal(err)
	}

	// Process the placemarks
	for _, placemark := range kml.Placemarks {
		fmt.Println("Name:", placemark.Name)

		if len(placemark.ExtendedData.Data) > 0 {
			fmt.Println("Extended Data:")
			for _, data := range placemark.ExtendedData.Data {
				fmt.Println("  - Name:", data.Name)
				fmt.Println("    Value:", data.Value)
			}
		}

		if placemark.Point.Coordinates != "" {
			fmt.Println("Point Coordinates:", placemark.Point.Coordinates)
		}

		if placemark.LineString.Coordinates != "" {
			fmt.Println("LineString Coordinates:", placemark.LineString.Coordinates)
		}

		fmt.Println()
	}
}