package secutils

import (
	"testing"
)

func TestTrasposeMatrix(t *testing.T) {
	first := []byte{byte(48), byte(49), byte(50)}
	second := []byte{byte(51), byte(52), byte(53)}

	m := make([][]byte, 2)
	m[0] = first
	m[1] = second
	//PrintArrayOfArray(m)
	transposed := TrasposeArray(m)
	//PrintArrayOfArray(transposed)
	if transposed[0][0] != byte(48) && transposed[0][1] != byte(51) {
		t.Error("Matrix transposition error  !")
	}
	if transposed[1][0] != byte(49) && transposed[1][1] != byte(52) {
		t.Error("Matrix transposition error  !")
	}
	if transposed[2][0] != byte(50) && transposed[2][1] != byte(53) {
		t.Error("Matrix transposition error  !")
	}

}

func TestDivideArrayInBlocks(t *testing.T) {
	bytes := []byte{byte(48), byte(49), byte(50), byte(51), byte(52), byte(53), byte(54), byte(55)}

	array := DivideArrayInBlocks(bytes, 3)
	if len(array) != 3 {
		t.Error("Matrix block division in blocks error  !")
	}
}
