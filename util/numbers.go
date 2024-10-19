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

func AbsError(a, b *big.Float) *big.Float {
	// TODO: Unit test error calculations. They seem a bit unstable (although it works).
	// TODO: Can I do this in one liner??
	res := new(big.Float).Sub(a, b)
	res.Abs(res)
	return res
}

func RelError(a, b *big.Float) *big.Float {
	// TODO: Unit test error calculations. They seem a bit unstable (although it works).
	res := new(big.Float).Sub(a, b)
	// TODO: After doing this change, I'm not sure if I should compare it (using bigMax)
	// the same way as with the absolute error.
	// Research what's the methodology for comparing this.
	res.Quo(res, b)
	return res
}
