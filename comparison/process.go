package comparison

import (
	"log"
	"math/big"
)

type ComparisonEntry struct {
	Lhs            Comparable
	Rhs            Comparable
	CmpRes         ComparisonResult
	HasRealNumbers bool
	MaxErr         *big.Float
}

func handleNumArrays(lhs, rhs NumArray, relativeErr bool, error *big.Float) (ComparisonResult, *big.Float, bool) {
	cmp, diff := compareNums(lhs, rhs, error, relativeErr)

	return cmp, diff, lhs.HasRealNumbers() || lhs.HasRealNumbers()
}

func Process(
	lhsCh,
	rhsCh chan string,
	error *big.Float,
	relativeErr bool,
	entriesCh chan ComparisonEntry,
) {
	for {
		l, ok1 := <-lhsCh
		r, ok2 := <-rhsCh

		if !ok1 && !ok2 {
			break
		}

		lhs := LineToComparable(l)
		rhs := LineToComparable(r)

		var cmp ComparisonResult
		maxErr := big.NewFloat(0)
		hasRealNumbers := false

		if lhs.Type() != rhs.Type() {
			cmp = ComparisonResults.Incorrect
		} else if lhs.Type() == ComparableTypes.NumArray {
			cmp, maxErr, hasRealNumbers = handleNumArrays(
				lhs.(NumArray),
				rhs.(NumArray),
				relativeErr,
				error,
			)
		} else if lhs.Type() == ComparableTypes.RawString {
			cmp = compareStrings(lhs.(RawString), rhs.(RawString))
		} else {
			log.Fatalf("Wrong type")
		}

		entry := ComparisonEntry{
			Lhs:            lhs,
			Rhs:            rhs,
			MaxErr:         maxErr,
			HasRealNumbers: hasRealNumbers,
			CmpRes:         cmp,
		}

		// TODO: This can be simplified, I think.
		// And when I do, more things can be simplified, probably. (maybe just use one enum)
		// (usage of enum)

		entriesCh <- entry
	}

	close(entriesCh)
}
