package helpers

import (
	"fmt"	
	"os"
	"github.com/TwiN/go-color"
	"strconv"
	"strings"
	"log"
)

func Option(min int, max int) int {
	fmt.Print("\n ENTER INPUT > ")
	var selection string
	fmt.Scanln(&selection)
	num, err := strconv.Atoi(selection)
    if err != nil {
		fmt.Println(color.Ize(color.Red, "  [!] INVALID INPUT"))
		return Option(min, max)
    } else {
		if (num == min) {
			fmt.Println(color.Ize(color.Blue, " Farewell and fly high!"))
			os.Exit(1)
			return 0
		} else if (num > min  && num < max + 1) {
			return num
		} else {
			fmt.Println(color.Ize(color.Red, "  [!] INVALID INPUT"))
			return Option(min, max)
		}
    }
}

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

func PrintError(message string) {
	fmt.Println(color.Ize(color.Red, "[ERROR] " + message))
}

func PrintErrorLog(message string, err error) {
	fmt.Println(color.Ize(color.Red, message))
	log.Println(color.Ize(color.Red, "[ERROR]"), err)
}