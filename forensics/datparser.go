package forensics

import (
	"io/ioutil"
	"github.com/TwiN/go-color"
	"fmt"
)


func DATParser() {
	options, _ := ioutil.ReadFile("txt/dat.txt")	
	fmt.Println(color.Ize(color.Cyan, "\n" + string(options) + "\n"))
	var selection int = Option(0, 4)

	if (selection == 1){
		ParseDatToCSV()
	} else if (selection == 2) {
		ParseDatToTXT()
	} else if (selection == 3) {
		ParseDatToKML()
	}

	return
}

func ParseDatToCSV() {

}

func ParseDatToTXT() {

}

func ParseDatToKML() {

}