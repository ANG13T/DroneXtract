package forensics

import (
	"io/ioutil"
	"github.com/TwiN/go-color"
	"fmt"
	"os"
	"io"
	// "bufio"
	"unsafe"
	"log"
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
	fmt.Print("\n ENTER DAT FILE PATH > ")
	var path string
	fmt.Scanln(&path)
	_, err := os.Open(path)
	// TODO: change this
	force := false
	//outFn := "../output/output.csv"
 
	if err != nil {
		fmt.Println(color.Ize(color.Red, "  [!] INVALID DAT FILE"))
	}

	inFile, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer inFile.Close()

	var fileHeader [128]byte
	_, err = inFile.Read(fileHeader[:])
	if err != nil {
		log.Fatal(err)
	}

	buildBytes := *(*[5]byte)(unsafe.Pointer(&fileHeader[16])) // Read 5 bytes starting from offset 16
	build := string(buildBytes[:])

	if build != "BUILD" {
		if !force {
			err := NotDATFileError{filename: path}
			log.Fatal(err)
		} else {
			fmt.Printf("*** WARNING: %s is not a recognized DJI DAT file but will be processed anyway because the FORCE flag was set. ***\n", path)
			_, err = inFile.Seek(0, io.SeekStart) // set the pointer to the beginning of the file because this is an unrecognized file type and we don't want to risk missing data
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	// outFile, err := os.OpenFile(outFn, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer outFile.Close()

	// var byteValue [1]byte
	// _, err = io.ReadFull(inFile, byteValue[:])
	// if err != nil {
	// 	fmt.Println(color.Ize(color.Red, "  [!] INVALID DAT FILE"))
	// }

	// if byteValue[0] != 0x55 {
	// 	alternateStructure := true
	// }
	
}

func ParseDatToTXT() {

}

func ParseDatToKML() {

}