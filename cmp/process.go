package cmp

import (
	"cpdiff/big"
)

func findGlobalResult(cmpRanges []verdictRange) Verdict {
	approx := false

	for _, r := range cmpRanges {
		if r.Result == Verdicts.Incorrect {
			return Verdicts.Incorrect
		}

		if r.Result == Verdicts.Approx {
			approx = true
		}
	}

	if approx {
		return Verdicts.Approx
	}

	return Verdicts.Correct
}

func bothNumbers(lhs, rhs Comparable) bool {
	return lhs.Type() == ComparableTypes.NumArray &&
		rhs.Type() == ComparableTypes.NumArray
}

func applyStringFallbackHeuristic(lhs, rhs *Comparable, line1, line2 string) {
	// Heuristic: when one side has numbers and the other side has strings,
	// turn both into strings.
	// When strings are compared, each character is colored
	// green/red depending if it matches or not.
	// This case would get completely red simply because the
	// types are different:
	// 1011 (treated as number)
	// 0011 (treated as string)
	// In this case, the user may be trying to compare binary strings,
	// in which case it's best to show the differences bit by bit.
	// If both begin with a non-zero digit, then the only solution is to
	// disable numeric conversion (CLI flag).
	if bothNumbers(*lhs, *rhs) {
		return
	}

	*lhs = newComparableNonNumeric(line1)
	*rhs = newComparableNonNumeric(line2)
}

func Process(
	lhsCh,
	rhsCh chan string,
	allowedError big.Decimal,
	useRelativeErr bool,
	useNumbers bool,
	entriesCh chan ComparisonEntry,
) {
	for {
		line1, ok1 := <-lhsCh
		line2, ok2 := <-rhsCh

		if !ok1 && !ok2 {
			break
		}

		var lhs, rhs Comparable

		if useNumbers {
			lhs = newComparable(line1)
			rhs = newComparable(line2)

			applyStringFallbackHeuristic(&lhs, &rhs, line1, line2)
		} else {
			lhs = newComparableNonNumeric(line1)
			rhs = newComparableNonNumeric(line2)
		}

		entriesCh <- newComparisonEntry(lhs, rhs, useRelativeErr, allowedError)
	}

	close(entriesCh)
}
