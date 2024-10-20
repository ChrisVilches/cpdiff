package cli

import (
	"cpdiff/cmp"
	"cpdiff/util"
	"strings"

	"github.com/fatih/color"
)

func colorSubstrings(s string, entry cmp.ComparisonEntry) (string, error) {
	res := strings.Builder{}

	for _, elem := range entry.CmpRanges {
		from := elem.From
		to := min(elem.To, len(s))

		if from > to {
			break
		}

		c := resultColor(elem.Result)
		_, err := res.WriteString(color.New(c).Sprint(s[from:to]))

		if err != nil {
			return "", err
		}
	}

	return res.String(), nil
}

func colorFields(s string, entry cmp.ComparisonEntry) (string, error) {
	res := strings.Builder{}
	prev := 0
	j := 0

	for idx, i := range util.StringFieldsKeepWhitespace(s) {
		for entry.CmpRanges[j].To <= idx {
			j++
		}

		c := resultColor(entry.CmpRanges[j].Result)

		_, err := res.WriteString(color.New(c).Sprint(s[prev:i]))

		if err != nil {
			return "", err
		}

		prev = i
	}

	return res.String(), nil
}

func colorAll(s string, entry cmp.ComparisonEntry) (string, error) {
	return color.New(resultColor(entry.CmpRes)).Sprint(s), nil
}
