package main

import (
	"flag"
	"fmt"
	"path/filepath"
	"uploadAndDownload/file"
	"uploadAndDownload/file/csv"
)


var input string
var output string

func init() {
	flag.StringVar(&input,"input","","")
	flag.StringVar(&output,"output","","")
	flag.Parse()
}

func main() {
	var inputExt,outputExt = getFileType(input),getFileType(output)
	var readFunc,writeFunc = readFun[inputExt],writeFun[outputExt]
	//var Type string
	rows,err := readFunc(input)
	if err != nil{
		panic(err)
	}

	fmt.Println("-----")
	fmt.Println(rows)
	fmt.Println("-----")

	rows = append(rows, []string{"13452676545","lisi"})

	err = writeFunc(output, rows)
	fmt.Println(err)
}

func getFileType(fileName string) Type {
	ext := filepath.Ext(fileName)
	switch Type(ext) {
	case CSV:
		return CSV
	case Excel:
		return Excel
	default:
		panic(fmt.Sprintf("invalid file ext :%s", ext))
	}

	return CSV
}



type Type string
const(
	CSV Type = ".csv"
	Excel Type = ".xlsx"
)


type WriteFun func(string,[]file.Row)error
type ReadFun func(string)([]file.Row,error)

var writeFun = map[Type]WriteFun{
	CSV:csv.Write,
}

var readFun = map[Type]ReadFun{
	CSV:csv.Read,
}
