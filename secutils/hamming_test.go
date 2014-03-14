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
		t.Error("Expected 3, got ", resultRoses)
	}

	var resultBin int
	expectedBin := 2
	resultBin = HammingDistance("1011101", "1001001")
	if resultBin != expectedBin {
		t.Error("Expected 2, got ", resultBin)
	}

	var resultDec int
	expectedDec := 3
	resultDec = HammingDistance("2173896", "2233796")
	if resultDec != expectedDec {
		t.Error("Expected 3, got ", resultDec)
	}

	var resultWokka int
	expectedWokka := 4
	resultWokka = HammingDistance("wokka wokka!!!", "this is a test")
	if resultWokka != expectedWokka {
		t.Error("Expected 4, got ", resultWokka)
	}

	fmt.Println("Test HammingDistance: OK")
}

func TestMin(t *testing.T) {
	expected := 3
	res := Min(3, 4, 5)
	if res != expected {
		t.Error("Expected 3, got ", res)
	}
	fmt.Println("Test Min: OK")
}

func TestLevenshteinDistance(t *testing.T) {
	k := "kitten"
	s := "sitting"
	expected := 3
	res := LevenshteinDistance(k, s)
	//fmt.Printf("Leven: %d\n", res)
	if res != expected {
		t.Error("Expected 3, got ", res)
	}

	k1 := "this is a test"
	s1 := "wokka wokka!!!"
	expected1 := 14
	res1 := LevenshteinDistance(k1, s1)
	if res1 != expected1 {
		t.Error("Expected 3, got ", res)
	}
	fmt.Println("Test LevenshteinDistance: OK")
}
