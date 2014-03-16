package aesutils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	//"fmt"
	"github.com/silenteh/security_suite/secutils"
	"io"
)

var BlockSize int = aes.BlockSize

// It's important to remember that ciphertexts must be authenticated
// (i.e. by using crypto/hmac) as well as being encrypted in order to
// be secure.

func GenerateRandomKey() []byte {
	key := make([]byte, BlockSize)
	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		panic(err)
	}
	return key
}

//#######################################################################################################################
//#######################################################################################################################
// GOLANG NATIVE CRYPTO FUNCTIONS !!!
// USE THESE !!!
func DecryptAES_ECB(cipherTextBytes, keyBytes []byte) (string, error) {

	cipherTextLen := len(cipherTextBytes)
	mod := cipherTextLen % BlockSize

	if mod != 0 {
		panic("ciphertext is not a multiple of the block size")
	}

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		panic(err)
	}

	numBlocks := cipherTextLen / BlockSize

	blockBytes := make([]byte, BlockSize)
	var buffer bytes.Buffer

	for i := 0; i < numBlocks; i++ {
		block.Decrypt(blockBytes, cipherTextBytes)
		buffer.Write(blockBytes)
		cipherTextBytes = cipherTextBytes[BlockSize:]
	}

	decrypedData := buffer.Bytes()
	if IsPkcs7(decrypedData, BlockSize) {
		decrypedData = Unpad_Pkcs7(decrypedData, BlockSize)
	}

	text := secutils.BytesToASCIIString(decrypedData)
	return text, nil
}

func EncryptAES_CBC(plainText, key string) ([]byte, error) {

	keyBytes := secutils.ASCIIStringToBytes(key)
	plainTextBytes := secutils.ASCIIStringToBytes(plainText)
	plainTextLen := len(plainTextBytes)
	mod := plainTextLen % BlockSize
	// CBC mode works on blocks so plaintexts may need to be padded to the
	// next whole block. For an example of such padding, see
	// https://tools.ietf.org/html/rfc5246#section-6.2.3.2. Here we'll
	// assume that the plaintext is already of the correct length.
	if mod != 0 {
		//panic("plaintext is not a multiple of the block size")
		plainTextBytes = Pkcs7(plainTextBytes, BlockSize)
		//fmt.Printf("%s\n", plainTextBytes)
		plainTextLen = len(plainTextBytes)
	}

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertextBytes := make([]byte, BlockSize+plainTextLen)
	iv := ciphertextBytes[:BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertextBytes[BlockSize:], plainTextBytes)

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
	// fmt.Printf("Full lenght: %d\n", fullCipherTextLen)
	// fmt.Printf("Block size: %d\n", blockSize)

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if fullCipherTextLen < BlockSize {
		panic("ciphertext too short")
	}

	iv := fullCipherTextBytes[:BlockSize]
	cipherTextBytes := fullCipherTextBytes[BlockSize:]
	cipherTextLen := len(cipherTextBytes)
	//fmt.Printf("Only cipher text lenght: %d\n", cipherTextLen)

	mod := cipherTextLen % BlockSize
	//fmt.Printf("Mod: %d\n", mod)

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
	ciphertextBytes := make([]byte, BlockSize+len(plaintextBytes))
	iv := ciphertextBytes[:BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(ciphertextBytes[BlockSize:], plaintextBytes)
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

//#######################################################################################################################
//#######################################################################################################################
// DO NOT USE THESE !!!!!!!!!!!!!
// This function encrypts a single block
// it can be used to encrypt in ECB mode or in CBC mode
func Encrypt_Block(blockBytes, keyBytes []byte) []byte {

	keyLen := len(keyBytes)
	blockLen := len(blockBytes)
	mod := blockLen % keyLen

	// ECB and CBC requires PADDING before encrypting !
	if mod != 0 {
		panic("ciphertext is not a multiple of the block size")
	}

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		panic(err)
	}

	cipherBytes := make([]byte, BlockSize)
	block.Encrypt(cipherBytes, blockBytes)

	return cipherBytes

}

// This function decrypts a single block
// it can be used to decrypt in ECB mode or in CBC mode
func Decrypt_Block(blockBytes, keyBytes []byte) []byte {

	keyLen := len(keyBytes)
	blockLen := len(blockBytes)
	mod := blockLen % keyLen

	// ECB and CBC requires PADDING before encrypting !
	if mod != 0 {
		panic("ciphertext is not a multiple of the block size")
	}

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		panic(err)
	}

	cipherBytes := make([]byte, BlockSize)
	block.Decrypt(cipherBytes, blockBytes)

	return cipherBytes

}

