package src

import (
	"errors"
)

var ErrInvalidFileExt = errors.New("invalid file ext")

type SendPackage interface {
	Read(string2 string) ([]string, error)
	Write(string, []string) error
}

type Row []string
