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
