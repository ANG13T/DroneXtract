package steganographics

import (
	"github.com/TwiN/go-color"
	"fmt"
	"strings"
)

func GenTableHeader(name string) {
	fmt.Println(color.Ize(color.Blue, "\n    ╔══════════════════════════════════════════════════════════════════════════════╗"))
	var amount = (78 - len(name)) / 2
	fmt.Println(color.Ize(color.Blue, "    ║" +  strings.Repeat(" ", amount) + name + strings.Repeat(" ", amount + 1) + "║"))
	fmt.Println(color.Ize(color.Blue, "    ╠══════════════════════════════════════════════════════════════════════════════╣"))
}

func GenRowString(intro string, input string) {
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