package parsing

import (
	"strings"
	"path/filepath"
	"github.com/TwiN/go-color"
	"fmt"
	"github.com/ANG13T/DroneXtract/helpers"
)

func ExecuteParser(index int) {
	filePath := helpers.FileInputString()
	switch in := index; in {
		case 1:
			suite := NewDJI_CSV_Parser(filePath)
			suite.ParseContents()
		case 2:
			suite := NewDJI_KML_Parser(filePath)
			suite.ParseContents()
		case 3:
			suite := NewDJI_GPX_Parser(filePath)
			suite.ParseContents()
	}
}

func CheckFileFormat(path string, exten string) bool {
	extension := strings.ToLower(filepath.Ext(path))
	return (extension == exten)
}

func GenTableHeader(name string, containBreak bool) {
	if containBreak {
		fmt.Println(color.Ize(color.Blue, "\n    ╔══════════════════════════════════════════════════════════════════════════════╗"))
	} else {
		fmt.Println(color.Ize(color.Blue, "    ╔══════════════════════════════════════════════════════════════════════════════╗"))
	}
	var amount = (78 - len(name)) / 2
	var extraPadding = 1
	if len(name) % 2 == 0 {
		extraPadding = 0
	}
	fmt.Println(color.Ize(color.Blue, "    ║" +  strings.Repeat(" ", amount) + name + strings.Repeat(" ", amount + extraPadding) + "║"))
	fmt.Println(color.Ize(color.Blue, "    ╠══════════════════════════════════════════════════════════════════════════════╣"))
}

func GenTableHeaderModified(name string) {
	fmt.Println(color.Ize(color.Blue, "    ╠══════════════════════════════════════════════════════════════════════════════╣"))
	var amount = (78 - len(name)) / 2
	var extraPadding = 1
	if len(name) % 2 == 0 {
		extraPadding = 0
	}
	fmt.Println(color.Ize(color.Blue, "    ║" +  strings.Repeat(" ", amount) + name + strings.Repeat(" ", amount + extraPadding) + "║"))
	fmt.Println(color.Ize(color.Blue, "    ╠══════════════════════════════════════════════════════════════════════════════╣"))
}

func GenRowString(intro string, input string) {
	if input == "UNSPECIFIED" {
		return
	}
	var totalCount int = 4 + len(intro) + len(input) + 2
	var useCount = 80 - totalCount
	if useCount < 0 { 
		useCount = 0
	}
	var val = "    ║ " + intro + ": " + input + strings.Repeat(" ", useCount) + " ║"
	fmt.Println(color.Ize(color.Blue, val))
}

func GenTableFooter() {
	fmt.Println(color.Ize(color.Blue, "    ╚══════════════════════════════════════════════════════════════════════════════╝"))
}