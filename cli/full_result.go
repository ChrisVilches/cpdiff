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

func fmtVerdict(correct, total int, duration string) string {
	var c *color.Color
	var msg string
	if correct == total {
		msg = "Accepted"
		c = color.New(color.FgGreen)
	} else {
		msg = "Wrong Answer"
		c = color.New(color.FgRed)
	}

	res := ""
	c.Add(color.Bold)
	res += c.Sprint(msg)
	c.Add(color.ResetBold)
	res += c.Sprintf(" %d/%d%s", correct, total, duration)
	return res
}

func (v fullResult) print(
	startTime, endTime time.Time,
	opts options,
) {
	var duration string

	if opts.showDuration {
		duration = fmt.Sprintf(" (%s)", endTime.Sub(startTime))
	}

	fmt.Println(fmtVerdict(v.correct, v.totalLines, duration))

	if v.aborted {
		fmt.Println(color.RedString("Aborted"))
	}

	if v.hasRealNumbers {
		errType := "absolute"

		if opts.useRelativeError {
			errType = "relative"
		}

		fmt.Println(
			color.YellowString("Max error found was %s (using %s error of %s)",
				v.maxErr.String(),
				errType,
				opts.error.String(),
			),
		)
	}
}
