package forensics

import (
	"io/ioutil"
	"github.com/iskaa02/qalam/gradient"
)


func DATParser() {
	options, _ := ioutil.ReadFile("txt/dat.txt")
	opt,_:=gradient.NewGradient("#1179ef", "cyan")
	opt.Print("\n" + string(options))
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