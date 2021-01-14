package excel

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"github.com/abusizhishen/fileConvert/src"
)

func Read(fileName string) (rows []src.Row, err error) {
	wb, err := xlsx.OpenFile(fileName)
	if err != nil {
		return
	}

	if len(wb.Sheet) == 0 {
		err = fmt.Errorf("empty exsl")
		return
	}

	var sheet = wb.Sheets[0]
	for _, row := range sheet.Rows {
		var data src.Row
		for _, cell := range row.Cells {
			data = append(data, cell.Value)
		}
		rows = append(rows, data)
	}

	return
}

func Write(fileName string, rows []src.Row) error {
	wb := xlsx.NewFile()

	sheet, err := wb.AddSheet("first")
	if err != nil {
		return err
	}

	for _, data := range rows {
		row := sheet.AddRow()
		for _, value := range data {
			cell := row.AddCell()
			cell.SetValue(value)
		}
	}

	return wb.Save(fileName)
}
