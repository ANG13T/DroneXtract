package steganography

// TODO -

// toGeoJSON
// 0 - support all files
// 5 - subtitle extractor from videos (MP4 to SRT)
// 7 - comments

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"io/ioutil"
	"os"
	"encoding/json"
	"encoding/csv"
)

var isoDateRegex = regexp.MustCompile(`[0-9]{4}-[0-9]{2}-[0-9]{2}T[0-9]{2}:[0-9]{2}:[0-9]{2}(\.[0-9]+)?Z`)

type DJI_SRT_Parser struct {
	fileName        string
	packets			[]SRT_Packet
}

func NewDJI_SRT_Parser(fileName string) *DJI_SRT_Parser {
	check := CheckFileFormat(fileName, ".srt")
	if check == false {
		PrintError("INVALID FILE FORMAT. MUST BE SRT FILE")
		return nil
	}

	parser := DJI_SRT_Parser{
		fileName: fileName,
		packets:            make([]SRT_Packet, 0),
	}
	return &parser
}


type SRT_Packet struct {
	frame_count string
	diff_time    string
	iso 		string
	shutter 	string
	fnum 		string
	ev 			string
	ct			string
	color_md	string
	focal_len 	string
	latitude 	string
	longtitude	string
	altitude	string
	date 		string
	time_stamp	string
	barometer 	string
}

type GeoFeatureJSONPointResult struct {
	Type       string                 `json:"type"`
	Geometry   PointGeometry          `json:"geometry"`
	Properties GeoProperty            `json:"properties"`
}

type GeoFeatureJSONLineStringResult struct {
	Type       string                 `json:"type"`
	Geometry   LineStringGeometry     `json:"geometry"`
	Properties GeoJSONEnding          `json:"properties"`
}

type GeoJSONResult struct {
	Type       string                 `json:"type"`
	CRS   	   CRS                    `json:"crs"`
	Features   []interface{} 		  `json:"features"`
}

type PointGeometry struct {
	Type        string      `json:"type"`
	Coordinates []float64 	`json:"coordinates"`
}

type LineStringGeometry struct {
	Type        string      `json:"type"`
	Coordinates [][]float64 	`json:"coordinates"`
}

type GeoJSONEnding struct {
	Source        string      `json:"source"`
	Timestamp     []int64     `json:"timestamp"`
	Name          string      `json:"name"`
	HomeLatitude  float64     `json:"HOME_LATITUDE"`
	HomeLongitude float64     `json:"HOME_LONGITUDE"`
	ISO           int32       `json:"ISO"`
	Shutter       string      `json:"SHUTTER"`
	FNUM          int32       `json:"FNUM"`
}

type GPS struct {
	Latitude       float64           `json:"latitude"`
	Longitude      float64           `json:"longitude"`
	Altitude   	   float64  		   `json:"altitude"`
}

type GPSPoint struct {
	Latitude 	   float64			`json:"latitude"`
	Longitude      float64          `json:"longitude"`
}

type JSONResult struct {
	FileName 	 string			`json:"fileName"`
	Packets      []GeoProperty  `json:"packets"`
}

type GeoProperty struct {
	FrameCount 	int64	`json:"FRAMECOUNT"`
	DiffTime   string	`json:"DIFF_TIME"`
	ISO 		int32	`json:"ISO"`
	Shutter     string	`json:"SHUTTER"`
	FNUM		int32	`json:"FNUM"`
	EV			int32	`json:"EV"`
	CT 			int64	`json:"CT"`
	ColorMd 	string	`json:"COLOR_MD"`
	FocalLen	int32	`json:"FOCAL_LEN"`
	Latitude	float64	`json:"LATITUDE"`
	Longitude	float64	`json:"LONGITUDE"`
	Altitude	float64	`json:"ALTITUDE"`
	Date 		string	`json:"DATE"`
	Timestamp  string	`json:"timestamp"`
	Barometer   float64	`json:"BAROMETER"`
}


type CRS struct {
	Type        string    `json:"type"`
	Properties map[string]interface{} `json:"properties"`
}


