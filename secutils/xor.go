package secutils

import (
	"math/big"
)

func Xor(s1, s2 string) *big.Int {

	s1Buf := new(big.Int).SetBytes([]byte(s1))
	s2Buf := new(big.Int).SetBytes([]byte(s2))
	result := new(big.Int).Xor(s1Buf, s2Buf)
	return result
}
