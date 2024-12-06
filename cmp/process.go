package cmp

import (
	"sync/atomic"

	"github.com/ChrisVilches/cpdiff/big"
)

type ComparisonEntry struct {
	LHS           Comparable
	RHS           Comparable
	Verdict       Verdict
	VerdictRanges []verdictRange
	MaxErr        big.Decimal
}

func newComparisonEntry(
	lhs,
	rhs Comparable,
	r []verdictRange,
	maxErr big.Decimal,
) ComparisonEntry {
	return ComparisonEntry{
		LHS:           lhs,
		RHS:           rhs,
		MaxErr:        maxErr,
		VerdictRanges: r,
		Verdict:       findGlobalVerdict(r),
	}
}

func findGlobalVerdict(cmpRanges []verdictRange) Verdict {
	approx := false

	for _, r := range cmpRanges {
		if r.Value == Verdicts.Incorrect {
			return Verdicts.Incorrect
		}

		if r.Value == Verdicts.Approx {
			approx = true
		}
	}

	if approx {
		return Verdicts.Approx
	}

	return Verdicts.Correct
}

func bothNumbers(lhs, rhs Comparable) bool {
	_, ok := lhs.(NumArray)
	return ok && SameType(lhs, rhs)
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
	// If both begin with a non-zero digit but the user still wants
	// to compare each digit, then the only solution is to disable
	// numeric conversion (CLI flag).
	if bothNumbers(*lhs, *rhs) {
		return
	}

	*lhs = newComparableNonNumeric(line1)
	*rhs = newComparableNonNumeric(line2)
}

func Process(
	lhsCh,
	rhsCh <-chan string,
	allowedError big.Decimal,
	useRelativeErr bool,
	useNumbers bool,
	chSize int,
	aborted *atomic.Bool,
) <-chan ComparisonEntry {
	entries := make(chan ComparisonEntry, chSize)

	go func() {
		for {
			line1, ok1 := <-lhsCh
			line2, ok2 := <-rhsCh

			if (!ok1 && !ok2) || aborted.Load() {
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

			ranges, maxErr := lhs.compare(rhs, useRelativeErr, allowedError)

			entries <- newComparisonEntry(lhs, rhs, ranges, maxErr)
		}

		close(entries)
	}()

	return entries
}
