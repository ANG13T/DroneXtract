package steganography

type DJI_XMP_Parser struct {
	fileName        string
}


func NewDJI_XMP_Parser(fileName string) *DJI_XMP_Parser {
	check := CheckFileFormat(fileName, "xmp")
	if check == false {
		PrintError("INVALID FILE FORMAT. MUST BE XMP FILE")
		return nil
	}
	
	parser := DJI_XMP_Parser{
		fileName: fileName,
	}
	return &parser
}

func (parser *DJI_XMP_Parser) Read() { 
	exif := NewDJI_EXIF_Parser(parser.fileName)
	exif.Read()
}