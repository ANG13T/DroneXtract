package parsing

import (
	"encoding/csv"
	"fmt"
	"github.com/ANG13T/DroneXtract/helpers"
	"os"
)

type DJI_CSV_Parser struct {
	fileName        string
}

func NewDJI_CSV_Parser(fileName string) *DJI_CSV_Parser {
	check := CheckFileFormat(fileName, ".csv")
	if check == false {
		helpers.PrintError("INVALID FILE FORMAT. MUST BE CSV FILE")
		return nil
	}

	parser := DJI_CSV_Parser{
		fileName: fileName,
	}
	return &parser
}

func (parser *DJI_CSV_Parser) ParseContents() {
	file, err := os.Open("data.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	// Create a new CSV reader
	reader := csv.NewReader(file)

	// Read all records from the CSV file
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print each record
	for _, record := range records {
		for _, value := range record {
			fmt.Printf("%s\t", value)
		}
		fmt.Println()
	}
}

func (parser *DJI_CSV_Parser) PrintContents() {

}


