package cmp

import "github.com/ChrisVilches/cpdiff/big"

type Comparable interface {
	Display() string
	ShortDisplay(int) string
	compare(Comparable, bool, big.Decimal) ([]verdictRange, big.Decimal)
}

func SameType(a, b Comparable) bool {
	switch a.(type) {
	case NumArray:
		_, ok := b.(NumArray)
		return ok
	case RawString:
		_, ok := b.(RawString)
		return ok
	case Empty:
		_, ok := b.(Empty)
		return ok
	default:
		panic("Not handled")
	}
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
	Value    Verdict
}

func appendVerdictRange(
	list []verdictRange,
	item verdictRange,
) []verdictRange {
	if len(list) == 0 {
		return append(list, item)
	}

	last := &list[len(list)-1]

	if last.To == item.From && last.Value == item.Value {
		last.To = item.To

		return list
	}

	return append(list, item)
}