func (packet *SRT_Packet) printSRTPacket(length string) {
	title := "FRAME " + checkEmptyField(packet.frame_count)
	if packet.frame_count == "1" {
		GenTableHeader(title, true)
	} else {
		GenTableHeaderModified(title)
	}
	
	GenRowString("FRAME COUNT", checkEmptyField(packet.frame_count))
	GenRowString("DIFF TIME", checkEmptyField(packet.diff_time))
	GenRowString("ISO", checkEmptyField(packet.iso))
	GenRowString("SHUTTER", checkEmptyField(packet.shutter))
	GenRowString("FNUM", checkEmptyField(packet.fnum))
	GenRowString("EV", checkEmptyField(packet.ev))
	GenRowString("CT", checkEmptyField(packet.ct))
	GenRowString("COLOR MD", checkEmptyField(packet.color_md))
	GenRowString("FOCAL EN", checkEmptyField(packet.focal_len))
	GenRowString("LATITUDE", checkEmptyField(packet.latitude))
	GenRowString("LONGITUDE", checkEmptyField(packet.longtitude))
	GenRowString("ALTITUDE", checkEmptyField(packet.altitude))
	GenRowString("DATE", checkEmptyField(packet.date))
	GenRowString("TIME STAMP", checkEmptyField(packet.time_stamp))
	GenRowString("BAROMETER", checkEmptyField(packet.barometer))
	
	if packet.frame_count == length { 
		GenTableFooter()
	} 

}

