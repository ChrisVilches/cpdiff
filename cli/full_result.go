package cli

import (
	"fmt"
	"github.com/ChrisVilches/cpdiff/big"
	"github.com/ChrisVilches/cpdiff/cmp"
	"time"

	"github.com/fatih/color"
)

type fullResult struct {
	totalLines     int
	correct        int
	hasRealNumbers bool
	aborted        bool
	printedLines   int
	maxErr         big.Decimal
}

func (v fullResult) toError() error {
	if v.correct == v.totalLines {
		return nil
	}

	return NotAcceptedError{}
}

func (v fullResult) putEntry(entry cmp.ComparisonEntry) fullResult {
	if entry.Verdict != cmp.Verdicts.Incorrect {
		v.correct++
	}

	return fullResult{
		totalLines:     v.totalLines + 1,
		correct:        v.correct,
		hasRealNumbers: v.hasRealNumbers || entry.HasRealNumbers,
		maxErr:         big.Max(v.maxErr, entry.MaxErr),
	}
}

func newFullResult() fullResult {
	return fullResult{
		totalLines:     0,
		correct:        0,
		hasRealNumbers: false,
		maxErr:         big.NewZero(),
	}
}

func (v fullResult) print(
	startTime, endTime time.Time,
	opts options,
) {
	var duration string

	if opts.showDuration {
		duration = fmt.Sprintf(" (%s)", endTime.Sub(startTime))
	}

	if v.totalLines == v.correct {
		printfLnColor(
			color.FgGreen,
			opts.showColor,
			"Accepted %d/%d%s",
			v.correct,
			v.totalLines,
			duration,
		)
	} else {
		printfLnColor(
			color.FgRed,
			opts.showColor,
			"Wrong Answer %d/%d%s",
			v.correct,
			v.totalLines,
			duration,
		)
	}

	if v.aborted {
		printfLnColor(color.FgRed, opts.showColor, "Aborted")
	}

	if v.hasRealNumbers {
		errType := "absolute"

		if opts.useRelativeError {
			errType = "relative"
		}

		printfLnColor(
			color.FgYellow,
			opts.showColor,
			"Max error found was %s (using %s error of %s)",
			v.maxErr.String(),
			errType,
			opts.error.String(),
		)
	}
}
