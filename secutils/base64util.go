package secutils

import (
	"encoding/base64"
	"log"
)

func Base64StringToBytes(base64String string) []byte {

	data, err := base64.StdEncoding.DecodeString(base64String)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func Base64StringToASCII(base64String string) string {
	data := Base64StringToBytes(base64String)
	return string(data[:])
}

func BytesToBase64(bytes []byte) string {
	return base64.StdEncoding.EncodeToString(bytes)
}

func ASCIIStringToBase64(asciiString string) string {
	bytes := ASCIIStringToBytes(asciiString)
	return BytesToBase64(bytes)
}

func Base64ToHex(base64String string) string {

	base64Bytes := Base64StringToBytes(base64String)
	return BytesToHex(base64Bytes)
}
