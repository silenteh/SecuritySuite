package aesutils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"github.com/silenteh/security_suite/secutils"
	"io"
)

var BlockSize int = aes.BlockSize

// It's important to remember that ciphertexts must be authenticated
// (i.e. by using crypto/hmac) as well as being encrypted in order to
// be secure.

func DecryptAES_EBC(cipherTextBytes, keyBytes []byte) (string, error) {

	blockSize := aes.BlockSize
	cipherTextLen := len(cipherTextBytes)
	mod := cipherTextLen % blockSize

	if mod != 0 {
		panic("ciphertext is not a multiple of the block size")
	}

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		panic(err)
	}

	numBlocks := cipherTextLen / blockSize

	blockBytes := make([]byte, blockSize)
	var buffer bytes.Buffer

	for i := 0; i < numBlocks; i++ {
		block.Decrypt(blockBytes, cipherTextBytes)
		buffer.Write(blockBytes)
		cipherTextBytes = cipherTextBytes[blockSize:]
	}

	unpadded := Unpad_Pkcs7(buffer.Bytes(), blockSize)

	text := secutils.BytesToASCIIString(unpadded)
	return text, nil
}

func EncryptAES_CBC(plainText, key string) ([]byte, error) {

	keyBytes := secutils.ASCIIStringToBytes(key)
	plainTextBytes := secutils.ASCIIStringToBytes(plainText)
	plainTextLen := len(plainTextBytes)
	mod := plainTextLen % aes.BlockSize
	// CBC mode works on blocks so plaintexts may need to be padded to the
	// next whole block. For an example of such padding, see
	// https://tools.ietf.org/html/rfc5246#section-6.2.3.2. Here we'll
	// assume that the plaintext is already of the correct length.
	if mod != 0 {
		//panic("plaintext is not a multiple of the block size")
		plainTextBytes = Pkcs7(plainTextBytes, aes.BlockSize)
		fmt.Printf("%s\n", plainTextBytes)
		plainTextLen = len(plainTextBytes)
	}

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertextBytes := make([]byte, aes.BlockSize+plainTextLen)
	iv := ciphertextBytes[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertextBytes[aes.BlockSize:], plainTextBytes)

	// It's important to remember that ciphertexts must be authenticated
	// (i.e. by using crypto/hmac) as well as being encrypted in order to
	// be secure.

	//hexCipherText := secutils.BytesToHex(ciphertextBytes)
	return ciphertextBytes, nil
}

func DecryptAES_CBC(cipherText, keyBytes []byte) (string, error) {

	fullCipherTextBytes := cipherText
	//fullCipherTextBytes := secutils.ASCIIStringToBytes(asciiCipherText) // this contains the Cipher Text and the IV
	//keyBytes := secutils.ASCIIStringToBytes(key)
	fullCipherTextLen := len(fullCipherTextBytes)
	fmt.Printf("Full lenght: %d\n", fullCipherTextLen)
	fmt.Printf("Block size: %d\n", aes.BlockSize)

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if fullCipherTextLen < aes.BlockSize {
		panic("ciphertext too short")
	}

	iv := fullCipherTextBytes[:aes.BlockSize]
	cipherTextBytes := fullCipherTextBytes[aes.BlockSize:]
	cipherTextLen := len(cipherTextBytes)
	fmt.Printf("Only cipher text lenght: %d\n", cipherTextLen)

	mod := cipherTextLen % aes.BlockSize
	fmt.Printf("Mod: %d\n", mod)

	// we need to work out the padding !
	// CBC mode always works in whole blocks.
	if mod != 0 {
		panic("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)

	decryptedBytes := make([]byte, cipherTextLen)
	// CryptBlocks can work in-place if the two arguments are the same.
	mode.CryptBlocks(decryptedBytes, cipherTextBytes)

	// If the original plaintext lengths are not a multiple of the block
	// size, padding would have to be added when encrypting, which would be
	// removed at this point. For an example, see
	// https://tools.ietf.org/html/rfc5246#section-6.2.3.2. However, it's
	// critical to note that ciphertexts must be authenticated (i.e. by
	// using crypto/hmac) before being decrypted in order to avoid creating
	// a padding oracle.

	//fmt.Printf("%s\n", cipherTextBytes)
	text := secutils.BytesToASCIIString(decryptedBytes)
	return text, nil
}

// DO NOT USE THIS !!!!!!!
// UNLESS YOU KNOW WHAT YOU ARE DOING !
func EncryptAES_CFB(dst, src, key, iv []byte) error {
	aesBlockEncrypter, err := aes.NewCipher([]byte(key))
	if err != nil {
		return err
	}
	aesEncrypter := cipher.NewCFBEncrypter(aesBlockEncrypter, iv)
	aesEncrypter.XORKeyStream(dst, src)
	return nil
}

func DecryptAES_CFB(dst, src, key, iv []byte) error {
	aesBlockDecrypter, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil
	}
	aesDecrypter := cipher.NewCFBDecrypter(aesBlockDecrypter, iv)
	aesDecrypter.XORKeyStream(dst, src)
	return nil
}

// THIS IS PROBABLY ONE OF THE BEST
func EncryptAES_CTR(plainText, key string) (string, error) {

	keyBytes := secutils.ASCIIStringToBytes(key)
	plaintextBytes := secutils.ASCIIStringToBytes(plainText)

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertextBytes.
	ciphertextBytes := make([]byte, aes.BlockSize+len(plaintextBytes))
	iv := ciphertextBytes[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(ciphertextBytes[aes.BlockSize:], plaintextBytes)
	hexCipherText := secutils.BytesToHex(ciphertextBytes)
	return hexCipherText, nil

	// CTR mode is the same for both encryption and decryption, so we can
	// also decrypt that ciphertextBytes with NewCTR.

	// plaintext2 := make([]byte, len(plaintextBytes))
	// stream = cipher.NewCTR(block, iv)
	// stream.XORKeyStream(plaintext2, ciphertextBytes[aes.BlockSize:])

}

func DecryptAES_CTR(dst, src, key, iv []byte) error {
	aesBlockDecrypter, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil
	}
	aesDecrypter := cipher.NewCTR(aesBlockDecrypter, iv)
	aesDecrypter.XORKeyStream(dst, src)
	return nil
}
