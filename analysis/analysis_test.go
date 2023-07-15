package analysis


type DJI_Analysis struct {
	fileName        string
	outputPath		string
}

func NewDJI_Analysis(fileName string) *DJI_Analysis {
	parser := DJI_Analysis{
		fileName: fileName,
	}
	return &parser
}

func (parser *DJI_Analysis) ExecuteAnalysis() {
	
}
