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
	fmt.Printf("1. Test EncryptAES_CBC: %s\n", clearText)

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
	fmt.Printf("2. Test DencryptAES_CBC: %s\n", clearText)
}

func TestEncryptAES_ECB_Manual(t *testing.T) {

	key := "YELLOW SUBMARINE"

	expectedHexCipher := "d1aa4f6578926542fbb6dd876cd20508"
	expectedHexCipherPadding := "d1aa4f6578926542fbb6dd876cd20508be02542eda8a4ae9cd80e9ce20751237"

	text := "YELLOW SUBMARINE"
	textPadding := "YELLOW SUBMARINE!"

	cipherText := EncryptAES_ECB_Manual(text, key)
	cipherTextHex := secutils.BytesToHex(cipherText)
	if expectedHexCipher != cipherTextHex {
		t.Error("Test Encrypt AES_ECB Manually FAILED !")
	}
	fmt.Printf("3. Test EncryptAES_ECB_Manual: %s\n", cipherTextHex)

	cipherTextPadding := EncryptAES_ECB_Manual(textPadding, key)
	cipherTextHexPadding := secutils.BytesToHex(cipherTextPadding)
	if expectedHexCipherPadding != cipherTextHexPadding {
		t.Error("Test Encrypt AES_ECB Manually + PADDING FAILED !")
	}
	fmt.Printf("3. Test EncryptAES_ECB_Manual + PADDING: %s\n", cipherTextHexPadding)

}

func TestDecryptAES_ECB_Manual(t *testing.T) {

	key := "YELLOW SUBMARINE"

	expetedText := "YELLOW SUBMARINE"
	expetedTextPadding := "YELLOW SUBMARINE!"

	cipherText := "d1aa4f6578926542fbb6dd876cd20508"
	cipherTextPadding := "d1aa4f6578926542fbb6dd876cd20508be02542eda8a4ae9cd80e9ce20751237"

	textBytes := secutils.HexStringToBytes(cipherText)
	text := DecryptAES_ECB_Manual(textBytes, []byte(key))
	if expetedText != secutils.BytesToASCIIString(text) {
		t.Error("Test Decrypt AES_ECB Manually FAILED !")
	}
	fmt.Printf("4. Test DecryptAES_ECB_Manual: %s\n", text)

	textBytesPadding := secutils.HexStringToBytes(cipherTextPadding)
	textPadding := DecryptAES_ECB_Manual(textBytesPadding, []byte(key))
	if expetedTextPadding != secutils.BytesToASCIIString(textPadding) {
		t.Error("Test Decrypt AES_ECB Manually + PADDING FAILED !")
	}
	fmt.Printf("4. Test DecryptAES_ECB_Manual + PADDING: %s\n", textPadding)

}

func TestEncryptAES_CBC_Manual(t *testing.T) {
	text := "YELLOW SUBMARINE"
	key := "YELLOW SUBMARINE"
	cipherTextBytes := EncryptAES_CBC_Manual(text, key)
	cipherHex := secutils.BytesToHex(cipherTextBytes)
	fmt.Printf("5. Test EncryptAES_CBC Manual: %s\n", cipherHex)
}

//d7ab340a60337584edbab1a82d6b0977
func TestDecryptAES_CBC_Manual(t *testing.T) {
	text := secutils.HexStringToBytes("a6af0c64422b96e8923c1ebe819e82553e8e38f74e1693a9cde95502229ddcc0")
	key := secutils.ASCIIStringToBytes("YELLOW SUBMARINE")
	//cipherTextBytes, _ := DecryptAES_CBC(text, key)
	cipherTextBytes := DecryptAES_CBC_Manual(text, key)

	//cipherHex := secutils.BytesToHex(cipherTextBytes)
	fmt.Printf("6. Test DecryptAES_CBC Manual: %s\n", cipherTextBytes)
}
