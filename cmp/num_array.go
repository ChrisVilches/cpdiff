package cmp

import (
	"cpdiff/big"
	"fmt"
)

type NumArray struct {
	nums    []big.Decimal
	rawData string
}

func (NumArray) Type() ComparableType {
	return ComparableTypes.NumArray
}

func (n NumArray) HasRealNumbers() bool {
	for _, num := range n.nums {
		if !num.IsInt() {
			return true
		}
	}

	return false
}

func (n NumArray) Display() string {
	return n.rawData
}

func (n NumArray) ShortDisplay() string {
	if len(n.nums) == 1 {
		return n.rawData
	}

	return fmt.Sprintf("(%d numbers...)", len(n.nums))
}

func compareNums(
	first,
	second NumArray,
	allowedError big.Decimal,
	useRelativeErr bool,
) ([]cmpRange, big.Decimal) {
	n := len(first.nums)
	m := len(second.nums)
	size := min(n, m)
	res := []cmpRange{}
	maxErr := big.Zero()

	for i := 0; i < size; i++ {
		var v ComparisonResult
		a := first.nums[i]
		b := second.nums[i]

		if a.ExactEq(b) {
			v = CmpRes.Correct
		} else {
			approx := false
			var err big.Decimal

			if useRelativeErr {
				approx, err = a.ApproxEqRelError(b, allowedError)
			} else {
				approx, err = a.ApproxEqAbsError(b, allowedError)
			}

			if approx {
				v = CmpRes.Approx
				maxErr = big.Max(maxErr, err)
			} else {
				v = CmpRes.Incorrect
			}
		}

		res = appendCmpRange(res, cmpRange{From: i, To: i + 1, Result: v})
	}

	if n != m {
		res = appendCmpRange(
			res,
			cmpRange{From: size, To: max(n, m), Result: CmpRes.Incorrect},
		)
	}

	return res, maxErr
}
