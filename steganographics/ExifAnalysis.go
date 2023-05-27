package steganographics

import (
	"github.com/barasher/go-exiftool"
	"github.com/ANG13T/DroneXtract/forensics"
	"fmt"
)


func ExampleExiftool_Read() {
	et, err := exiftool.NewExiftool()
	if err != nil {
		fmt.Printf("Error when intializing: %v\n", err)
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