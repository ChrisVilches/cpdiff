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

type SingleComparisonResult int

const (
	SingleComparisonCorrect SingleComparisonResult = iota
	SingleComparisonApproxCorrect
	SingleComparisonIncorrect
)

type ComparisonEntry struct {
	Lhs     Comparable
	Rhs     Comparable
	Verdict SingleComparisonResult
}

func handleNumArrays(lhs, rhs *NumArray, relativeErr bool, biggestDifference, error *big.Float, hasRealNumbers *bool) int {
	var diff big.Float
	var cmp int
	cmp, diff = compareNums(lhs, rhs, error, relativeErr)

	if biggestDifference.Cmp(&diff) == -1 {
		biggestDifference.Set(&diff)
	}

	*hasRealNumbers = *hasRealNumbers || lhs.HasRealNumbers() || lhs.HasRealNumbers()
	return cmp
}

func Process(
	lhsStream,
	rhsStream chan string,
	error big.Float,
	relativeErr bool,
	comparisonResults chan ComparisonEntry,
	ret chan ProcessResult,
) {

	result := ProcessResult{}

	for {
		l, ok1 := <-lhsStream
		r, ok2 := <-rhsStream

		if !ok1 && !ok2 {
			break
		}

		lhs := LineToComparable(l)
		rhs := LineToComparable(r)

		var cmp int

		if lhs.Type() != rhs.Type() {
			cmp = Incorrect
		} else if lhs.Type() == "num_array" {
			cmp = handleNumArrays(
				lhs.(*NumArray),
				rhs.(*NumArray),
				relativeErr,
				&result.BiggestDifference,
				&error,
				&result.HasRealNumbers,
			)
		} else if lhs.Type() == "string" {
			cmp = compareStrings(lhs.(*RawString), rhs.(*RawString))
		} else {
			log.Fatalf("Wrong type")
		}

		entry := ComparisonEntry{Lhs: lhs, Rhs: rhs}

		// TODO: Should allow to remove colors (a CLI flag)
		if cmp == Correct {
			result.Correct++
			entry.Verdict = SingleComparisonCorrect
		} else if cmp == Approx {
			result.Correct++
			result.Approx++
			entry.Verdict = SingleComparisonApproxCorrect
		} else {
			entry.Verdict = SingleComparisonIncorrect
		}

		comparisonResults <- entry

		result.TotalLines++
	}

	ret <- result
}
