package main

import "fmt"

const (
	simpleStrType   = '+'
	SimpleErrorType = '-'
	IntType         = ':'
	BulkStrType     = '$'
	ArrayType       = '*'
	NullType        = '_'
	BoolType        = '#'
	DoubleType      = ','
	BigNumType      = '('
	BulkErrorType   = '!'
	VerbatimStrType = '='
	MapType         = '%'
	SetType         = '~'
	PushType        = '>'
)

type RESP interface {
	Parse(msg string) RESPReq
	Process(req *RESPReq) RESPRes
}

func main() {
	fmt.Println("Hello protocol")
}
