package steganography

import (
	"image"
	"image/png"
	"os"

	_ "github.com/mdouchement/dng"
)

var (
	input  = "/Users/angelinatsuboi/Desktop/DJI-Forensics/dataset/DJI_0234.dng"
	output = "/Users/angelinatsuboi/Desktop/DJI-Forensics/dataset/output.png"
)

func DNGtoPNG() {
	fi, err := os.Open(input)
	check(err)
	defer fi.Close()

	m, _, err := image.Decode(fi)
	check(err)

	fo, err := os.Create(output)
	check(err)

	png.Encode(fo, m)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}