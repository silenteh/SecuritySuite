package secutils

import (
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

// A function to turn a map into a FrequencyTable, then sort and return it.
func sortMapByValue(frequencyTableMap map[string]int) FrequencyTable {
	frequencyTable := make(FrequencyTable, len(frequencyTableMap))
	i := 0
	for letter, frequency := range frequencyTableMap {
		frequencyTable[i] = LetterFrequency{letter, frequency}
		i++
	}
	sort.Sort(frequencyTable)
	return frequencyTable
}

func GenerateTable(data string) FrequencyTable {

	frequencyTableMap := make(map[string]int)
	splittedString := []byte(data)

	for i := range splittedString {
		frequencyTableMap[string(splittedString[i])] += 1 // increase the fequency
	}

	frequencyTable := sortMapByValue(frequencyTableMap)
	return frequencyTable

}
