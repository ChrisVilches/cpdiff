package cli

import (
	"cpdiff/cmp"
	"cpdiff/util"
	"slices"
	"strings"

	"github.com/fatih/color"
)

func colorSubstrings(s string, entry cmp.ComparisonEntry) (string, error) {
	res := strings.Builder{}

	for _, elem := range entry.VerdictRanges {
		from := elem.From
		to := min(elem.To, len(s))

		if from >= to {
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
	pos := slices.Collect(util.StringFieldsKeepWhitespace(s))
	res := strings.Builder{}
	prev := 0

	for _, elem := range entry.VerdictRanges {
		lastPos := pos[min(elem.To, len(pos))-1]
		from := prev
		to := min(lastPos, len(s))

		if from >= to {
			break
		}

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
	return color.New(resultColor(entry.Verdict)).Sprint(s), nil
}
