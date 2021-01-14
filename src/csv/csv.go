package csv

import (
	"encoding/csv"
	"github.com/abusizhishen/fileConvert/src"
	"io"
	"os"
	"unsafe"
)

func Read(fileName string) ([]src.Row, error) {
	csvFile, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer csvFile.Close()
	// Parse the file
	r := csv.NewReader(csvFile)
	var rows []src.Row
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		rows = append(rows, record)
	}

	return rows, nil
}

func Write(fileName string, rows []src.Row) error {
	csvFile, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer csvFile.Close()

	bits := *(*[][]string)(unsafe.Pointer(&rows))
	r := csv.NewWriter(csvFile)
	err = r.WriteAll(bits)
	return err
}
