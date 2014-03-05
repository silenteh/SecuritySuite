package secutils

import (
	"fmt"
	"testing"
)

func TestHammingDistance(t *testing.T) {
	var resultRoses int
	expectedRoses := 3
	resultRoses = HammingDistance("toned", "roses")
	if resultRoses != expectedRoses {
		t.Error("Expected 3, got ", expectedRoses)
	}

	var resultBin int
	expectedBin := 2
	resultBin = HammingDistance("1011101", "1001001")
	if resultBin != expectedBin {
		t.Error("Expected 2, got ", expectedBin)
	}

	var resultDec int
	expectedDec := 3
	resultDec = HammingDistance("2173896", "2233796")
	if resultDec != expectedDec {
		t.Error("Expected 3, got ", expectedDec)
	}

	fmt.Println("Test HammingDistance: OK")
}
