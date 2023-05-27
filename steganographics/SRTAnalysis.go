package steganographics

// subtitle extratoir
// SRT Viewer

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
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
	customProperties map[string]interface{}
}

func (p *DJI_SRT_Parser) srtToObject(srt string) []map[string]interface{} {
	maybeParseNumbers := func(d string) interface{} {
		if isNum(d) {
			num, err := strconv.ParseFloat(d, 64)
			if err != nil {
				return d
			}
			return num
		}
		return d
	}

	// Convert SRT strings file into an array of objects
	converted := make([]map[string]interface{}, 0)
	timecodeRegEx := regexp.MustCompile(`(\d{2}:\d{2}:\d{2},\d{3})\s-->\s`)
	packetRegEx := regexp.MustCompile(`^\d+$`)
	arrayRegEx := regexp.MustCompile(`\b([A-Z_a-z]+)\(([-\+\w.,/]+)\)`)
	valueRegEx := regexp.MustCompile(`\b([A-Z_a-z]+)\s?:[\s\[a-z_A-Z\]]?([-\+\d./]+)\w{0,3}\b`)
	dateRegEx := regexp.MustCompile(`\d{4}[-.]\d{1,2}[-.]\d{1,2} \d{1,2}:\d{2}:\d{2,}`)
	accurateDateRegex := regexp.MustCompile(`(\d{4}[-.]\d{1,2}[-.]\d{1,2} \d{1,2}:\d{2}:\d{2}),(\w{3}),(\w{3})`)
	accurateDateRegex2 := regexp.MustCompile(`(\d{4}[-.]\d{1,2}[-.]\d{1,2} \d{1,2}:\d{2}:\d{2})[,.](\w{3})`)
	isDJIFPV := regexp.MustCompile(`font size="28"`).MatchString(srt) &&
		regexp.MustCompile(`\d{4}-\d{1,2}-\d{1,2} \d{1,2}:\d{2}:\d{2}.\d{3}`).MatchString(srt) &&
		regexp.MustCompile(`\[altitude: \d.*\]`).MatchString(srt)

	// Split difficult Phantom4Pro format
	srt = strings.ReplaceAll(srt, ",", ":separator:")
	srt = regexp.MustCompile(`\(([^\)]+)\)`).ReplaceAllStringFunc(srt, func(match string) string {
		match = strings.ReplaceAll(match, ",", ":separator:")
		match = strings.ReplaceAll(match, " ", "")
		return match
	})
	srt = strings.ReplaceAll(srt, ", ", " ")
	srt = strings.ReplaceAll(srt, "Â", "")
	srt = strings.ReplaceAll(srt, "°", "")
	srt = strings.ReplaceAll(srt, "B0", "")
	lines := strings.Split(srt, "\n")
	srtLines := make([]string, 0)
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) > 0 {
			srtLines = append(srtLines, line)
		}
	}

	for _, line := range srtLines {
		var match []string
		if packetRegEx.MatchString(line) {
			// New packet
			converted = append(converted, make(map[string]interface{}))
		} else if m := timecodeRegEx.FindStringSubmatch(line); len(m) > 0 {
			converted[len(converted)-1]["TIMECODE"] = m[1]
		} else {
			for _, m = range arrayRegEx.FindAllStringSubmatch(line, -1) {
				key := m[1]
				values := strings.Split(m[2], ",")
				parsedValues := make([]interface{}, len(values))
				for i, v := range values {
					parsedValues[i] = maybeParseNumbers(v)
				}
				converted[len(converted)-1][key] = parsedValues
			}
			for _, m = range valueRegEx.FindAllStringSubmatch(line, -1) {
				key := m[1]
				value := maybeParseNumbers(m[2])
				converted[len(converted)-1][key] = value
			}
			if m := isoDateRegex.FindStringSubmatch(line); len(m) > 0 {
				converted[len(converted)-1]["DATE"] = line
			} else if m := accurateDateRegex.FindStringSubmatch(line); len(m) > 0 {
				converted[len(converted)-1]["DATE"] = m[1] + ":" + m[2] + "." + m[3]
			} else if m := accurateDateRegex2.FindStringSubmatch(line); len(m) > 0 {
				converted[len(converted)-1]["DATE"] = m[1] + "." + m[2]
			} else if m := dateRegEx.FindStringSubmatch(line); len(m) > 0 {
				converted[len(converted)-1]["DATE"] = strings.ReplaceAll(m[0], ":"+m[2], "."+m[2])
			} else if isDJIFPV && regexp.MustCompile(`\[altitude: \d.*\]`).MatchString(line) {
				// Correct altitude divided by 10 problem in DJI FPV drone
				altitude := converted[len(converted)-1]["altitude"].(string)
				altitudeValue, err := strconv.Atoi(altitude)
				if err == nil {
					converted[len(converted)-1]["altitude"] = strconv.Itoa(altitudeValue * 10)
				}
			}
		}
	}

	if len(converted) < 1 || len(converted[0]) == 0 {
		fmt.Println("Error converting object")
		return nil
	}
	return converted
}

func (p *DJI_SRT_Parser) millisecondsPerSample(metadata map[string]interface{}, milliseconds int) []map[string]interface{} {
	newArr := metadata["packets"].([]map[string]interface{})

	millisecondsPerSampleTIMECODE := func(amount int) []map[string]interface{} {
		lastTimecode := 0
		newResArr := make([]map[string]interface{}, 0)

		getMilliseconds := func(timecode string) int {
			m := strings.Split(timecode, ",")
			t := strings.Split(m[0], ":")

			hours, _ := strconv.Atoi(t[0])
			minutes, _ := strconv.Atoi(t[1])
			seconds, _ := strconv.Atoi(t[2])
			milliseconds, _ := strconv.Atoi(m[1])

			totalMilliseconds := (hours*60*60 + minutes*60 + seconds) * 1000
			return totalMilliseconds + milliseconds
		}

		for i := 0; i < len(newArr); i++ {
			millisecondsFromTimecode := getMilliseconds(newArr[i]["TIMECODE"].(string))

			if millisecondsFromTimecode < lastTimecode {
				continue
			}

			newResArr = append(newResArr, newArr[i])
			lastTimecode = millisecondsFromTimecode + amount
		}

		return newResArr
	}

	if len(newArr) > 0 && newArr[0]["TIMECODE"] != nil {
		if milliseconds != 0 {
			newArr = millisecondsPerSampleTIMECODE(milliseconds)
		}
		p.millisecondsSample = milliseconds
	}

	return newArr
}

	
func isNum(val string) bool {
	match, _ := regexp.MatchString(`^[-+\d.,]+$`, val)
	return match
}