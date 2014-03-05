package secutils

import (
	"fmt"
	"math/big"
	"testing"
)

func TestXor(t *testing.T) {

	// 1 xor 1 = 0
	var res_1xor1 *big.Int
	res_1xor1 = Xor("1", "1")
	if res_1xor1.Sign() != 0 {
		t.Error("Expected 0, got ", res_1xor1)
	}

	// 0 xor 0 = 0
	var res_0xor0 *big.Int
	res_0xor0 = Xor("0", "0")
	if res_0xor0.Sign() != 0 {
		t.Error("Expected 0, got ", res_0xor0)
	}

	// 1 xor 0 = 1
	var res_1xor0 *big.Int
	res_1xor0 = Xor("1", "0")
	if res_1xor0.Sign() != 1 {
		t.Error("Expected 1, got ", res_1xor0)
	}

	// 0 xor 1 = 1
	var res_0xor1 *big.Int
	res_0xor1 = Xor("0", "1")
	if res_0xor1.Sign() != 1 {
		t.Error("Expected 1, got ", res_0xor1)
	}

	fmt.Println("Test Xor: OK")

}
