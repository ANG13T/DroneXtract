package steganographics

import (
	"github.com/TwiN/go-color"
	"fmt"
	"strings"
	"log"
)

func GenTableHeader(name string) {
	fmt.Println(color.Ize(color.Blue, "\n    ╔══════════════════════════════════════════════════════════════════════════════╗"))
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

func PrintError(message string) {
	fmt.Println(color.Ize(color.Red, "[ERROR] " + message))
}

func PrintErrorLog(message string, err error) {
	fmt.Println(color.Ize(color.Red, message))
	log.Println(color.Ize(color.Red, "[ERROR]"), err)
}