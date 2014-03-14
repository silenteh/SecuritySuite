package secutils

import (
	//"fmt"
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

// VERY inefficient but uses less memory than the matrix approach good for small strings only !
// func LevenshteinDistance(s string, len_s int, t string, len_t int) int {

// 	/* base case: empty strings */
// 	if len_s == 0 {
// 		return len_t
// 	}
// 	if len_t == 0 {
// 		return len_s
// 	}

// 	sBytes := ASCIIStringToBytes(s)
// 	tBytes := ASCIIStringToBytes(t)

// 	 test if last characters of the strings match
// 	var cost int
// 	if sBytes[len_s-1] == tBytes[len_t-1] {
// 		cost = 0
// 	} else {
// 		cost = 1
// 	}

// 	return Min(
// 		LevenshteinDistance(s, len_s-1, t, len_t)+1,
// 		LevenshteinDistance(s, len_s, t, len_t-1)+1,
// 		LevenshteinDistance(s, len_s-1, t, len_t-1)+cost)

// }

func LevenshteinDistance(s, t string) int {
	// degenerate cases
	len_s := len(s)
	len_t := len(t)

	if s == t {
		return 0
	}
	if len_s == 0 {
		return len_t
	}
	if len_t == 0 {
		return len_s
	}

	// create two work vectors of integer distances
	v0 := make([]int, len_t+1) //new int[t.Length + 1]; // int[]
	v1 := make([]int, len_t+1) // int[]

	// initialize v0 (the previous row of distances)
	// this row is A[0][i]: edit distance for an empty s
	// the distance is just the number of characters to delete from t
	for i := 0; i < len_t+1; i++ {
		v0[i] = i
	}

	for i := 0; i < len_s; i++ {
		// calculate v1 (current row distances) from the previous row v0

		// first element of v1 is A[i+1][0]
		//   edit distance is delete (i+1) chars from s to match empty t
		v1[0] = i + 1

		// use formula to fill in the rest of the row
		for j := 0; j < len_t; j++ {

			var cost int
			if s[i] == t[j] {
				cost = 0
			} else {
				cost = 1
			}

			v1[j+1] = Min(v1[j]+1, v0[j+1]+1, v0[j]+cost)
		}

		// copy v1 (current row) to v0 (previous row) for next iteration
		for j := 0; j < len_t+1; j++ {
			v0[j] = v1[j]
		}
	}

	return v1[len_t]
}

// int is a signed integer in Golang
func Min(x, y, z int) int {

	var currentMin int
	if x <= y {
		currentMin = x
	} else {
		currentMin = y
	}

	var finalMin int
	if currentMin <= z {
		finalMin = currentMin
	} else {
		finalMin = z
	}

	return finalMin
}
