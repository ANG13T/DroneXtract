package steganography

import (
	"github.com/barasher/go-exiftool"
	"github.com/ANG13T/DroneXtract/forensics"
	"fmt"
)

type DJI_EXIF_Parser struct {
	fileName        string
}


func NewDJI_EXIF_Parser(fileName string) *DJI_EXIF_Parser {
	parser := DJI_EXIF_Parser{
		fileName: fileName,
	}
	return &parser
}

func (parser *DJI_EXIF_Parser) Read() {

}


func (parser *DJI_EXIF_Parser) ExampleExiftool_Read() {
	et, err := exiftool.NewExiftool()
	if err != nil {
		fmt.Printf("Error when intializing: %v\n", err)
		if err == `error when executing command: exec: "exiftool.exe": executable file not found in %PATH%` {
			PrintError("EXIF TOOL NOT INSTALLED. VISIT https://exiftool.org/install.html FOR INSTRUCTIONS")
		}
		return
	}
	defer et.Close()

	fileInfos := et.ExtractMetadata("/Users/angelinatsuboi/Desktop/DJI-Forensics/dataset/DJI_0001.jpg")

	for _, fileInfo := range fileInfos {
		if fileInfo.Err != nil {
			fmt.Printf("Error concerning %v: %v\n", fileInfo.File, fileInfo.Err)
			continue
		}

		forensics.GenTableHeader("EXIF Analysis");
		for k, v := range fileInfo.Fields {
			formattedValue := fmt.Sprintf("%v", v)
			forensics.GenRowString(k, formattedValue)
		}
		forensics.GenTableFooter();
	}
}