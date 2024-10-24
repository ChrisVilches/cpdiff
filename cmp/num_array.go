package cmp

import (
	"fmt"
	"github.com/ChrisVilches/cpdiff/big"
)

type NumArray struct {
	nums    []big.Decimal
	rawData string
}

func (n NumArray) Display() string {
	return n.rawData
}

func (n NumArray) ShortDisplay(int) string {
	if len(n.nums) == 1 {
		return n.rawData
	}

	return fmt.Sprintf("(%d numbers...)", len(n.nums))
}

func (n NumArray) compare(
	c Comparable,
	useRelativeErr bool,
	allowedError big.Decimal,
) ([]verdictRange, big.Decimal) {
	if rhs, ok := c.(NumArray); ok {
		return compareNums(n, rhs, allowedError, useRelativeErr)
	}

	return []verdictRange{{Value: Verdicts.Incorrect}}, big.Decimal{}
}

func compareNums(
	first,
	second NumArray,
	allowedError big.Decimal,
	useRelativeErr bool,
) ([]verdictRange, big.Decimal) {
	n := len(first.nums)
	m := len(second.nums)
	size := min(n, m)
	res := []verdictRange{}
	maxErr := big.NewZero()

	for i := 0; i < size; i++ {
		var v Verdict
		a := first.nums[i]
		b := second.nums[i]

		if a.ExactEq(b) {
			v = Verdicts.Correct
		} else {
			approx := false
			var err big.Decimal

			if useRelativeErr {
				approx, err = a.ApproxEqRelError(b, allowedError)
			} else {
				approx, err = a.ApproxEqAbsError(b, allowedError)
			}

			if approx {
				v = Verdicts.Approx
				maxErr = big.Max(maxErr, err)
			} else {
				v = Verdicts.Incorrect
			}
		}

		res = appendVerdictRange(
			res,
			verdictRange{From: i, To: i + 1, Value: v},
		)
	}

	if n != m {
		res = appendVerdictRange(
			res,
			verdictRange{From: size, To: max(n, m), Value: Verdicts.Incorrect},
		)
	}

	return res, maxErr
}
