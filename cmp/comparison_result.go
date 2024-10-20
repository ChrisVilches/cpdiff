package cmp

// TODO: Naming is inconsistent (type and enum)
type ComparisonResult int

var CmpRes = struct {
	Correct   ComparisonResult
	Incorrect ComparisonResult
	Approx    ComparisonResult
}{
	Correct:   0,
	Incorrect: 1,
	Approx:    2,
}

type cmpRange struct {
	From, To int
	Result   ComparisonResult
}

func appendCmpRange(list []cmpRange, newValue cmpRange) []cmpRange {
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
