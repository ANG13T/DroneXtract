package parsing

import (
	"strings"
	"path/filepath"
)

func CheckFileFormat(path string, exten string) bool {
	extension := strings.ToLower(filepath.Ext(path))
	return (extension == exten)
}