package comparison

import (
	"log"
	"math/big"
)

type ComparisonEntry struct {
	Lhs            Comparable
	Rhs            Comparable
	CmpRes         ComparisonResult
	CmpRanges      []cmpRange
	HasRealNumbers bool
	MaxErr         *big.Float
}

func handleNumArrays(lhs, rhs NumArray, relativeErr bool, error *big.Float) ([]cmpRange, *big.Float, bool) {
	cmpRanges, diff := compareNums(lhs, rhs, error, relativeErr)

	return cmpRanges, diff, lhs.HasRealNumbers() || lhs.HasRealNumbers()
}

func findGlobalResult(cmpRanges []cmpRange) ComparisonResult {
	approx := false

	for _, r := range cmpRanges {
		if r.Result == CmpRes.Incorrect {
			return CmpRes.Incorrect
		}

		if r.Result == CmpRes.Approx {
			approx = true
		}
	}

	if approx {
		return CmpRes.Approx
	}

	return CmpRes.Correct
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

		var cmp []cmpRange
		maxErr := big.NewFloat(0)
		hasRealNumbers := false
		var globalResult ComparisonResult

		if lhs.Type() != rhs.Type() {
			globalResult = CmpRes.Incorrect
		} else if lhs.Type() == ComparableTypes.NumArray {
			cmp, maxErr, hasRealNumbers = handleNumArrays(
				lhs.(NumArray),
				rhs.(NumArray),
				relativeErr,
				error,
			)
		} else if lhs.Type() == ComparableTypes.RawString {
			cmp = compareStrings(lhs.(RawString), rhs.(RawString))
		} else if lhs.Type() == ComparableTypes.Empty {
			globalResult = CmpRes.Correct
		} else {
			log.Fatalf("Wrong type")
		}

		if len(cmp) != 0 {
			globalResult = findGlobalResult(cmp)
		}

		entry := ComparisonEntry{
			Lhs:            lhs,
			Rhs:            rhs,
			MaxErr:         maxErr,
			HasRealNumbers: hasRealNumbers,
			CmpRanges:      cmp,
			CmpRes:         globalResult,
		}

		entriesCh <- entry
	}

	close(entriesCh)
}
