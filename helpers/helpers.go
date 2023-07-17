package helpers

import (
	"fmt"	
	"os"
	"github.com/TwiN/go-color"
	"strconv"
	"github.com/joho/godotenv"
	"path/filepath"
	"strings"
	"log"
)

func FileInputString() string {
	fmt.Print("\n ENTER FILE PATH > ")
	var selection string
	fmt.Scanln(&selection)
	return selection
}

func OutputPathString() string {
	fmt.Print("\n ENTER OUTPUT PATH > ")
	var selection string
	fmt.Scanln(&selection)
	return selection
}


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
		} else if (num == max - 1) {
			return -1
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

func CheckFileFormat(path string, exten string) bool {
	extension := strings.ToLower(filepath.Ext(path))
	return (extension == exten)
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


func GenTableFooter() {
	fmt.Println(color.Ize(color.Blue, "    ╚══════════════════════════════════════════════════════════════════════════════╝"))
}

func PrintLog(message string) {
	fmt.Println(color.Ize(color.Cyan, message))
}

func PrintError(message string) {
	fmt.Println(color.Ize(color.Red, "[ERROR] " + message))
}

func PrintErrorLog(message string, err error) {
	fmt.Println(color.Ize(color.Red, message))
	log.Println(color.Ize(color.Red, "[ERROR]"), err)
}

func PrintValidLog(message string) {
	fmt.Println(color.Ize(color.Green, message))
}

func PrintInvalidLog(message string) {
	fmt.Println(color.Ize(color.Red, message))
}

func GetEnvVariable(key string) int {
	err := godotenv.Load(".env")
  
	if err != nil {
		PrintError("FAILED TO LOAD .env FILE")
	}

	val, err2 := strconv.Atoi(os.Getenv(key))

	if err2 != nil {
		PrintError("INVALID ENV VARIABLE: " + key)
	}
  
	return val
}

func GetEnvVariances() []float64 {
	err := godotenv.Load(".env")
  
	if err != nil {
		PrintError("FAILED TO LOAD .env FILE")
	}

	return stringToFloat64Array(os.Getenv("ANALYSIS_MAX_VARIANCE"))
}

func stringToFloat64Array(input string) []float64 {
	strValues := strings.Fields(input)
	values := make([]float64, len(strValues))

	for i, strValue := range strValues {
		value, err := strconv.ParseFloat(strValue, 64)
		if err != nil {
			fmt.Printf("Error parsing value at index %d: %v\n", i, err)
			return nil
		}
		values[i] = value
	}

	return values
}