package secutils

import (
	"fmt"
	"testing"
)

func TestASCIIStringToHex(t *testing.T) {

	data := "the quick brown fox jumps over the lazy dog"
	//fmt.Printf("%x\n", data)
	expected := "74686520717569636b2062726f776e20666f78206a756d7073206f76657220746865206c617a7920646f67"
	result := ASCIIStringToHex(data)
	if result != expected {
		t.Error("ASCIIStringToHex failed !")
	}

	fmt.Println("Test ASCIIStringToHex: OK")
}

func TestHexStringToASCII(t *testing.T) {

	expected := "the quick brown fox jumps over the lazy dog"
	//fmt.Printf("%x\n", data)
	data := "74686520717569636b2062726f776e20666f78206a756d7073206f76657220746865206c617a7920646f67"
	result := HexStringToASCII(data)
	if result != expected {
		t.Error("HexStringToASCII failed !")
	}

	fmt.Println("Test HexStringToASCII: OK")
}

func TestHexStringToBytes(t *testing.T) {

	expected := "the quick brown fox jumps over the lazy dog"
	//fmt.Printf("%x\n", data)
	data := "74686520717569636b2062726f776e20666f78206a756d7073206f76657220746865206c617a7920646f67"
	resultBytes := HexStringToBytes(data)
	result := BytesToASCIIString(resultBytes)
	if result != expected {
		t.Error("HexStringToBytes failed !")
	}

	fmt.Println("Test HexStringToBytes: OK")
}

func TextHexToBase64(t *testing.T) {

	hex := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	expectedBase64 := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	result := HexToBase64(hex)
	if result != expectedBase64 {
		t.Error("HexToBase64 failed !")
	}

	fmt.Println("Test HexToBase64: OK")

}
