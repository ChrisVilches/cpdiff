package cli

import (
	"cpdiff/cmp"
	"cpdiff/util"
	"github.com/fatih/color"
)

func colorSubstrings(s string, entry cmp.ComparisonEntry) string {
	res := ""

	for _, elem := range entry.CmpRanges {
		from := elem.From
		to := min(elem.To, len(s))
		if from > to {
			break
		}

		c := resultColor(elem.Result)
		res += color.New(c).Sprint(s[from:to])
	}

	return res
}

func colorFields(s string, entry cmp.ComparisonEntry) string {
	res := ""
	prev := 0
	j := 0

	for idx, i := range util.StringFieldsKeepWhitespace(s) {
		for entry.CmpRanges[j].To <= idx {
			j++
		}

		c := resultColor(entry.CmpRanges[j].Result)

		res += color.New(c).Sprint(s[prev:i])

		prev = i
	}

	return res
}

func colorAll(s string, entry cmp.ComparisonEntry) string {
	return color.New(resultColor(entry.CmpRes)).Sprint(s)
}
