package steganography

import (
	"github.com/barasher/go-exiftool"
	"github.com/ANG13T/DroneXtract/forensics"
	"fmt"
	"os"
	"encoding/csv"
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

func (parser *DJI_EXIF_Parser) ExportToTXT(outputPath string) {
	if len(outputPath) == 0 {
		outputPath = "../output/exif-analysis.txt"
	}

	check := CheckFileFormat(outputPath, ".txt")
	if check == false {
		PrintError("INVALID OUTPUT FILE FORMAT. MUST BE TXT FILE")
		return
	}

	file, err := os.Create(outputPath)
	if err != nil {
		PrintErrorLog("FAILED TO CREATE TXT FILE", err)
		return
	}
	defer file.Close()

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

		for k, v := range fileInfo.Fields {
			formattedValue := fmt.Sprintf("%v", v)

			// Write the output to the file
			output := k + ":" + formattedValue + "\n"
			_, err = file.WriteString(output)

			
			if err != nil {
				PrintErrorLog("FAILED TO WRITE TO TXT FILE", err)
				return
			}
		}
	}

}	

func (parser *DJI_EXIF_Parser) ExportToCSV(outputPath string) {
	if len(outputPath) == 0 {
		outputPath = "../output/exif-analysis.csv"
	}

	check := CheckFileFormat(outputPath, ".csv")
	if check == false {
		PrintError("INVALID OUTPUT FILE FORMAT. MUST BE CSV FILE")
		return
	}

	file, err := os.Create(outputPath)
	if err != nil {
		PrintErrorLog("FAILED TO CREATE CSV FILE", err)
		return
	}
	defer file.Close()

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

	writer := csv.NewWriter(file)
	defer writer.Flush()

	fileInfos := et.ExtractMetadata(parser.fileName)

	for _, fileInfo := range fileInfos {
		if fileInfo.Err != nil {
			PrintErrorLog("COULD NOT READ FILE", fileInfo.Err)
			continue
		}

		for k, v := range fileInfo.Fields {
			formattedValue := fmt.Sprintf("%v", v)

			// Write the output to the file
			output := []string{k, formattedValue}
			err := writer.Write(output)

			
			if err != nil {
				PrintErrorLog("FAILED TO WRITE TO CSV FILE", err)
				return
			}
		}
	}
}

func (parser *DJI_EXIF_Parser) ExportToJSON(outputPath string) {

}

// export to csv, json