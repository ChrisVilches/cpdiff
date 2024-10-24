package cli

import (
	"github.com/ChrisVilches/cpdiff/cmp"
	"github.com/ChrisVilches/cpdiff/util"
	"slices"
	"strings"

	"github.com/fatih/color"
)

var green = color.New(color.FgGreen)
var red = color.New(color.FgRed)
var yellow = color.New(color.FgYellow)

func resultColor(res cmp.Verdict) *color.Color {
	switch res {
	case cmp.Verdicts.Correct:
		return green
	case cmp.Verdicts.Approx:
		return yellow
	default:
		return red
	}
}

func colorSubstrings(s string, entry cmp.ComparisonEntry) (string, error) {
	res := strings.Builder{}

	for _, elem := range entry.VerdictRanges {
		from := elem.From
		to := min(elem.To, len(s))

		if from >= to {
			break
		}

		c := resultColor(elem.Value)
		_, err := res.WriteString(c.Sprint(s[from:to]))

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

		c := resultColor(elem.Value)
		_, err := res.WriteString(c.Sprint(s[from:to]))

		if err != nil {
			return "", err
		}

		prev = to
	}

	return res.String(), nil
}

func colorAll(s string, entry cmp.ComparisonEntry) (string, error) {
	return resultColor(entry.Verdict).Sprint(s), nil
}
