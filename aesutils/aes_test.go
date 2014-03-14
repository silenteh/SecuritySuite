package aesutils

import (
	"fmt"
	"github.com/silenteh/security_suite/secutils"
	"testing"
)

func TestEncryptAES_CBC(t *testing.T) {

	text := "YELLOW SUBMARINE"
	key := "YELLOW SUBMARINE"
	cipherTextBytes, _ := EncryptAES_CBC(text, key)
	keyBytes := secutils.ASCIIStringToBytes(key)
	clearText, _ := DecryptAES_CBC(cipherTextBytes, keyBytes)
	fmt.Printf("Decrypted: %s\n", clearText)

	if clearText != text {
		t.Error("Test EncryptAES_CBC and DencryptAES_CBC error !")
	}
}

func TestDencryptAES_CBC(t *testing.T) {

	text := "YELLOW SUBMARINE!"
	key := "YELLOW SUBMARINE"
	cipherTextBytes, _ := EncryptAES_CBC(text, key)
	keyBytes := secutils.ASCIIStringToBytes(key)
	clearText, _ := DecryptAES_CBC(cipherTextBytes, keyBytes)
	fmt.Printf("Decrypted: %s\n", clearText)

	// if clearText != text {
	// 	t.Error("Test EncryptAES_CBC and DencryptAES_CBC error !")
	// }
}
