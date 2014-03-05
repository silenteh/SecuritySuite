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

func BytesToHex(bytes []byte) string {
	return hex.EncodeToString(bytes)
}

func ASCIIStringToHex(asciiString string) string {
	bytes := ASCIIStringToBytes(asciiString)
	return BytesToHex(bytes)
}

func HexToBase64(hexString string) string {

	hexBytes := HexStringToBytes(hexString)
	return BytesToBase64(hexBytes)
}
