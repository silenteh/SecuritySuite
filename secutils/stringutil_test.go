package secutils

import (
	"fmt"
	"testing"
)

func TestASCIIStringToBytes(t *testing.T) {

	data := "this is a test string"
	dataBytes := ASCIIStringToBytes(data)
	dataBytesToSring := string(dataBytes[:])
	if dataBytesToSring != data {
		t.Error("ASCII string to bytes and back conversion error !!")
	}
	fmt.Println("Test ASCIIStringToBytes: OK")
}
