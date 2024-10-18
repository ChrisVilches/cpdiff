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

func compareNums(first, second NumArray, error *big.Float, relativeErr bool) (ComparisonResult, *big.Float) {
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

		// TODO: Unit test error calculations. They seem a bit unstable (although it works).
		diff := new(big.Float).Sub(a, b)

		if relativeErr {
			// TODO: After doing this change, I'm not sure if I should compare it (using bigMax)
			// the same way as with the absolute error.
			// Research what's the methodology for comparing this.
			diff.Quo(diff, b)
		}

		diff.Abs(diff)

		maxErr.Set(util.BigMax(maxErr, diff))

		if diff.Cmp(error) == -1 {
			approx = true
		} else {
			ok = false
		}
	}

	if !ok {
		return ComparisonResults.Incorrect, maxErr
	}

	if approx {
		return ComparisonResults.Approx, maxErr
	} else {
		return ComparisonResults.Correct, maxErr
	}
}