func (parser *DJI_SRT_Parser) SRTToObject(srt string) []SRT_Packet {
	converted := make([]SRT_Packet, 0)
	test_regex := regexp.MustCompile(`\[(\w+)\s*:\s*([^]]+)\]`)
	test_regex_2 := regexp.MustCompile(`([A-Za-z]+):([-.\w/]+)`)
	diffTimeRegex := regexp.MustCompile(`\bDiffTime\s*:\s*([^ ]+)`)
	timecodeRegEx := regexp.MustCompile(`(\d{2}:\d{2}:\d{2},\d{3})\s-->\s`)
	packetRegEx := regexp.MustCompile(`^\d+$`)
	GPSRegEx := regexp.MustCompile(`GPS\(([-.\d]+,[-.\d]+,[-.\d]+)\)`)
	GPSRegEx2 := regexp.MustCompile(`GPS\((-?\d+\.\d+),(-?\d+\.\d+),(-?\d+\.\d+)M\)`)
	//arrayRegEx := regexp.MustCompile(`\b([A-Z_a-z]+)\(([-\+\w.,/]+)\)`)
	dateRegEx := regexp.MustCompile(`\d{4}[-.]\d{1,2}[-.]\d{1,2} \d{1,2}:\d{2}:\d{2,}`)
	accurateDateRegex := regexp.MustCompile(`(\d{4}[-.]\d{1,2}[-.]\d{1,2} \d{1,2}:\d{2}:\d{2}),(\w{3}),(\w{3})`)
	accurateDateRegex2 := regexp.MustCompile(`(\d{4}[-.]\d{1,2}[-.]\d{1,2} \d{1,2}:\d{2}:\d{2})[,.](\w{3})`)

	isDJIFPV := regexp.MustCompile(`font size="28"`).MatchString(srt) &&
		regexp.MustCompile(`\d{4}-\d{1,2}-\d{1,2} \d{1,2}:\d{2}:\d{2}.\d{3}`).MatchString(srt) &&
		regexp.MustCompile(`\[altitude: \d.*\]`).MatchString(srt)

	// Split difficult Phantom4Pro format
	srt = regexp.MustCompile(`.*-->.*`).ReplaceAllStringFunc(srt, func(match string) string {
		return strings.ReplaceAll(match, ",", ":separator:")
	})
	srt = regexp.MustCompile(`\(([^\)]+)\)`).ReplaceAllStringFunc(srt, func(match string) string {
		match = strings.ReplaceAll(match, ",", ":separator:")
		match = strings.ReplaceAll(match, " ", "")
		return match
	})
	srt = strings.ReplaceAll(srt, ", ", " ")
	srt = strings.ReplaceAll(srt, "Â", "")
	srt = strings.ReplaceAll(srt, "°", "")
	srt = strings.ReplaceAll(srt, "B0", "")
	srt = strings.ReplaceAll(srt, ":separator:", ",")

	// Split others
	lines := strings.Split(srt, "\n")
	for i := range lines {
		lines[i] = strings.TrimSpace(lines[i])
		lines[i] = regexp.MustCompile(`([a-zA-Z])\s([-\d])`).ReplaceAllString(lines[i], "$1:$2")
		lines[i] = regexp.MustCompile(`([a-zA-Z])\s\(`).ReplaceAllString(lines[i], "$1(")
		lines[i] = regexp.MustCompile(`([a-zA-Z])\.([a-zA-Z])`).ReplaceAllString(lines[i], "$1_$2")
		lines[i] = regexp.MustCompile(`([a-zA-Z])\/(\d)`).ReplaceAllString(lines[i], "$1:$2")
	}

	lines = filterEmptyLines(lines) 

	for _, line := range lines {
		var match []string
		matched := packetRegEx.MatchString(fmt.Sprintf("%s", line))

		if matched || len(line) < 5{
			// new packet
			converted = append(converted, SRT_Packet{})
			converted[len(converted)-1].frame_count = line
		} else if match = timecodeRegEx.FindStringSubmatch(line); match != nil {
			values := strings.Split(match[1], ",")
			converted[len(converted)-1].time_stamp = values[0]
		} else {
			matches_2 := test_regex.FindAllStringSubmatch(line, -1)

			matches_3 := test_regex_2.FindAllStringSubmatch(line, -1)

			properties := make(map[string]string)
			selectedArr := matches_2

			if len(selectedArr) == 0 {
				selectedArr = matches_3
			}

			for _, match := range selectedArr {
				if len(match) == 3 {
					key := match[1]
					value := match[2]
					properties[key] = value
				}
			}

			// Print the extracted property-value pairs
			for key, value := range properties {
				switch strings.ToLower(key) {
				case "iso":
					converted[len(converted)-1].iso = value
				case "shutter":
					converted[len(converted)-1].shutter = value
				case "fnum":
					converted[len(converted)-1].fnum = value
				case "ev":
					converted[len(converted)-1].ev = value
				case "ct":
					converted[len(converted)-1].ct = value
				case "color_md":
					converted[len(converted)-1].color_md = value
				case "focal_len":
					converted[len(converted)-1].focal_len = value
				case "latitude":
					converted[len(converted)-1].latitude = value
				case "longtitude":
					converted[len(converted)-1].longtitude = value
				case "barometer":
					converted[len(converted)-1].barometer = value
				case "altitude":
					// Correct altitude divided by 10 problem in DJI FPV drone
					if isDJIFPV {
						alt, _ := strconv.Atoi(value)
						converted[len(converted)-1].altitude = strconv.Itoa(alt * 10)
					} else {
						converted[len(converted)-1].altitude = value
					}	
				}
			}

			diff_match := diffTimeRegex.FindStringSubmatch(line)

			if len(diff_match) == 2 {
				converted[len(converted)-1].diff_time = diff_match[1]
			}

			if match = accurateDateRegex.FindStringSubmatch(line); match != nil {
				converted[len(converted)-1].date = match[1] + ":" + match[2] + "." + match[3]
			} else if match = accurateDateRegex2.FindStringSubmatch(line); match != nil {
				converted[len(converted)-1].date = match[1] + "." + match[2]
			} else if match = dateRegEx.FindStringSubmatch(line); match != nil {
				converted[len(converted)-1].date = match[0]
			}

			match := GPSRegEx.FindStringSubmatch(line)

			match2 := GPSRegEx2.FindStringSubmatch(line)

			if len(match) > 1 {
				gpsValue := match[1]
				gpsVals := strings.Split(gpsValue, ",")
				converted[len(converted)-1].latitude = gpsVals[0]
				converted[len(converted)-1].longtitude = gpsVals[1]
				converted[len(converted)-1].altitude = gpsVals[2]
			}

			if len(match2) > 1 {
				converted[len(converted)-1].latitude = match2[1]
				converted[len(converted)-1].longtitude = match2[2]
				converted[len(converted)-1].altitude = match2[3]
			}
		}
	}

	if len(converted) < 1 || (len(converted) == 1 && checkNullPacket(converted[0])) {
		PrintError("ERROR PARSING SRT FILE")
		return nil
	}

	return converted
}

func (parser *DJI_SRT_Parser) GeneratePackets() {
	// Check if Valid File Path
	content, err := ioutil.ReadFile(parser.fileName)

	if err != nil {
		PrintErrorLog("INVALID FILE PATH", err)
	}

	string_content := string(content)

	if checkValidFileContents(string_content) {
		parser.packets = parser.SRTToObject(string_content)
	} 
}

func (parser *DJI_SRT_Parser) PrintAllPackets() {
	for _, packet := range parser.packets {
		amount := len(parser.packets) 
		packet.printSRTPacket(strconv.Itoa(amount))
	}
}



