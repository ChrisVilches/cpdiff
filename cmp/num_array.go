package cmp

import (
	"cpdiff/util"
	"fmt"
	"math/big"
)

type NumArray struct {
	nums    []*big.Float
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

func compareNums(first, second NumArray, error *big.Float, useRelativeErr bool) ([]cmpRange, *big.Float) {
	n := len(first.nums)
	m := len(second.nums)
	size := min(n, m)
	res := []cmpRange{}
	maxErr := big.NewFloat(0)

	for i := 0; i < size; i++ {
		var v ComparisonResult
		a := first.nums[i]
		b := second.nums[i]

		if a.Cmp(b) == 0 {
			v = CmpRes.Correct
		} else {
			var err *big.Float

			if useRelativeErr {
				err = util.RelError(a, b)
			} else {
				err = util.AbsError(a, b)
			}

			if err.Cmp(error) == -1 {
				v = CmpRes.Approx
				maxErr.Set(util.BigMax(maxErr, err))
			} else {
				v = CmpRes.Incorrect
			}
		}

		res = appendCmpRange(res, cmpRange{From: i, To: i + 1, Result: v})
	}

	if n != m {
		res = appendCmpRange(res, cmpRange{From: size, To: max(n, m), Result: CmpRes.Incorrect})
	}

	return res, maxErr
}
