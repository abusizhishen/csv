package convert

import (
	"github.com/abusizhishen/fileConvert/src"
	"github.com/abusizhishen/fileConvert/src/csv"
	"github.com/abusizhishen/fileConvert/src/excel"
	"path/filepath"
)

func getFileType(fileName string) (Type, error) {
	ext := filepath.Ext(fileName)
	switch Type(ext) {
	case CSV:
		return CSV, nil
	case Excel:
		return Excel, nil
	default:
		return "", src.ErrInvalidFileExt
	}
}

type Type string

const (
	CSV   Type = ".csv"
	Excel Type = ".xlsx"
)

type WriteFun func(string, []src.Row) error
type ReadFun func(string) ([]src.Row, error)

var writeFun = map[Type]WriteFun{
	CSV:   csv.Write,
	Excel: excel.Write,
}

var readFun = map[Type]ReadFun{
	CSV:   csv.Read,
	Excel: excel.Read,
}

func Convert(inputFile, outPutFile string) error {
	inputExt, err := getFileType(inputFile)
	if err != nil {
		return err
	}
	outputExt, err := getFileType(outPutFile)
	if err != nil {
		return err
	}

	var readFunc, writeFunc = readFun[inputExt], writeFun[outputExt]
	rows, err := readFunc(inputFile)
	if err != nil {
		return err
	}

	return writeFunc(outPutFile, rows)
}

type Data interface {
	ToRow() []src.Row
}

func LoadCity(fileName string) (citys Citys, err error) {
	inputExt, err := getFileType(fileName)
	if err != nil {
		return
	}
	rows, err := readFun[inputExt](fileName)
	if err != nil {
		return
	}

	//	忽略行首

	return ToCity(rows[1:])
}

func LoadUser(fileName string) (users Users, err error) {
	inputExt, err := getFileType(fileName)
	if err != nil {
		return
	}
	rows, err := readFun[inputExt](fileName)
	if err != nil {
		return
	}

	return ToUser(rows[1:])
}
