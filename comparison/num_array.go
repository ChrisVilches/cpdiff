package comparison

import (
	"fmt"
	"math/big"
)

type NumArray struct {
	nums    []big.Float
	rawData string
}

func (*NumArray) Type() ComparableType {
	return ComparableTypes.NumArray
}

func (n *NumArray) HasRealNumbers() bool {
	for _, num := range n.nums {
		if !num.IsInt() {
			return true
		}
	}

	return false
}

// TODO: Should this be reference?
func (n *NumArray) Display() string {
	return n.rawData
}

func (n *NumArray) ShortDisplay() string {
	if len(n.nums) == 1 {
		return n.rawData
	}

	return fmt.Sprintf("(%d numbers...)", len(n.nums))
}

func compareNums(first, second *NumArray, error *big.Float, relativeErr bool) (ComparisonResult, big.Float) {
	// biggestDifference := new(big.Float)
	// TODO: I think this prevents doing heap operations? Just return the bare object?
	var biggestDifference big.Float

	if len(first.nums) != len(second.nums) {
		return ComparisonResults.Incorrect, biggestDifference
	}

	approx := false
	ok := true

	for i := range len(first.nums) {
		a := first.nums[i]
		b := second.nums[i]

		if a.Cmp(&b) == 0 {
			continue
		}

		// TODO: Unit test error calculations. They seem a bit unstable (although it works).
		diff := new(big.Float)
		diff.Sub(&a, &b)

		if relativeErr {
			// TODO: This will crash when b is zero!!!! How to calculate it correctly??
			diff.Quo(diff, &b)
		}

		diff.Abs(diff)

		if biggestDifference.Cmp(diff) == -1 {
			biggestDifference.Set(diff)
		}

		if diff.Cmp(error) == -1 {
			approx = true
		} else {
			ok = false
		}
	}

	if !ok {
		return ComparisonResults.Incorrect, biggestDifference
	}

	if approx {
		return ComparisonResults.Approx, biggestDifference
	} else {
		return ComparisonResults.Correct, biggestDifference
	}
}
