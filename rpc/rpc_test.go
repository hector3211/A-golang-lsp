package rpc_test

import (
	"golang-lsp/rpc"
	"testing"
)

type EncodingExample struct {
	Testing bool `json:"testing"`
}

func TestEncodeMessage(t *testing.T) {
	expected := "Content-Length: 16\r\n\r\n{\"testing\":true}"
	actual := rpc.EncodeMessage(EncodingExample{Testing: true})

	if expected != actual {
		t.Fatalf("Expected: %s, Actual: %s", expected, actual)
	}
}

func TestDecodeMessage(t *testing.T) {
	incoming := "Content-Length: 15\r\n\r\n{\"Method\":\"hi\"}"
	method, content, err := rpc.DecodeMessage([]byte(incoming))
	contentLength := len(content)

	if err != nil {
		t.Fatal(err)
	}

	if contentLength != 15 {
		t.Fatalf("expected: %d, Got: %d", 16, contentLength)
	}

	if method != "hi" {
		t.Fatalf("expected: hi, got: %s", method)
	}
}
