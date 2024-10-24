package cmp

import "github.com/ChrisVilches/cpdiff/big"

type RawString struct {
	value string
}

const (
	ellipses = "..."
)

func (r RawString) Display() string {
	return r.value
}

func (r RawString) ShortDisplay(maxLength int) string {
	if len(r.value) > maxLength {
		l := max(maxLength-len(ellipses), 1)
		return r.value[0:l] + ellipses
	}

	return r.value
}

func (r RawString) compare(
	c Comparable,
	_ bool,
	_ big.Decimal,
) ([]verdictRange, big.Decimal) {
	if rhs, ok := c.(RawString); ok {
		ranges := compareStrings(r, rhs)
		return ranges, big.Decimal{}
	}

	return []verdictRange{{Value: Verdicts.Incorrect}}, big.Decimal{}
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
			verdictRange{From: i, To: i + 1, Value: v},
		)
	}

	if n != m {
		res = appendVerdictRange(
			res,
			verdictRange{From: size, To: max(n, m), Value: Verdicts.Incorrect},
		)
	}

	return res
}
