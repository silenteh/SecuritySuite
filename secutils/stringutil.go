package secutils

import (
	"bufio"
	//"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"os"
)

type Char struct {
	Byte   byte
	String string
}

var ASCII []Char = createASCII()

// creates an ASCII table with the first 32 bytes removed - mostly control chars
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

func LoadFile(filename string) []byte {

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return bytes

}

func LoadFileLines(filename string) ([]string, error) {

	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	//i := 0
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		//println(scanner.Text())
		//fmt.Printf("%d - %s\n", i, scanner.Text())
		//i++
	}
	return lines, scanner.Err()

}
