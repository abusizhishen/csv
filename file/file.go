package file

type SendPackage interface {
	Read(string2 string)([]string,error)
	Write(string,[]string)error
}

type Row []string