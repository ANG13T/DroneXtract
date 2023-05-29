package steganographics

// TODO -
// 0 - fix everything
// 1 - Finish Table Display
// 2 - toGeoJSON
// 3 - toCSV
// 4 - toMGJSON
// 5 - subtitle extractor
// 6 - display metadata about SRT file
// 7 - comments

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"io/ioutil"
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

func (packet *SRT_Packet) printSRTPacket() {
	title := "Frame " + checkEmptyField(packet.frame_count)
	GenTableHeader(title)
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
	GenTableFooter()
}

func (parser *DJI_SRT_Parser) SRTToObject(srt string) []SRT_Packet {
	converted := make([]SRT_Packet, 0)
	test_regex := regexp.MustCompile(`\[(\w+)\s*:\s*([^]]+)\]`)
	test_regex_2 := regexp.MustCompile(`([A-Za-z]+):([-.\w/]+)`)
	diffTimeRegex := regexp.MustCompile(`\bDiffTime\s*:\s*([^ ]+)`)
	timecodeRegEx := regexp.MustCompile(`(\d{2}:\d{2}:\d{2},\d{3})\s-->\s`)
	packetRegEx := regexp.MustCompile(`^\d+$`)
	arrayRegEx := regexp.MustCompile(`\b([A-Z_a-z]+)\(([-\+\w.,/]+)\)`)
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

	lines = filterEmptyLines(lines) //maybe

	for _, line := range lines {
		var match []string
		matched := packetRegEx.MatchString(line)


		if matched {
			// new packet
			converted = append(converted, SRT_Packet{})
			fmt.Println("LINE 1: ", line)
			converted[len(converted)-1].frame_count = line
		} else if match = timecodeRegEx.FindStringSubmatch(line); match != nil {
			fmt.Println("timestamp", match[1])
			values := strings.Split(match[1], ",")
			converted[len(converted)-1].time_stamp = values[0]
			fmt.Println("LINE 2: ", line)
		} else {
			for _, match := range arrayRegEx.FindStringSubmatch(line) {
				//values := strings.Split(match[2], ",")
				// converted[len(converted)-1].mapMatch = convertValues(values)
				// fmt.Println("LINE 3: ", converted[len(converted)-1].mapMatch)
				//fmt.Println("LINE 33: ", values)
				fmt.Println("MATCHES: ", match)
			}
			
			fmt.Println("LINE 4: ", line)

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
				fmt.Printf("Key: %s, Value: %s\n", key, value)

				switch key {
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
		}
	}

	if len(converted) < 1 || (len(converted) == 1 && checkNullPacket(converted[0])) {
		PrintError("ERROR PARSING SRT FILE")
		return nil
	}

	return converted
}

func (parser *DJI_SRT_Parser) GeneratePackets(path string) {
	// Check if Valid File Path
	content, err := ioutil.ReadFile(path)

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
		packet.printSRTPacket()
	}
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