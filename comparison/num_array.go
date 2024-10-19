package comparison

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

func compareNums(first, second NumArray, error *big.Float, useRelativeErr bool) (ComparisonResult, *big.Float) {
	maxErr := big.NewFloat(0)

	if len(first.nums) != len(second.nums) {
		return ComparisonResults.Incorrect, maxErr
	}

	approx := false
	ok := true

	for i := range len(first.nums) {
		a := first.nums[i]
		b := second.nums[i]

		if a.Cmp(b) == 0 {
			continue
		}

		var err *big.Float

		if useRelativeErr {
			err = util.RelError(a, b)
		} else {
			err = util.AbsError(a, b)
		}

		maxErr.Set(util.BigMax(maxErr, err))

		if err.Cmp(error) == -1 {
			approx = true
		} else {
			ok = false
		}
	}

	if !ok {
		return ComparisonResults.Incorrect, maxErr
	} else if approx {
		return ComparisonResults.Approx, maxErr
	} else {
		return ComparisonResults.Correct, maxErr
	}
}
