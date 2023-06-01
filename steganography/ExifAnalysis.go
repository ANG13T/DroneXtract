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

// export to txt, csv, and json

func (parser *DJI_EXIF_Parser) Read() {
	et, err := exiftool.NewExiftool()
	if err != nil {
		if err.Error() == `error when executing command: exec: "exiftool.exe": executable file not found in %PATH%` {
			PrintError("EXIF TOOL NOT INSTALLED. VISIT https://exiftool.org/install.html FOR INSTRUCTIONS")
		} else {
			PrintErrorLog("COULD NOT INITIALIZE EXIF TOOL", err)
		}
		return
	}
	defer et.Close()


	fileInfos := et.ExtractMetadata(parser.fileName)

	for _, fileInfo := range fileInfos {
		if fileInfo.Err != nil {
			PrintErrorLog("COULD NOT READ FILE", fileInfo.Err)
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