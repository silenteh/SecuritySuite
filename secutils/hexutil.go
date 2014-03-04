package secutils

import (
	"encoding/hex"
	"log"
)

func HexStringToBytes(hexString string) []byte {

	data, err := hex.DecodeString(hexString)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func HexStringToASCII(hexString string) string {
	data := HexStringToBytes(hexString)
	return string(data[:])
}

func ASCIIStringToHex(asciiString string) string {
	bytes := ASCIIStringToBytes(asciiString)
	return hex.EncodeToString(bytes)
}

func ASCIIStringToBytes(asciiString string) []byte {
	return []byte(asciiString)
}