func (parser *DJI_SRT_Parser) PrintFileMetadata() {
	file, err := os.Open(parser.fileName)
	if err != nil {
		PrintErrorLog("INVALID FILE", err)
		return
	}
	defer file.Close()

	// Get file information
	fileInfo, err := file.Stat()
	if err != nil {
		PrintErrorLog("UNABLE TO OBTAIN FILE METADATA", err)
		return
	}

	fileName := fileInfo.Name()
	fileSize := fileInfo.Size()
	modTime := fileInfo.ModTime()

	GenTableHeader("Parsing SRT Job", true)

	// Print metadata
	GenRowString("File Name", fileName)
	GenRowString("File Size (bytes)", strconv.FormatInt(fileSize, 10))
	GenRowString("Last Modified Time", modTime.Format("2006-01-02 15:04:05"))

	GenTableFooter()
}


func (parser *DJI_SRT_Parser) ExportToJSON(outputPath string) {
	if len(outputPath) == 0 {
		outputPath = "../output/srt-analysis.json"
	}

	check := CheckFileFormat(outputPath, ".json")
	if check == false {
		PrintError("INVALID OUTPUT FILE FORMAT. MUST BE JSON FILE")
		return
	}

	file, err := os.Create(outputPath)
	if err != nil {
		PrintErrorLog("FAILED TO CREATE JSON FILE", err)
		return
	}
	defer file.Close()

	result := JSONResult{
		FileName: parser.fileName,
		Packets: []GeoProperty{},
	}

	for _, packet := range parser.packets {
		conv := parser.PacketToGeoFeatureJSON(packet)
		result.Packets = append(result.Packets, conv.Properties)
	}

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(result)

	if err != nil {
		PrintErrorLog("FAILED TO ENCODE GEOJSON", err)
		return
	}

}

func (parser *DJI_SRT_Parser) ExportToGeoJSON(outputPath string) {
	if len(outputPath) == 0 {
		outputPath = "../output/srt-analysis.geojson"
	}

	check := CheckFileFormat(outputPath, ".geojson")
	if check == false {
		PrintError("INVALID OUTPUT FILE FORMAT. MUST BE GEOJSON FILE")
		return
	}

	file, err := os.Create(outputPath)
	if err != nil {
		PrintErrorLog("FAILED TO CREATE GEOJSON FILE", err)
		return
	}
	defer file.Close()

	result := GeoJSONResult{
		Type: "FeatureCollection",
		CRS: CRS{
			Type: "name",
			Properties: map[string]interface{}{
				"name": "urn:ogc:def:crs:OGC:1.3:CRS84",
			},
		},
		Features: nil,
	}

	initial_packet := parser.PacketToGeoFeatureJSON(parser.packets[0]).Properties
	
	geo_features := GeoJSONEnding{
		Source: "dji-srt-parser",
		Timestamp: []int64{},
		Name: parser.fileName,
		HomeLatitude: initial_packet.Latitude,
		HomeLongitude: initial_packet.Longitude,
		ISO: initial_packet.ISO,
		Shutter: initial_packet.Shutter,
		FNUM: initial_packet.FNUM,
	}

	ending := GeoFeatureJSONLineStringResult{
		Type: "Feature",
		Properties: geo_features,
		Geometry: LineStringGeometry{
			Type: "LineString",
			Coordinates: [][]float64{},
		},
	}

	for _, packet := range parser.packets {
		conv := parser.PacketToGeoFeatureJSON(packet)
		convGeoProp := conv.Properties
		result.Features = append(result.Features, conv)
		endData := ending.Properties
		endData.Timestamp = append(endData.Timestamp, strToInt64(packet.time_stamp))
		geoArr := []float64{
			convGeoProp.Latitude, 
			convGeoProp.Longitude, 
			convGeoProp.Altitude,
		}
		ending.Geometry.Coordinates = append(ending.Geometry.Coordinates, geoArr)
	}

	result.Features = append(result.Features, ending)

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(result)

	if err != nil {
		PrintErrorLog("FAILED TO ENCODE GEOJSON", err)
		return
	}

}

func (parser *DJI_SRT_Parser) ExporttoMGJSON(outputPath string) {
	if len(outputPath) == 0 {
		outputPath = "../output/srt-analysis.mgjson"
	}

	check := CheckFileFormat(outputPath, ".mgjson")
	if check == false {
		PrintError("INVALID OUTPUT FILE FORMAT. MUST BE MGJSON FILE")
		return
	}

	file, err := os.Create(outputPath)
	if err != nil {
		PrintErrorLog("FAILED TO CREATE MGJSON FILE", err)
		return
	}
	defer file.Close()
}

