package aesutils

import (
	"bytes"
	//"fmt"
	"log"
)

func Pkcs7(data []byte, blockSize int) []byte {

	dataLen := len(data)

	//if dataLen < blockSize {
	//	log.Fatal("The lenght of the data cannot be < than the block size !!!")
	//}

	mod := dataLen % blockSize

	// it does not need padding
	if mod == 0 {
		return data
	}

	totalPadding := blockSize - mod
	var buffer bytes.Buffer
	buffer.Write(data)
	for i := 0; i < totalPadding; i++ {
		buffer.WriteByte(byte(totalPadding))
	}

	return buffer.Bytes()

}

func Unpad_Pkcs7(data []byte, blockSize int) []byte {

	dataLen := len(data)

	if dataLen < blockSize {
		log.Fatal("The lenght of the data cannot be less than the block size !!!")
	}

	lastByte := data[dataLen-1]
	num := int(lastByte)

	unpaddedData := data[0 : dataLen-num]
	return unpaddedData

}

func IsPkcs7(data []byte, blockSize int) bool {

	dataLen := len(data)

	lastByte := data[dataLen-1]
	num := int(lastByte)

	if num >= blockSize {
		return false
	}

	start := dataLen - num
	end := dataLen
	subData := data[start:end]

	total := num * num

	//fmt.Printf("****** num: %d -  total: %d\n", num, total)

	currentTotal := 0
	for index := range subData {
		subNum := int(subData[index])
		currentTotal = currentTotal + subNum
	}

	if currentTotal == total {
		return true
	}

	return false

}
