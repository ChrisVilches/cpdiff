package util

import "math/big"

func BigFloatOutsideRange(bigFloat *big.Float, lo, hi float64) bool {
	bad1 := bigFloat.Cmp(new(big.Float).SetFloat64(lo)) <= 0
	bad2 := bigFloat.Cmp(new(big.Float).SetFloat64(hi)) >= 0

	return bad1 || bad2
}

func BigMax(a, b *big.Float) *big.Float {
	if a == nil {
		return b
	}

	if b == nil {
		return a
	}

	if a == nil && b == nil {
		return nil
	}

	if a.Cmp(b) == 1 {
		return a
	}

	return b
}

func AbsError(a, b *big.Float) *big.Float {
	var r big.Float

	return r.Abs(r.Sub(a, b))
}

func RelError(a, b *big.Float) *big.Float {
	r := big.NewFloat(0)

	if b.Cmp(r) == 0 {
		return r.SetInf(false)
	}

	return r.Abs(r.Quo(r.Sub(a, b), b))

	// TODO: After doing this change, I'm not sure
	//    if I should compare it (using bigMax)
	// the same way as with the absolute error.
	// Research what's the methodology for comparing this.
}
