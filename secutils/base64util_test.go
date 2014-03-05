package secutils

import (
	"fmt"
	"testing"
)

func TestASCIIStringToBase64(t *testing.T) {

	data := "the quick brown fox jumps over the lazy dog"
	expected := "dGhlIHF1aWNrIGJyb3duIGZveCBqdW1wcyBvdmVyIHRoZSBsYXp5IGRvZw=="
	result := ASCIIStringToBase64(data)
	if result != expected {
		t.Error("ASCIIStringToBase64 failed !")
	}

	fmt.Println("Test ASCIIStringToBase64: OK")
}

func TestBase64StringToBytes(t *testing.T) {

	expected := "the quick brown fox jumps over the lazy dog"
	data := "dGhlIHF1aWNrIGJyb3duIGZveCBqdW1wcyBvdmVyIHRoZSBsYXp5IGRvZw=="
	resultBytes := Base64StringToBytes(data)
	result := BytesToASCIIString(resultBytes)
	if result != expected {
		t.Error("Base64StringToBytes failed !")
	}

	fmt.Println("Test Base64StringToBytes: OK")
}

func TestBase64StringToASCII(t *testing.T) {

	expected := "the quick brown fox jumps over the lazy dog"
	data := "dGhlIHF1aWNrIGJyb3duIGZveCBqdW1wcyBvdmVyIHRoZSBsYXp5IGRvZw=="
	result := Base64StringToASCII(data)
	if result != expected {
		t.Error("Base64StringToASCII failed !")
	}

	fmt.Println("Test Base64StringToASCII: OK")
}

func TextBase64ToHex(t *testing.T) {

	base64 := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	expectedHex := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	result := Base64ToHex(base64)
	if result != expectedHex {
		t.Error("Base64ToHex failed !")
	}

	fmt.Println("Test Base64ToHex: OK")

}
