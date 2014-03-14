package secutils

import (
	"bytes"
	"log"
	"math/big"
)

func Xor(s1, s2 string) *big.Int {

	if len(s1) != len(s2) {
		log.Fatal("The two strings have different length !")
	}

	s1Buf := new(big.Int).SetBytes([]byte(s1))
	s2Buf := new(big.Int).SetBytes([]byte(s2))
	result := new(big.Int).Xor(s1Buf, s2Buf)
	return result
}

func XorBytes(b1, b2 []byte) []byte {
	if len(b1) != len(b2) {
		log.Fatal("The two bytes array have different length !")
	}
	var byteBuffer bytes.Buffer
	for index := range b1 {
		b1Byte := b1[index]
		b2Byte := b2[index]
		xored := XorSingleByte(b1Byte, b2Byte)
		byteBuffer.WriteByte(xored)
	}
	return byteBuffer.Bytes()
}

func XorSingleByte(b1, b2 byte) byte {
	return b1 ^ b2
}
