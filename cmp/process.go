package cmp

import (
	"log"
	"math/big"
)

type ComparisonEntry struct {
	LHS            Comparable
	RHS            Comparable
	CmpRes         ComparisonResult
	CmpRanges      []cmpRange
	HasRealNumbers bool
	MaxErr         *big.Float
}

func handleNumArrays(
	lhs,
	rhs NumArray,
	relativeErr bool,
	allowedError *big.Float,
) ([]cmpRange, *big.Float, bool) {
	cmpRanges, diff := compareNums(lhs, rhs, allowedError, relativeErr)

	return cmpRanges, diff, lhs.HasRealNumbers() || rhs.HasRealNumbers()
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

func newComparisonEntry(
	lhs, rhs Comparable,
	useRelativeErr bool,
	allowedError *big.Float,
) ComparisonEntry {
	e := ComparisonEntry{LHS: lhs, RHS: rhs}

	if lhs.Type() != rhs.Type() {
		e.CmpRes = CmpRes.Incorrect
	} else {
		switch lhs.Type() {
		case ComparableTypes.RawString:
			e.CmpRanges = compareStrings(lhs.(RawString), rhs.(RawString))
		case ComparableTypes.NumArray:
			e.CmpRanges, e.MaxErr, e.HasRealNumbers = handleNumArrays(
				lhs.(NumArray),
				rhs.(NumArray),
				useRelativeErr,
				allowedError,
			)
		case ComparableTypes.Empty:
			e.CmpRes = CmpRes.Correct
		default:
			panic("Wrong type")
		}
	}

	if len(e.CmpRanges) != 0 {
		e.CmpRes = findGlobalResult(e.CmpRanges)
	}

	return e
}

func Process(
	lhsCh,
	rhsCh chan string,
	allowedError *big.Float,
	useRelativeErr bool,
	entriesCh chan ComparisonEntry,
) {
	for {
		line1, ok1 := <-lhsCh
		line2, ok2 := <-rhsCh

		if !ok1 && !ok2 {
			break
		}

		lhs := LineToComparable(line1)
		rhs := LineToComparable(line2)

		entriesCh <- newComparisonEntry(lhs, rhs, useRelativeErr, allowedError)
	}

	close(entriesCh)
}
