package steganography

import (
	"github.com/ANG13T/DroneXtract/helpers"
)

type DJI_XMP_Parser struct {
	fileName        string
}

func NewDJI_XMP_Parser(fileName string) *DJI_XMP_Parser {
	check := CheckFileFormat(fileName, ".xmp")
	if check == false {
		PrintError("INVALID FILE FORMAT. MUST BE XMP FILE")
		return nil
	}
	
	parser := DJI_XMP_Parser{
		fileName: fileName,
	}
	return &parser
}

func (parser *DJI_XMP_Parser) ExecuteXMPAnalysis(input int) {
	switch in := input; in {
		case 1:
			parser.Read()
		case 2:
			outputPath := helpers.FileInputString()
			parser.ExportToTXT(outputPath)
		case 3:
			outputPath := helpers.FileInputString()
			parser.ExportToCSV(outputPath)
		case 4:
			outputPath := helpers.FileInputString()
			parser.ExportToJSON(outputPath)
	}
}

func (parser *DJI_XMP_Parser) Read() { 
	exif := NewDJI_EXIF_Parser(parser.fileName)
	exif.Read()
}

func (parser *DJI_XMP_Parser) ExportToTXT(outputPath string) { 
	exif := NewDJI_EXIF_Parser(parser.fileName)
	exif.ExportToTXT(outputPath)
}

func (parser *DJI_XMP_Parser) ExportToCSV(outputPath string) { 
	exif := NewDJI_EXIF_Parser(parser.fileName)
	exif.ExportToCSV(outputPath)
}

func (parser *DJI_XMP_Parser) ExportToJSON(outputPath string) { 
	exif := NewDJI_EXIF_Parser(parser.fileName)
	exif.ExportToJSON(outputPath)
}
