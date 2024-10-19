package cmp

type RawString struct {
	value string
}

const (
	ellipses  = "..."
	maxLength = 10
)

func (r RawString) Display() string {
	return r.value
}

func (r RawString) ShortDisplay() string {
	if len(r.value) > maxLength {
		return r.value[0:maxLength-len(ellipses)] + ellipses
	}

	return r.value
}

func (RawString) Type() ComparableType {
	return ComparableTypes.RawString
}

func compareStrings(rs1, rs2 RawString) []cmpRange {
	n := len(rs1.value)
	m := len(rs2.value)
	size := min(n, m)
	res := []cmpRange{}

	for i := 0; i < size; i++ {
		var v ComparisonResult
		if rs1.value[i] == rs2.value[i] {
			v = CmpRes.Correct
		} else {
			v = CmpRes.Incorrect
		}

		res = appendCmpRange(res, cmpRange{From: i, To: i + 1, Result: v})
	}

	if n != m {
		res = appendCmpRange(
			res,
			cmpRange{From: size, To: max(n, m), Result: CmpRes.Incorrect},
		)
	}

	return res
}
