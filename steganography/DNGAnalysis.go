package steganography

import (
	"image"
	"image/png"
	"os"

	_ "github.com/mdouchement/dng"
)

type DJI_DNG_Parser struct {
	fileName        string
}

func NewDJI_DNG_Parser(fileName string) *DJI_DNG_Parser {
	check := CheckFileFormat(fileName, "dng")
	if check == false {
		PrintError("INVALID FILE FORMAT. MUST BE DNG FILE")
		return nil
	}
	
	parser := DJI_DNG_Parser{
		fileName: fileName,
	}
	return &parser
}

func (parser *DJI_DNG_Parser) Read() { 
	exif := NewDJI_EXIF_Parser(parser.fileName)
	exif.Read()
}

func (parser *DJI_DNG_Parser) DNGtoPNG(outputFileName string) {
	fi, err := os.Open(parser.fileName)
	check("COULD NOT READ FILE", err)
	defer fi.Close()

	m, _, err := image.Decode(fi)
	check("CORRUPT IMAGE FILE", err)

	fo, err := os.Create(outputFileName)
	check("INVALID OUTPUT FILE", err)

	png.Encode(fo, m)
}

func (parser *DJI_DNG_Parser) ExportToTXT(outputPath string) { 
	exif := NewDJI_EXIF_Parser(parser.fileName)
	exif.ExportToTXT(outputPath)
}

func (parser *DJI_DNG_Parser) ExportToCSV(outputPath string) { 
	exif := NewDJI_EXIF_Parser(parser.fileName)
	exif.ExportToCSV(outputPath)
}

func (parser *DJI_DNG_Parser) ExportToJSON(outputPath string) { 
	exif := NewDJI_EXIF_Parser(parser.fileName)
	exif.ExportToJSON(outputPath)
}


// helper functions

func check(errorName string, err error) {
	if err != nil {
		PrintErrorLog(errorName, err)
	}
}