func (parser *DJI_SRT_Parser) ExportToCSV(outputPath string) {
	if len(outputPath) == 0 {
		outputPath = "../output/srt-analysis.csv"
	}

	check := CheckFileFormat(outputPath, ".csv")
	if check == false {
		PrintError("INVALID OUTPUT FILE FORMAT. MUST BE CSV FILE")
		return
	}

	file, err := os.Create(outputPath)
	if err != nil {
		PrintErrorLog("FAILED TO CREATE CSV FILE", err)
		return
	}
	defer file.Close()

	data := [][]string{
		{"TIMESTAMP", "HOME.LATITUDE", "HOME.LONGITUDE", "DATE", "GPS.LATITUDE", "GPS.LONGITUDE", "GPS.ALTITUDE", "BAROMETER", "FRAME COUNT", "DIFF TIME", "ISO", "SHUTTER", "FNUM", "EV", "CT", "COLOR MD", "FOCAL LEN", "NAME"},
	}

	writer := csv.NewWriter(file)

	err = writer.WriteAll(data)
	if err != nil {
		PrintErrorLog("ERROR WRITING TO CSV FILE", err)
	}

	writer.Flush()

	if err := writer.Error(); err != nil {
		PrintErrorLog("ERROR WRITING TO CSV FILE", err)
	}

}

// Helpers

func isNum(d string) bool {
	_, err := strconv.ParseFloat(d, 64)
	return err == nil
}

func (parser *DJI_SRT_Parser) PacketToGeoFeatureJSON(packet SRT_Packet) GeoFeatureJSONPointResult {
	converted_long, _ := strconv.ParseFloat(packet.longtitude, 64)
	converted_lat, _ := strconv.ParseFloat(packet.latitude, 64)

	geo_prop := GeoProperty{
		Shutter:     packet.shutter,
		FrameCount:  strToInt64(packet.frame_count),
		DiffTime:   packet.diff_time,
		ISO: 		 strToInt32(packet.iso),
		FNUM: 		 strToInt32(packet.fnum),
		EV:			 strToInt32(packet.ev),
		CT: 	     strToInt64(packet.ct),
		ColorMd: 	 packet.color_md,
		FocalLen:	 strToInt32(packet.focal_len),
		Latitude:	converted_lat,
		Longitude:	converted_long,
		Altitude:	strToFloat64(packet.altitude),
		Date: 		packet.date,
		Timestamp: packet.time_stamp,
		Barometer:  strToFloat64(packet.barometer),
	}

	amount := len(parser.packets) 
	packet.printSRTPacket(strconv.Itoa(amount))

	result := GeoFeatureJSONPointResult{
		Type: "Feature",
		Geometry: PointGeometry{
			Type: "Point",
			Coordinates: []float64{
				converted_long, 
				converted_lat,
				strToFloat64(packet.altitude),
			},
		},
		Properties: geo_prop,
		
	}

	return result
}


func filterEmptyLines(lines []string) []string {
	result := make([]string, 0, len(lines))
	for _, line := range lines {
		if len(line) > 0 {
			result = append(result, line)
		}
	}
	return result
}

func maybeParseNumbers(value string) interface{} {
	if number, err := strconv.Atoi(value); err == nil {
		return number
	}
	if number, err := strconv.ParseFloat(value, 64); err == nil {
		return number
	}
	return value
}

func convertValues(values []string) []interface{} {
	converted := make([]interface{}, len(values))
	for i, value := range values {
		converted[i] = maybeParseNumbers(value)
	}
	return converted
}

func isElementExist(s []string, str string) bool {
	for _, v := range s {
	  if v == str {
		return true
	  }
	}
	return false
}

func checkEmptyField(s string) string{
	if len(s) == 0 {
		return "UNSPECIFIED"
	} else {
		return s
	}
}

func checkValidFileContents(fileContent string) bool {
	if len(fileContent) == 0 {
		PrintError("INVALID FILE CONTENT IS EMPTY")
		return false
	}

	return true
}

func checkNullPacket(packet SRT_Packet) bool {
	return (packet.diff_time == "" && packet.iso == "" && packet.shutter == "" && packet.fnum == "" && packet.ev == "" && packet.ct == "" && packet.color_md == "" && packet.focal_len == "" && packet.latitude == "" && packet.longtitude == "" && packet.altitude == "" && packet.date == "" && packet.time_stamp == "")
}

func strToInt32(input string) int32 {
	val, _ := strconv.ParseInt(input, 10, 32)
	return int32(val)
}

func strToInt64(input string) int64 {
	val, _ := strconv.ParseInt(input, 10, 64)
	return val
}

func strToFloat64(input string) float64 {
	val, _ := strconv.ParseFloat(input, 64)
	return val
}
