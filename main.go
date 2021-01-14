package main

import (
	"flag"
	"fmt"
	"github.com/abusizhishen/fileConvert/convert"
)

var in string
var out string

func init() {
	flag.StringVar(&in, "in", "", "")
	flag.StringVar(&out, "out", "", "")
	flag.Parse()
}

func main() {
	err := convert.Convert(in,out)
	fmt.Println(err)
}


