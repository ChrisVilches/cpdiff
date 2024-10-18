package util

import "math/big"

func BigFloatOutsideRange(bigFloat *big.Float, lo, hi float64) bool {
	bad1 := bigFloat.Cmp(new(big.Float).SetFloat64(lo)) <= 0
	bad2 := bigFloat.Cmp(new(big.Float).SetFloat64(hi)) >= 0
	return bad1 || bad2
}

func BigMax(a, b *big.Float) *big.Float {
	if a.Cmp(b) == 1 {
		return a
	} else {
		return b
	}
}
