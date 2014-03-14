package secutils

import (
	"fmt"
)

type KeyMatch struct {
	KeySize  int
	Distance float64
}

type KeyMatches []KeyMatch

func (ft KeyMatches) Swap(i, j int)      { ft[i], ft[j] = ft[j], ft[i] }
func (ft KeyMatches) Len() int           { return len(ft) }
func (ft KeyMatches) Less(i, j int) bool { return ft[i].Distance < ft[j].Distance }

type Row []byte
type Matrix []Row

func TrasposeArray(array [][]byte) [][]byte {
	transposed := make([][]byte, len(array[0]))

	// init the transposed sub arrays
	for x, _ := range transposed {
		transposed[x] = make([]byte, len(array))
	}

	// loop through the original array
	for y, s := range array {
		for x, e := range s {
			transposed[x][y] = e
		}
	}
	return transposed
}

func TrasposeMatrix(array [][]byte) Matrix {
	r := make(Matrix, len(array[0]))
	// init
	for x, _ := range r {
		r[x] = make(Row, len(array))
	}
	// fill
	for y, s := range array {
		for x, e := range s {
			r[x][y] = e
		}
	}
	return r
}

func PrintMatrix(m Matrix) {
	for _, b := range m {
		fmt.Println(string(b))
	}
}

func PrintArrayOfArray(m [][]byte) {
	for _, b := range m {
		fmt.Println(string(b))
	}
}

func DivideArrayInBlocks(array []byte, blockSize int) [][]byte {

	arraysize := len(array)
	size := arraysize / blockSize
	mod := arraysize % blockSize
	if mod != 0 {
		size = size + 1
	}

	blocks := make([][]byte, size)

	for index := range blocks {
		if mod != 0 && index == size-1 {
			blocks[index] = array[arraysize-mod:]
		} else {
			starts := index * blockSize
			ends := starts + blockSize
			blocks[index] = array[starts:ends]
		}
	}

	return blocks
}
