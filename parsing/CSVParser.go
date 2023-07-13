package parsing

import (
	"encoding/csv"
	"fmt"
	"github.com/ANG13T/DroneXtract/helpers"
	"github.com/jedib0t/go-pretty/v6/table"
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
	tableValue := table.NewWriter()
    tableValue.SetOutputMirror(os.Stdout)

	file, err := os.Open(parser.fileName)
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

	columns := records[0]

	tableValue.AppendHeader(ArrayToRow(columns))

	tableBody := []table.Row{}

	// Print each record
	for _, record := range records {
		row := []string{}
		for _, value := range record {
			row = append(row, value)
		}
		tableBody = append(tableBody, ArrayToRow(row))
		fmt.Println()
	}

	tableValue.AppendRows(tableBody)
	fmt.Printf("Table without any customizations:\n%s", tableValue.Render())
}

func ArrayToRow(input []string) table.Row {
	row := make(table.Row, len(input))
	for i, value := range input {
		row[i] = value
	}
	return row
}

func (parser *DJI_CSV_Parser) PrintContents() {

}


