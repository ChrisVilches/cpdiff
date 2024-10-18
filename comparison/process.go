package comparison

import (
	"log"
	"math/big"
)

type ProcessResult struct {
	Correct           int
	Approx            int
	TotalLines        int
	HasRealNumbers    bool
	BiggestDifference big.Float
}

type LineCmpResult int

var LineCmpResults = struct {
	Correct   LineCmpResult
	Incorrect LineCmpResult
	Approx    LineCmpResult
}{
	Correct:   0,
	Incorrect: 1,
	Approx:    2,
}

type ComparisonEntry struct {
	Lhs     Comparable
	Rhs     Comparable
	Verdict LineCmpResult
}

func handleNumArrays(lhs, rhs *NumArray, relativeErr bool, biggestDifference, error *big.Float, hasRealNumbers *bool) ComparisonResult {
	var diff big.Float
	var cmp ComparisonResult
	cmp, diff = compareNums(lhs, rhs, error, relativeErr)

	if biggestDifference.Cmp(&diff) == -1 {
		biggestDifference.Set(&diff)
	}

	*hasRealNumbers = *hasRealNumbers || lhs.HasRealNumbers() || lhs.HasRealNumbers()
	return cmp
}

func Process(
	lhsCh,
	rhsCh chan string,
	error big.Float,
	relativeErr bool,
	comparisonResults chan ComparisonEntry,
	ret chan ProcessResult,
) {

	result := ProcessResult{}

	for {
		l, ok1 := <-lhsCh
		r, ok2 := <-rhsCh

		if !ok1 && !ok2 {
			break
		}

		lhs := LineToComparable(l)
		rhs := LineToComparable(r)

		var cmp ComparisonResult

		if lhs.Type() != rhs.Type() {
			cmp = ComparisonResults.Incorrect
		} else if lhs.Type() == ComparableTypes.NumArray {
			cmp = handleNumArrays(
				lhs.(*NumArray),
				rhs.(*NumArray),
				relativeErr,
				&result.BiggestDifference,
				&error,
				&result.HasRealNumbers,
			)
		} else if lhs.Type() == ComparableTypes.RawString {
			cmp = compareStrings(lhs.(*RawString), rhs.(*RawString))
		} else {
			log.Fatalf("Wrong type")
		}

		entry := ComparisonEntry{Lhs: lhs, Rhs: rhs}

		if cmp == ComparisonResults.Correct {
			result.Correct++
			entry.Verdict = LineCmpResults.Correct
		} else if cmp == ComparisonResults.Approx {
			result.Correct++
			result.Approx++
			entry.Verdict = LineCmpResults.Approx
		} else {
			entry.Verdict = LineCmpResults.Incorrect
		}

		comparisonResults <- entry

		result.TotalLines++
	}

	ret <- result
}
