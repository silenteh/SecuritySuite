package secutils

import (
	"math/big"
)

type Char struct {
	Byte   byte
	String string
}

var ASCII []Char = createASCII()

func createASCII() []Char {

	max := 127 - 32
	chars := make([]Char, max)

	index := 0
	for i := 32; i < 127; i++ {
		char := byte(i)
		chars[index] = Char{char, string(char)}
		index++
	}
	return chars

}

func ASCIIStringToBytes(asciiString string) []byte {
	return []byte(asciiString)
}

func BytesToASCIIString(bytes []byte) string {
	return string(bytes[:])
}

func BigIntToBytes(bigInt *big.Int) []byte {
	return bigInt.Bytes()
}

func BigIntToHex(bigInt *big.Int) string {
	bytes := BigIntToBytes(bigInt)
	return BytesToHex(bytes)
}

func BigIntToBase64(bigInt *big.Int) string {
	bytes := BigIntToBytes(bigInt)
	return BytesToBase64(bytes)
}
