package protocol

import (
	"testing"

	"github.com/B-AJ-Amar/gokv/internal/store"
)

func TestProcessSetAndGet(t *testing.T) {
	mem := store.NewInMemoryStore()
	resp := &RESP{}
	// SET key value
	setReq := &RESPReq{cmd: "set", argsLen: 3, args: []string{"set", "mykey", "value"}}
	setRes, err := resp.Process(setReq, mem)
	if err != nil {
		t.Fatalf("Process SET failed: %v", err)
	}
	if setRes.msgType != SpecialRes || setRes.message != SetResMsg {
		t.Errorf("SET response incorrect: got type %d, msg %q", setRes.msgType, setRes.message)
	}

	// GET key
	getReq := &RESPReq{cmd: "get", argsLen: 2, args: []string{"get", "mykey"}}
	getRes, err := resp.Process(getReq, mem)
	if err != nil {
		t.Fatalf("Process GET failed: %v", err)
	}
	if getRes.msgType != BulkStrRes || getRes.message != "value" {
		t.Errorf("GET response incorrect: got type %d, msg %q", getRes.msgType, getRes.message)
	}
}

func TestProcessPing(t *testing.T) {
	mem := store.NewInMemoryStore()
	resp := &RESP{}
	pingReq := &RESPReq{cmd: "ping", argsLen: 1, args: []string{"ping"}}
	pingRes, err := resp.Process(pingReq, mem)
	if err != nil {
		t.Fatalf("Process PING failed: %v", err)
	}
	if pingRes.msgType != SpecialRes || pingRes.message != PongResMsg {
		t.Errorf("PING response incorrect: got type %d, msg %q", pingRes.msgType, pingRes.message)
	}
}

func TestProcessHello(t *testing.T) {
	mem := store.NewInMemoryStore()
	resp := &RESP{}
	helloReq := &RESPReq{cmd: "hello", argsLen: 2, args: []string{"hello", "2"}}
	helloRes, err := resp.Process(helloReq, mem)
	if err != nil {
		t.Fatalf("Process HELLO failed: %v", err)
	}
	if helloRes.msgType != SpecialRes || helloRes.message != Hello2ResMsg {
		t.Errorf("HELLO response incorrect: got type %d, msg %q", helloRes.msgType, helloRes.message)
	}

	helloReq = &RESPReq{cmd: "hello", argsLen: 2, args: []string{"hello", "3"}}
	helloRes, err = resp.Process(helloReq, mem)
	if err != nil {
		t.Fatalf("Process HELLO failed: %v", err)
	}
	if helloRes.msgType != SpecialRes || helloRes.message != HelloResMsg {
		t.Errorf("HELLO response incorrect: got type %d, msg %q", helloRes.msgType, helloRes.message)
	}
}

func TestProcessUnknownCommand(t *testing.T) {
	mem := store.NewInMemoryStore()
	resp := &RESP{}
	req := &RESPReq{cmd: "unknown", argsLen: 1, args: []string{"unknown"}}
	res, err := resp.Process(req, mem)
	if err != nil {
		t.Fatalf("Process unknown command failed: %v", err)
	}
	if res.msgType != ErrorRes || res.message != "ERR unknown command" {
		t.Errorf("Unknown command response incorrect: got type %d, msg %q", res.msgType, res.message)
	}
}
