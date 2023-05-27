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

type SRT_Packet struct {
	date		string
	mapMatch	[]interface{}
	timeCode 	string
	dateStamp 	string
	altitude 	string
}

func (parser *DJI_SRT_Parser) SRTToObject(srt string) {
	maybeParseNumbers := func(d string) interface{} {
		if isNum(d) {
			num, _ := strconv.ParseFloat(d, 64)
			return num
		}
		return d
	}

	converted := make([]SRT_Packet, 0)
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
		} else if match = timecodeRegEx.FindStringSubmatch(line); match != nil {
			converted[len(converted)-1].timeCode = match[1]
		} else {
			for _, match := range arrayRegEx.FindAllStringSubmatch(line, -1) {
				values := strings.Split(match[2], ",")
				converted[len(converted)-1].mapMatch = convertValues(values)
			}
			for _, match = valueRegEx.FindStringSubmatch(line); match != nil; match = valueRegEx.FindStringSubmatch(line) {
				converted[len(converted)-1].mapMatch = maybeParseNumbers(match[2])
			}
			if match = isoDateRegex.FindStringSubmatch(line); match != nil {
				converted[len(converted)-1].dateStamp = line
			} else if match = accurateDateRegex.FindStringSubmatch(line); match != nil {
				converted[len(converted)-1].dateStamp = match[1] + ":" + match[2] + "." + match[3]
			} else if match = accurateDateRegex2.FindStringSubmatch(line); match != nil {
				converted[len(converted)-1].dateStamp = match[1] + "." + match[2]
			} else if match = dateRegEx.FindStringSubmatch(line); match != nil {
				converted[len(converted)-1].dateStamp = strings.ReplaceAll(match[0], ":"+match[2]+match[3]+"$", "."+match[2])
			} else if isDJIFPV && regexp.MustCompile(`\[altitude: \d.*\]`).MatchString(line) {
				// Correct altitude divided by 10 problem in DJI FPV drone
				altitude := converted[len(converted)-1].altitude
				converted[len(converted)-1].altitude = strconv.Itoa(maybeParseNumbers(altitude)*10)
			}
		}
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
