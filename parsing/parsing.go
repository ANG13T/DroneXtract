package parsing

import (
	"strings"
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
