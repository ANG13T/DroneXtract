package steganography

// TODO -

// 2 - toGeoJSON + toCSV + toMGJSON


// 0 - support all files
// 5 - subtitle extractor from videos (MP4 to SRT)
// 7 - comments
// 8 - test suite

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"io/ioutil"
	"os"
)

var isoDateRegex = regexp.MustCompile(`[0-9]{4}-[0-9]{2}-[0-9]{2}T[0-9]{2}:[0-9]{2}:[0-9]{2}(\.[0-9]+)?Z`)

type DJI_SRT_Parser struct {
	fileName        string
	metadata        map[string]interface{}
	rawMetadata     []interface{}
	smoothened      int
	millisecondsSample int
	loaded          bool
	isMultiple      bool
	packets			[]SRT_Packet
}

func NewDJI_SRT_Parser(fileName string) *DJI_SRT_Parser {
	check := CheckFileFormat(fileName, "srt")
	if check == false {
		PrintError("INVALID FILE FORMAT. MUST BE SRT FILE")
		return nil
	}

	parser := DJI_SRT_Parser{
		fileName: fileName,
		metadata:           make(map[string]interface{}),
		rawMetadata:        make([]interface{}, 0),
		smoothened:         0,
		millisecondsSample: 0,
		loaded:             false,
		isMultiple:         false,
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

type GeoJSONResult struct {
	Type       string                 `json:"type"`
	Geometry   Geometry               `json:"geometry"`
	Properties map[string]interface{} `json:"properties"`
}

type Geometry struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}


func (packet *SRT_Packet) printSRTPacket(length string) {
	title := "Frame " + checkEmptyField(packet.frame_count)
	if packet.frame_count == "1" {
		GenTableHeader(title, true)
	} else {
		GenTableHeaderModified(title)
	}
	
	GenRowString("Frame Count", checkEmptyField(packet.frame_count))
	GenRowString("Diff Time", checkEmptyField(packet.diff_time))
	GenRowString("ISO", checkEmptyField(packet.iso))
	GenRowString("Shutter", checkEmptyField(packet.shutter))
	GenRowString("FNUM", checkEmptyField(packet.fnum))
	GenRowString("EV", checkEmptyField(packet.ev))
	GenRowString("CT", checkEmptyField(packet.ct))
	GenRowString("Color MD", checkEmptyField(packet.color_md))
	GenRowString("Focal Len", checkEmptyField(packet.focal_len))
	GenRowString("Latitude", checkEmptyField(packet.latitude))
	GenRowString("Longitude", checkEmptyField(packet.longtitude))
	GenRowString("Altitude", checkEmptyField(packet.altitude))
	GenRowString("Date", checkEmptyField(packet.date))
	GenRowString("Time Stamp", checkEmptyField(packet.time_stamp))
	GenRowString("Barometer", checkEmptyField(packet.barometer))
	
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

func (parser *DJI_SRT_Parser) GeoJSONExtract(raw []SRT_Packet) {

}

func (parser *DJI_SRT_Parser) CreateGeoJSON(raw []SRT_Packet, waypoints bool) {
	// elevationOffset := 0
	// result := GeoJSONResult{
	// 	Type: "Feature",
	// 	Geometry: Geometry{
	// 		Type:        "Point",
	// 		Coordinates: []float64{},
	// 	},
	// 	Properties: make(map[string]interface{}),
	// }
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

// Helpers

func isNum(d string) bool {
	_, err := strconv.ParseFloat(d, 64)
	return err == nil
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

// toGeoJSON Helpers
func extractProps(childObj map[string]interface{}, pre string) []map[string]interface{} {
	var results []map[string]interface{}
	for child, value := range childObj {
		if childValue, ok := value.(map[string]interface{}); ok && childValue != nil {
			children := extractProps(childValue, pre+"_"+child)
			for _, child := range children {
				results = append(results, child)
			}
		} else {
			results = append(results, map[string]interface{}{"name": pre + "_" + child, "value": value})
		}
	}
	return results
}

// func GeoJSONExtract(obj map[string]interface{}, raw bool) GeoJSONResult {
// 	result := map[string]interface{}{
// 		"type": "Feature",
// 		"properties": map[string]interface{}{
// 			"source":    "dji-srt-parser",
// 			"timestamp": []interface{}{},
// 			"name":      cleanFileName("test"),
// 		},
// 		"geometry": map[string]interface{}{
// 			"type":        "Point",
// 			"coordinates": []interface{}{},
// 		},
// 	}

// 	for key, value := range obj {
// 		// if key == "DATE" {
// 		// 	result.properties["timestamp"] = value
// 		// } else if key == "GPS" {
// 		// 	result.Geometry.Coordinates = extractCoordinates(value, raw)
// 		// } else if subObj, ok := value.(map[string]interface{}); ok && subObj != nil {
// 		// 	children := extractProps(subObj, key)
// 		// 	for _, child := range children {
// 		// 		// result.Properties[child.Name] = child.Value
// 		// 		// TODO
// 		// 		fmt.Println(child)
// 		// 	}
// 		// } else {
// 		// 	result.Properties[key] = value
// 		// }
// 	}

// 	return result
// }

func createLinestring(features []map[string]interface{}, fileName string, customProperties map[string]interface{}) map[string]interface{} {
	result := map[string]interface{}{
		"type": "Feature",
		"properties": map[string]interface{}{
			"source":    "dji-srt-parser",
			"timestamp": []interface{}{},
			"name":      cleanFileName(fileName),
		},
		"geometry": map[string]interface{}{
			"type":        "LineString",
			"coordinates": []interface{}{},
		},
	}

	props := features[0]["properties"].(map[string]interface{})
	for prop, value := range props {
		if !containsString([]string{
			"DATE",
			"TIMECODE",
			"GPS",
			"timestamp",
			"BAROMETER",
			"DISTANCE",
			"SPEED_THREED",
			"SPEED_TWOD",
			"SPEED_VERTICAL",
			"HB",
		}, prop) {
			result["properties"].(map[string]interface{})[prop] = value
		}
	}

	for _, feature := range features {
		result["geometry"].(map[string]interface{})["coordinates"] = append(result["geometry"].(map[string]interface{})["coordinates"].([]interface{}), feature["geometry"].(map[string]interface{})["coordinates"])
		result["properties"].(map[string]interface{})["timestamp"] = append(result["properties"].(map[string]interface{})["timestamp"].([]interface{}), feature["properties"].(map[string]interface{})["timestamp"])
	}

	return result
}

func containsString(slice []string, str string) bool {
	for _, item := range slice {
		if item == str {
			return true
		}
	}
	return false
}

func cleanFileName(fileName string) string {
	re := regexp.MustCompile(`\.[^/.]+$`) // Regular expression to match file extension
	return re.ReplaceAllString(fileName, "") // Remove file extension
}

func notReady() interface{} {
	fmt.Println("Data not ready")
	return nil
}

func extractCoordinates(coordsObj map[string]interface{}, raw bool) []float64 {
	coordResult := make([]float64, 3)
	if raw {
		if gps, ok := coordsObj["GPS"].([]interface{}); ok && len(gps) >= 2 {
			if val, ok := gps[0].(float64); ok {
				coordResult[0] = val
			}
			if val, ok := gps[1].(float64); ok {
				coordResult[1] = val
			}
		}
	} else {
		if gps, ok := coordsObj["GPS"].(map[string]interface{}); ok {
			if val, ok := gps["LONGITUDE"].(float64); ok {
				coordResult[0] = val
			}
			if val, ok := gps["LATITUDE"].(float64); ok {
				coordResult[1] = val
			}
			if elevation := getElevation(coordsObj); elevation != nil {
				coordResult[2] = *elevation
			}
		}
	}
	return coordResult
}

func getElevationKey(src map[string]interface{}) string {
	if _, ok := src["ALTITUDE"]; ok {
		return "ALTITUDE"
	} else if _, ok := src["BAROMETER"]; ok {
		return "BAROMETER"
	} else if _, ok := src["HB"]; ok {
		return "HB"
	}
	return "ALTITUDE"
}

func getElevation(src map[string]interface{}) *float64 {
	if val, ok := src["ALTITUDE"].(float64); ok {
		return &val
	} else if val, ok := src["BAROMETER"].(float64); ok {
		return &val
	} else if val, ok := src["HB"].(float64); ok {
		return &val
	}
	return nil
}

func preProcess(array []SRT_Packet, fileName string) {
	
}