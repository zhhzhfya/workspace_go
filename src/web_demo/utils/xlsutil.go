package utils

import (
	"github.com/tealeg/xlsx"
	"fmt"
)

func WriteXls(){

	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell
	var err error

	file = xlsx.NewFile()
	sheet, err = file.AddSheet("Sheet1")
	if err != nil {
		fmt.Printf(err.Error())
	}
	row = sheet.AddRow()
	for a := 0; a < 10; a++ {
		cell = row.AddCell()
		cell.Value = "I am a cell!"
	}
	err = file.Save("MyXLSXFile.xlsx")
	if err != nil {
		fmt.Printf(err.Error())
	}
}