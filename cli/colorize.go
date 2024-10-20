package cli

import (
	"cpdiff/cmp"
	"cpdiff/util"
	"slices"
	"strings"

	"github.com/fatih/color"
)

// TODO: test these two-pointer methods.
// TODO: I think the min() does something weird (specially to the last part of the string)
func colorSubstrings(s string, entry cmp.ComparisonEntry) (string, error) {
	res := strings.Builder{}

	for _, elem := range entry.CmpRanges {
		from := elem.From
		to := min(elem.To, len(s))

		c := resultColor(elem.Result)
		_, err := res.WriteString(color.New(c).Sprint(s[from:to]))

		if err != nil {
			return "", err
		}
	}

	return res.String(), nil
}

func colorFields(s string, entry cmp.ComparisonEntry) (string, error) {
	pos := slices.Collect(util.StringFieldsKeepWhitespace(s))
	res := strings.Builder{}
	prev := 0

	for _, elem := range entry.CmpRanges {
		lastPos := pos[min(elem.To, len(pos))-1]
		from := prev
		to := min(lastPos, len(s))

		c := resultColor(elem.Result)
		_, err := res.WriteString(color.New(c).Sprint(s[from:to]))

		if err != nil {
			return "", err
		}

		prev = to
	}

	return res.String(), nil
}

func colorAll(s string, entry cmp.ComparisonEntry) (string, error) {
	return color.New(resultColor(entry.CmpRes)).Sprint(s), nil
}
