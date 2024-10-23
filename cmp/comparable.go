package cmp

import "github.com/ChrisVilches/cpdiff/big"

type Comparable interface {
	Type() ComparableType
	Display() string
	ShortDisplay(int) string
}

type ComparableType int

var ComparableTypes = struct {
	RawString ComparableType
	NumArray  ComparableType
	Empty     ComparableType
}{
	RawString: 0,
	NumArray:  1,
	Empty:     2,
}

type ComparisonEntry struct {
	LHS           Comparable
	RHS           Comparable
	Verdict       Verdict
	VerdictRanges []verdictRange
	MaxErr        big.Decimal
}

func newComparisonEntry(
	lhs, rhs Comparable,
	useRelativeErr bool,
	allowedError big.Decimal,
) ComparisonEntry {
	e := ComparisonEntry{LHS: lhs, RHS: rhs, MaxErr: big.NewZero()}

	if lhs.Type() != rhs.Type() {
		e.Verdict = Verdicts.Incorrect
		return e
	}

	switch lhs.Type() {
	case ComparableTypes.RawString:
		e.VerdictRanges = compareStrings(lhs.(RawString), rhs.(RawString))
	case ComparableTypes.NumArray:
		e.VerdictRanges, e.MaxErr = compareNums(
			lhs.(NumArray),
			rhs.(NumArray),
			allowedError,
			useRelativeErr,
		)
	case ComparableTypes.Empty:
		e.Verdict = Verdicts.Correct
	default:
		panic("Wrong type")
	}

	e.Verdict = findGlobalResult(e.VerdictRanges)
	return e
}

type Verdict int

var Verdicts = struct {
	Correct   Verdict
	Incorrect Verdict
	Approx    Verdict
}{
	Correct:   0,
	Incorrect: 1,
	Approx:    2,
}

type verdictRange struct {
	From, To int
	Result   Verdict
}

func appendVerdictRange(
	list []verdictRange,
	newValue verdictRange,
) []verdictRange {
	if len(list) == 0 {
		return append(list, newValue)
	}

	last := &list[len(list)-1]

	if last.To == newValue.From && last.Result == newValue.Result {
		last.To = newValue.To

		return list
	}

	return append(list, newValue)
}
