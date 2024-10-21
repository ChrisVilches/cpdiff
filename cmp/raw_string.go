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

func compareStrings(rs1, rs2 RawString) []verdictRange {
	n := len(rs1.value)
	m := len(rs2.value)
	size := min(n, m)
	res := []verdictRange{}

	for i := 0; i < size; i++ {
		var v Verdict
		if rs1.value[i] == rs2.value[i] {
			v = Verdicts.Correct
		} else {
			v = Verdicts.Incorrect
		}

		res = appendVerdictRange(res,
			verdictRange{From: i, To: i + 1, Result: v},
		)
	}

	if n != m {
		res = appendVerdictRange(
			res,
			verdictRange{From: size, To: max(n, m), Result: Verdicts.Incorrect},
		)
	}

	return res
}