// This should use a different key for every AES BLOCK !!!!
// Otherwise it is INSECURE!
func EncryptAES_ECB_Manual(plainText, key string) []byte {

	keyBytes := secutils.ASCIIStringToBytes(key)
	plainTextBytes := secutils.ASCIIStringToBytes(plainText)
	blocks := len(plainTextBytes) / BlockSize
	mod := len(plainTextBytes) % BlockSize

	if mod != 0 {
		blocks++
	}

	var buffer bytes.Buffer
	for block := 0; block < blocks; block++ {
		// this is the last block which may need padding
		if block == blocks-1 {
			start := block * BlockSize
			data := plainTextBytes[start:]

			// just make sure it is padded
			// the Pkcs7 function takes care of double check it
			data = Pkcs7(data, BlockSize)

			cipherBlock := Encrypt_Block(data, keyBytes)
			buffer.Write(cipherBlock)
		} else {
			start := block * BlockSize
			end := start + BlockSize
			data := plainTextBytes[start:end]
			cipherBlock := Encrypt_Block(data, keyBytes)
			buffer.Write(cipherBlock)
		}
	}

	return buffer.Bytes()

}
func DecryptAES_ECB_Manual(cipherTextBytes, keyBytes []byte) []byte {

	blocks := len(cipherTextBytes) / BlockSize
	mod := len(cipherTextBytes) % BlockSize

	if mod != 0 {
		panic("ciphertext is not a multiple of the block size ")
	}

	var buffer bytes.Buffer
	for block := 0; block < blocks; block++ {
		start := block * BlockSize
		end := start + BlockSize
		data := cipherTextBytes[start:end]
		cipherBlock := Decrypt_Block(data, keyBytes)
		buffer.Write(cipherBlock)
	}

	decryptedBytes := buffer.Bytes()

	if IsPkcs7(decryptedBytes, BlockSize) {
		decryptedBytes = Unpad_Pkcs7(decryptedBytes, BlockSize)
	}

	return decryptedBytes

}

func EncryptAES_CBC_Manual(plainText, key string) []byte {

	keyBytes := secutils.ASCIIStringToBytes(key)
	plainTextBytes := secutils.ASCIIStringToBytes(plainText)
	blocks := len(plainTextBytes) / BlockSize
	mod := len(plainTextBytes) % BlockSize

	if mod != 0 {
		blocks++
	}

	iv := make([]byte, BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	var buffer bytes.Buffer
	previousBlock := iv
	buffer.Write(iv)
	for block := 0; block < blocks; block++ {
		// this is the last block which may need padding

		if block == blocks-1 {
			start := block * BlockSize
			data := plainTextBytes[start:]

			// just make sure it is padded
			// the Pkcs7 function takes care of double check it
			originalData := Pkcs7(data, BlockSize)
			// XOR with the previous block
			xored := secutils.XorBytes(originalData, previousBlock)

			cipherBlock := Encrypt_Block(xored, keyBytes)
			buffer.Write(cipherBlock)
		} else {
			start := block * BlockSize
			end := start + BlockSize
			// XOR with the previous block
			originalData := plainTextBytes[start:end]
			xored := secutils.XorBytes(originalData, previousBlock)
			// Encrypt with the XORED data and assign the result to the previous block for the next cycle !
			cipherBlock := Encrypt_Block(xored, keyBytes)

			// #### THIS IS REALLY IMPORTANT !!!!!
			// otherwise it is completely broken !!!!
			previousBlock = cipherBlock

			buffer.Write(cipherBlock)
		}
	}

	return buffer.Bytes()

}
func DecryptAES_CBC_Manual(cipherTextBytes, keyBytes []byte) []byte {

	blocks := (len(cipherTextBytes) / BlockSize) - 1 // we have to remove the IV
	//mod := len(cipherTextBytes) % BlockSize

	//if mod != 0 {
	//	panic("ciphertext is not a multiple of the block size ")
	//}

	iv := cipherTextBytes[:BlockSize]
	remainingBytes := cipherTextBytes[BlockSize:]
	previousBlock := iv

	var buffer bytes.Buffer
	for block := 0; block < blocks; block++ {

		blockBytes := remainingBytes[:BlockSize]
		remainingBytes = remainingBytes[BlockSize:]

		decryptedBlock := Decrypt_Block(blockBytes, keyBytes)
		xored := secutils.XorBytes(previousBlock, decryptedBlock)

		buffer.Write(xored)
		previousBlock = blockBytes
	}

	decryptedBytes := buffer.Bytes()

	if IsPkcs7(decryptedBytes, BlockSize) {
		decryptedBytes = Unpad_Pkcs7(decryptedBytes, BlockSize)
	}

	return decryptedBytes
}
