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

	max := 128
	chars := make([]Char, max)

	for i := 0; i < 128; i++ {
		char := byte(i)
		chars[i] = Char{char, string(char)}
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
