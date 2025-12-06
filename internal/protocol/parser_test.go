package protocol

import (
	"testing"

	"bufio"
	"strings"
)

func TestParseSimpleCommand(t *testing.T) {
	input := "SET key value\r\n"
	reader := bufio.NewReader(strings.NewReader(input))
	resp := &RESP{}
	req, err := resp.Parse(reader)
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}
	if req.cmd != "SET" {
		t.Errorf("Expected cmd 'SET', got '%s'", req.cmd)
	}
	if req.argsLen != 2 {
		t.Errorf("Expected argsLen 2, got %d", req.argsLen)
	}
	if len(req.args) != 2 || req.args[0] != "key" || req.args[1] != "value" {
		t.Errorf("Expected args ['key', 'value'], got %v", req.args)
	}
}

func TestParseEmptyCommand(t *testing.T) {
	input := "\r\n"
	reader := bufio.NewReader(strings.NewReader(input))
	resp := &RESP{}
	_, err := resp.Parse(reader)
	if err == nil {
		t.Error("Expected error for empty command, got nil")
	}
}

func TestParseMalformedCommand(t *testing.T) {
	input := "SET\r\nvalue\r\n"
	reader := bufio.NewReader(strings.NewReader(input))
	resp := &RESP{}
	_, err := resp.Parse(reader)
	if err == nil {
		t.Error("Expected error for malformed command, got nil")
	}
}
