package excel

import (
	"fmt"
	"github.com/tealeg/xlsx/v3"
	_ "github.com/tealeg/xlsx/v3"
	"uploadAndDownload/file"
)


type Excel struct {
	CityId int `json:"city_id"`
	CityName string `json:"city_name"`
}

func Read(fileName string) (rows []string,err error) {
	wb, err := xlsx.OpenFile("../samplefile.xlsx")
	if err != nil {
		return
	}

	if len(wb.Sheet) == 0{
		err = fmt.Errorf("empty exsl")
		return
	}

	var sheet = wb.Sheets[0]
	var maxRow = sheet.MaxRow
	for i:=0;i<maxRow;i++{
		row,err := sheet.Row(i)
		if err != nil{
			return nil, err
		}


		row.ForEachCell()
	}


	fmt.Println("----")

	return rows,nil
}

func Write(fileName string, rows []file.Row) error {
	wb := xlsx.NewFile()

	sheet,err := wb.AddSheet("first")
	if err != nil{
		return err
	}

	for _,data := range rows{
		sheet.ro
		row := sheet.AddRow()
		for _,value := range data{
			cell := row.AddCell()
			cell.
		}
	}

}