package parsing

import (
	"io/ioutil"
	"github.com/TwiN/go-color"
	"fmt"
	"os"
	"io"
	// "bufio"
	"path/filepath"
	"unsafe"
	"log"
	"encoding/csv"
	"github.com/ANG13T/DroneXtract/models"
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
	outFn := "../output/output.csv"
 
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

	fieldnames := {
		"messageid", "offsetTime", "logDateTime", "time(millisecond)",
    "latitude", "longitude", "satnum", "gpsHealth", "altitude", "baroAlt", 
    "height", "accelX", "accelY", "accelZ", "accel", "gyroX", "gyroY", "gyroZ", "gyro", "errorX", "errorY", "errorZ", "error", "magX", "magY", "magZ", "magMod", 
    "velN", "velE", "velD", "vel", "velH", 
    "quatW", "quatX", "quatY", "quatZ", "roll", "pitch", "yaw", "yaw360", 
    "magYawX", "thrustAngle", "latitudeHP", "longitudeHP", 
    "imuTemp", "flyc_state", "flycStateStr", "nonGPSError", "nonGPSErrStr", 
    "DWflyCState", "connectedToRC", "current", "volt1", "volt2", "volt3", "volt4", "volt5", "volt6", "totalVolts", "voltSpread", "Watts", "batteryTemp(C)", "ratedCapacity", 
    "remainingCapacity", "percentageCapacity", "batteryUsefulTime", "voltagePercent", "batteryCycleCount", "batteryLifePercentage", "batteryBarCode", "minCurrent", "maxCurrent", 
    "avgCurrent", "minVolts", "maxVolts", "avgVolts", "minWatts", "maxWatts", "avgWatts", "Gimbal:roll", "Gimbal:pitch", "Gimbal:yaw", "Gimbal:Xroll", "Gimbal:Xpitch", "Gimbal:Xyaw", 
    "rFront", "lFront", "lBack", "rBack", 
    "rFrontSpeed", "lFrontSpeed", "lBackSpeed", "rBackSpeed", "rFrontLoad", "lFrontLoad", "lBackLoad", "rBackLoad", 
    "aileron", "elevator", "throttle", "rudder", "modeSwitch", "latitudeTablet", "longitudeTablet", "droneModel"
	}

	writer := csv.NewDictWriter(outFn)
	writer.Comma = ','
    writer.UseCRLF = false
	writer.Write(fieldnames)

	alternateStructure := false

	var b [1]byte
    _, err = path.Read(b[:])
    if err != nil {
        if err != io.EOF {
            panic(err)
        }
    }

    if b[0] != 0x55 {
        alternateStructure := true
        fmt.Println(alternateStructure)
    }

	kmlScale := 1
	kmlFile := path
	splPath := filepath.Split(inArg)
	fileName := splPath[len(splPath)-1]
	fileName = strings.Split(fileName, ".")[0]
	kmlFile = fileName + "-Map.kml"

	meta, err := os.Stat(inFn)
    if err != nil {
        panic(err)
    }

	message := models.NewMessage(meta, kmlFile, kmlScale)
	corruptPackets := 0
    unknownPackets := 0
    startIssue := true

	fmt.Println("done", message)


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