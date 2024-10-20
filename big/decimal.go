package big

import (
	"github.com/ericlagergren/decimal"
	"math/big"
)

type Decimal struct {
	inner *decimal.Big
}

func Zero() Decimal {
	return Decimal{inner: new(decimal.Big).SetFloat64(0)}
}

func isDecimalRepresentationCorrect(s string) bool {
	var b big.Float
	_, ok := b.SetString(s)
	return ok
}

func (a Decimal) String() string {
	return a.inner.String()
}

func NewFromString(s string) (Decimal, bool) {
	if !isDecimalRepresentationCorrect(s) {
		return Decimal{inner: nil}, false
	}

	val, ok := new(decimal.Big).SetString(s)

	if !ok {
		return Decimal{inner: nil}, false
	}

	return Decimal{inner: val}, true
}

func NewFromFloat64(f float64) Decimal {
	return Decimal{inner: new(decimal.Big).SetFloat64(f)}
}

func NewFromStringUnsafe(s string) Decimal {
	val, ok := NewFromString(s)

	if !ok {
		panic("Cannot build decimal.Big")
	}

	return val
}

func (a Decimal) IsInt() bool {
	return a.inner.IsInt()
}

func (a Decimal) isNil() bool {
	return a.inner == nil
}

func (a Decimal) ExactEq(b Decimal) bool {
	return a.inner.Cmp(b.inner) == 0
}

func (a Decimal) ApproxEqAbsError(b Decimal, err Decimal) (bool, Decimal) {
	diff := absError(a, b)
	return diff.inner.Cmp(err.inner) <= 0, diff
}

func (a Decimal) ApproxEqRelError(b Decimal, err Decimal) (bool, Decimal) {
	diff := relError(a, b)
	return diff.inner.Cmp(err.inner) <= 0, diff
}

func (a Decimal) InsideRange(lo, hi float64) bool {
	bad1 := a.inner.Cmp(NewFromFloat64(lo).inner) < 0
	bad2 := a.inner.Cmp(NewFromFloat64(hi).inner) > 0

	return !(bad1 || bad2)
}

func Max(a, b Decimal) Decimal {
	if a.inner == nil {
		return b
	}

	if b.inner == nil {
		return a
	}

	if a.inner.Cmp(b.inner) == 1 {
		return a
	}

	return b
}

func absError(a, b Decimal) Decimal {
	r := Zero()
	r.inner.Abs(r.inner.Sub(a.inner, b.inner))
	return r
}

func relError(a, b Decimal) Decimal {
	r := Zero()

	if b.ExactEq(r) {
		r.inner.SetInf(false)
		return r
	}

	r.inner.Abs(r.inner.Quo(r.inner.Sub(a.inner, b.inner), b.inner))
	return r
}
