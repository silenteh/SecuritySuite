package secutils

import "testing"

func TestGenerateTable(t *testing.T) {

	data := " silenteh "
	ft := GenerateTable(data)
	if ft[0].Frequency != 2 || ft[0].Letter != " " {
		t.Error("Frequency table calculation is wrong")
	}
}
