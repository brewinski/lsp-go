package rpc_test

import (
	"testing"

	"github.com/brewinski/lsp-go/rpc"
)

type EncodingExample struct {
	Testing bool
}

func TestEncodeMessage(t *testing.T) {
	expected := "Content-Length: 16\r\n\r\n{\"Testing\":true}"
	actual := rpc.EncodeMessage(EncodingExample{Testing: true})

	if actual != expected {
		t.Errorf("Expected %s but got %s", expected, actual)
	}
}

func TestDecodeMessage(t *testing.T) {
	incommingMessage := "Content-Length: 15\r\n\r\n{\"Method\":\"hi\"}"
	method, content, err := rpc.DecodeMessage([]byte(incommingMessage))
	contentLength := len(content)
	if err != nil {
		t.Errorf("Expected no error but got %s", err)
	}

	if method != "hi" {
		t.Errorf("Expected hi but got %s", method)
	}

	if contentLength != 15 {
		t.Errorf("Expected 15 but got %d", contentLength)
	}
}
