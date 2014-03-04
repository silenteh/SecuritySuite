package secutils

import (
	"log"
)

func HammingDistance(s1, s2 string) int {

	// check if the string have the same length
	if len(s1) != len(s2) {
		log.Fatalf("%s and %s have different length", s1, s2)
	}

	// define the variable to store the hamming weight
	dist := 0
	// xor them
	val := Xor(s1, s2).BitLen()
	// example of debug format
	//fmt.Printf("%d\n", val)
	// calculate the hamming weight
	for val > 0 {
		dist = dist + 1
		val &= val - 1
		//fmt.Printf("%b\n", val)
	}
	return dist
}
