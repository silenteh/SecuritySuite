package secutils

import (
	"fmt"
	"sort"
)

type LetterFrequency struct {
	Letter    string
	Frequency int
}

type FrequencyTable []LetterFrequency

func (ft FrequencyTable) Swap(i, j int)      { ft[i], ft[j] = ft[j], ft[i] }
func (ft FrequencyTable) Len() int           { return len(ft) }
func (ft FrequencyTable) Less(i, j int) bool { return ft[i].Frequency > ft[j].Frequency }

var CryptoMap map[string]int = preGenerateTableMapForCrypto()
var CryptoFrequencyTable FrequencyTable = SortMapByValue(CryptoMap)

// A function to turn a map into a FrequencyTable, then sort and return it.
func SortMapByValue(frequencyTableMap map[string]int) FrequencyTable {
	frequencyTable := make(FrequencyTable, len(frequencyTableMap))
	i := 0
	for letter, frequency := range frequencyTableMap {
		frequencyTable[i] = LetterFrequency{letter, frequency}
		i++
	}
	sort.Sort(frequencyTable)
	return frequencyTable
}

func createFrequencyTableMap(bytes []byte) map[string]int {
	frequencyTableMap := make(map[string]int)

	for i := range bytes {
		frequencyTableMap[string(bytes[i])] += 1 // increase the fequency
	}

	return frequencyTableMap
}

func createFrequencyTableMapFromStrings(stringList []string) map[string]int {
	frequencyTableMap := make(map[string]int)

	for i := range stringList {
		frequencyTableMap[string(stringList[i])] += 1 // increase the fequency
	}

	return frequencyTableMap
}

func createFrequencyTable(bytes []byte) FrequencyTable {
	frequencyTableMap := createFrequencyTableMap(bytes)

	frequencyTable := SortMapByValue(frequencyTableMap)
	return frequencyTable
}

func GenerateTable(data string) FrequencyTable {

	bytes := []byte(data)
	return createFrequencyTable(bytes)

}

func preGenerateTableMapForCrypto() map[string]int {

	bytes := LoadFile("../sherlock.txt")
	return createFrequencyTableMap(bytes)
}

func KeyScore(text string) int {
	bytes := []byte(text)
	bytesLen := len(bytes)
	total := 0

	for i := range bytes {
		b := bytes[i]
		total += CryptoMap[string(b)]
	}
	return (total / bytesLen)
}

func BlocksFrequency(bytes []byte, blockSize int) FrequencyTable {

	blocks := len(bytes) / blockSize
	mod := len(bytes) % blockSize
	if mod != 0 {
		fmt.Println("Cannot split the string because it is not a multiple of the block size")
		return nil
	}

	stringBlocks := make([]string, blocks)

	for block := range stringBlocks {
		start := block * blockSize
		end := start + blockSize
		stringBlocks[block] = BytesToHex(bytes[start:end])
		//fmt.Printf("%s\n", stringBlocks[block])
	}

	freqMap := createFrequencyTableMapFromStrings(stringBlocks)
	return SortMapByValue(freqMap)
}
