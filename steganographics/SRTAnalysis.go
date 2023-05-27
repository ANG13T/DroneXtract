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

