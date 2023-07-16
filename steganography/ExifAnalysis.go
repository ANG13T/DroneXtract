package steganography

import (
	"github.com/barasher/go-exiftool"
	"github.com/ANG13T/DroneXtract/helpers"
	"fmt"
	"os"
	"encoding/csv"
	"encoding/json"
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

func (parser *DJI_EXIF_Parser) ExecuteEXIFAnalysis(input int) {
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

func (parser *DJI_EXIF_Parser) Read() {
	et, err := exiftool.NewExiftool()
	if err != nil {
		if err.Error() == `error when executing command: exec: "exiftool.exe": executable file not found in %PATH%` {
			helpers.PrintError("EXIF TOOL NOT INSTALLED. VISIT https://exiftool.org/install.html FOR INSTRUCTIONS")
		} else {
			helpers.PrintErrorLog("COULD NOT INITIALIZE EXIF TOOL", err)
		}
		return
	}
	defer et.Close()


	fileInfos := et.ExtractMetadata(parser.fileName)

	for _, fileInfo := range fileInfos {
		if fileInfo.Err != nil {
			helpers.PrintErrorLog("COULD NOT READ FILE", fileInfo.Err)
			continue
		}

		helpers.GenTableHeader("EXIF Analysis");
		for k, v := range fileInfo.Fields {
			formattedValue := fmt.Sprintf("%v", v)
			helpers.GenRowString(k, formattedValue)
		}
		helpers.GenTableFooter();
	}
}

func (parser *DJI_EXIF_Parser) ExportToTXT(outputPath string) {
	if len(outputPath) == 0 {
		outputPath = "../output/exif-analysis.txt"
	}

	check := helpers.CheckFileFormat(outputPath, ".txt")
	if check == false {
		helpers.PrintError("INVALID OUTPUT FILE FORMAT. MUST BE TXT FILE")
		return
	}

	file, err := os.Create(outputPath)
	if err != nil {
		helpers.PrintErrorLog("FAILED TO CREATE TXT FILE", err)
		return
	}
	defer file.Close()

	et, err := exiftool.NewExiftool()
	if err != nil {
		if err.Error() == `error when executing command: exec: "exiftool.exe": executable file not found in %PATH%` {
			helpers.PrintError("EXIF TOOL NOT INSTALLED. VISIT https://exiftool.org/install.html FOR INSTRUCTIONS")
		} else {
			helpers.PrintErrorLog("COULD NOT INITIALIZE EXIF TOOL", err)
		}
		return
	}
	defer et.Close()

	fileInfos := et.ExtractMetadata(parser.fileName)

	for _, fileInfo := range fileInfos {
		if fileInfo.Err != nil {
			helpers.PrintErrorLog("COULD NOT READ FILE", fileInfo.Err)
			continue
		}

		for k, v := range fileInfo.Fields {
			formattedValue := fmt.Sprintf("%v", v)

			// Write the output to the file
			output := k + ":" + formattedValue + "\n"
			_, err = file.WriteString(output)

			
			if err != nil {
				helpers.PrintErrorLog("FAILED TO WRITE TO TXT FILE", err)
				return
			}
		}
	}

}	

func (parser *DJI_EXIF_Parser) ExportToCSV(outputPath string) {
	if len(outputPath) == 0 {
		outputPath = "../output/exif-analysis.csv"
	}

	check := helpers.CheckFileFormat(outputPath, ".csv")
	if check == false {
		helpers.PrintError("INVALID OUTPUT FILE FORMAT. MUST BE CSV FILE")
		return
	}

	file, err := os.Create(outputPath)
	if err != nil {
		helpers.PrintErrorLog("FAILED TO CREATE CSV FILE", err)
		return
	}
	defer file.Close()

	et, err := exiftool.NewExiftool()
	if err != nil {
		if err.Error() == `error when executing command: exec: "exiftool.exe": executable file not found in %PATH%` {
			helpers.PrintError("EXIF TOOL NOT INSTALLED. VISIT https://exiftool.org/install.html FOR INSTRUCTIONS")
		} else {
			helpers.PrintErrorLog("COULD NOT INITIALIZE EXIF TOOL", err)
		}
		return
	}
	defer et.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	fileInfos := et.ExtractMetadata(parser.fileName)

	for _, fileInfo := range fileInfos {
		if fileInfo.Err != nil {
			helpers.PrintErrorLog("COULD NOT READ FILE", fileInfo.Err)
			continue
		}

		for k, v := range fileInfo.Fields {
			formattedValue := fmt.Sprintf("%v", v)

			// Write the output to the file
			output := []string{k, formattedValue}
			err := writer.Write(output)

			
			if err != nil {
				helpers.PrintErrorLog("FAILED TO WRITE TO CSV FILE", err)
				return
			}
		}
	}
}

func (parser *DJI_EXIF_Parser) ExportToJSON(outputPath string) {
	if len(outputPath) == 0 {
		outputPath = "../output/exif-analysis.json"
	}

	check := helpers.CheckFileFormat(outputPath, ".json")
	if check == false {
		helpers.PrintError("INVALID OUTPUT FILE FORMAT. MUST BE JSON FILE")
		return
	}

	file, err := os.Create(outputPath)
	if err != nil {
		helpers.PrintErrorLog("FAILED TO CREATE JSON FILE", err)
		return
	}
	defer file.Close()

	et, err := exiftool.NewExiftool()
	if err != nil {
		if err.Error() == `error when executing command: exec: "exiftool.exe": executable file not found in %PATH%` {
			helpers.PrintError("EXIF TOOL NOT INSTALLED. VISIT https://exiftool.org/install.html FOR INSTRUCTIONS")
		} else {
			helpers.PrintErrorLog("COULD NOT INITIALIZE EXIF TOOL", err)
		}
		return
	}
	defer et.Close()

	fileInfos := et.ExtractMetadata(parser.fileName)

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(fileInfos[0])

	if err != nil {
		helpers.PrintErrorLog("FAILED TO ENCODE JSON", err)
		return
	}
}
