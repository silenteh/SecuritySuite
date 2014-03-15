package aesutils

import (
	"fmt"
	"testing"
)

func TestPkcs7(t *testing.T) {

	blockSize := 16

	data := make([]byte, 20)

	for index := range data {
		data[index] = byte(index + 48)
	}

	padded := Pkcs7(data, blockSize)
	fmt.Printf("%d\n", padded)

}

func TestIsPkcs7(t *testing.T) {

	blockSize := 6

	firstSample := []byte{byte(48), byte(49), byte(50), byte(51), byte(52), byte(53)} //make([]byte, 4)
	noPad := IsPkcs7(firstSample, blockSize)
	if noPad {
		t.Error("Pkcs7 detection error !")
	}

	secondSample := []byte{byte(48), byte(49), byte(50), byte(51)} //make([]byte, 4)
	padded := Pkcs7(secondSample, blockSize)
	twoPads := IsPkcs7(padded, blockSize)
	if twoPads == false {
		t.Error("Pkcs7 detection error !")
	}

	thirdSample := []byte{byte(48), byte(49), byte(50)} //make([]byte, 4)
	padded = Pkcs7(thirdSample, blockSize)
	threePads := IsPkcs7(padded, blockSize)
	if threePads == false {
		t.Error("Pkcs7 detection error !")
	}

}
