package secutils

import (
	"io/ioutil"
	"log"
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

	bytes, err := ioutil.ReadFile("sherlock.txt")
	if err != nil {
		log.Fatal(err)
	}
	return createFrequencyTableMap(bytes)
}

func KeyScore(key string) int {

	bytes := []byte(key)
	bytesLen := len(bytes)
	total := 0

	for i := range bytes {
		b := bytes[i]
		total += CryptoMap[string(b)]
	}

	return (total / bytesLen)
}
