package models

import (
	"os"
	"github.com/twpayne/go-kml/v3"
	"time"
)

type Message struct {
	fieldnames    []string
	tickNo        *uint32
	tickOffset    int
	rowOut        map[string]interface{}
	packetNum     int
	packets       []*Packet
	addedData     bool
	meta          os.FileInfo
	startUNIXTime int64
	gpsFrDict     map[uint32][]interface{}
	kmlFile       *os.File
	kmlWriter     *kml.KMLElement
	kmlRes        int
	pointCnt      int
}

func (m *Message) NewMessage(meta os.FileInfo, kmlFile *os.File, kmlScale int) {
	m.fieldnames = []string{
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
    "aileron", "elevator", "throttle", "rudder", "modeSwitch", "latitudeTablet", "longitudeTablet", "droneModel"}
	m.tickNo = nil
	m.tickOffset = 0
	m.rowOut = make(map[string]interface{})
	m.packetNum = 0
	m.packets = []*Packet{}
	m.addedData = false
	m.meta = meta
	m.gpsFrDict = make(map[uint32][]interface{})

	m.kmlFile = kmlFile
	m.kmlWriter = nil
	if m.kmlFile != nil {
		m.kmlWriter = kml.KML(kml.Placemark(
            kml.Name("Simple placemark"),
            kml.Description("Attached to the ground. Intelligently places itself at the height of the underlying terrain."),

        ))
		// m.kmlRes = kmlScale
	}

	m.startUNIXTime = time.Now().Unix()
}

