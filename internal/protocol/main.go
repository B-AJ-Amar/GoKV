package protocol

import (
	"fmt"
)

const (
	simpleStrType   = '+'
	SimpleErrorType = '-'
	IntType         = ':'
	BulkStrType     = '$'
	ArrayType       = '*'
)

const (
	PingReqType = iota
	HelloReqType
	SetReqType
	SetExReqType
	GetReqType
)

const (
	SetRes    = "+OK\r\n"
	Hello2Res = "*14\r\n$6\r\nserver\r\n$5\r\nGpKv\r\n$7\r\nversion\r\n$11\r\n0.0.0-alpha\r\n$5\r\nproto\r\n:2\r\n$2\r\nid\r\n:4\r\n$4\r\nmode\r\n$9\r\nstandalone\r\n$4\r\nrole\r\n$6\r\nmaster\r\n$7\r\nmodules\r\n*0\r\n"
	HelloRes  = "-NOPROTO unsupported protocol version\r\n"
)

type RESPReq struct {
	reqType int // XReqType
	argsLen int
	args    []string
}

type RESPRes struct {
	message string
}

type Protocol interface {
	Parse(msg string) (*RESPReq, error)
	Process(req *RESPReq) RESPRes
}

type RESP struct{}

func main() {
	fmt.Println("Hello protocol")
}
