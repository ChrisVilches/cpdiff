package cli

import (
	"bufio"
	"cpdiff/big"
	"cpdiff/cmp"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
)

const (
	separator = "\t\t"
)

func shouldSkipLine(line string, opts options) bool {
	return opts.skipEmptyLines && len(line) == 0
}

func resultColor(res cmp.ComparisonResult) color.Attribute {
	switch res {
	case cmp.CmpRes.Correct:
		return color.FgGreen
	case cmp.CmpRes.Approx:
		return color.FgYellow
	default:
		return color.FgRed
	}
}

func resultIcon(res cmp.ComparisonResult) string {
	switch res {
	case cmp.CmpRes.Correct:
		return ""
	case cmp.CmpRes.Approx:
		return "≈ "
	default:
		return "X "
	}
}

// A rather complex logic that picks the coloring function for each case.
// e.g. Number sequences have ranges related to each item, not each
// character in the raw string. so they are handled differently from strings.
func getColorFn(
	entry cmp.ComparisonEntry,
	short bool,
) func(s string, entry cmp.ComparisonEntry) (string, error) {
	if short || entry.LHS.Type() != entry.RHS.Type() {
		return colorAll
	}

	if entry.LHS.Type() == cmp.ComparableTypes.NumArray {
		return colorFields
	} else if entry.LHS.Type() == cmp.ComparableTypes.RawString {
		return colorSubstrings
	}

	return colorAll
}

func showComparisonLine(
	entry cmp.ComparisonEntry,
	opts options, line int,
) (bool, error) {
	if opts.showOnlyWrong && entry.CmpRes != cmp.CmpRes.Incorrect {
		return false, nil
	}

	pre := ""

	if opts.showLineNum {
		pre = `#{}`
		pre = fmt.Sprintf("%d\t", line)
	}

	var lhsText, rhsText string

	if opts.short {
		lhsText = entry.LHS.ShortDisplay()
		rhsText = entry.RHS.ShortDisplay()
	} else {
		lhsText = entry.LHS.Display()
		rhsText = entry.RHS.Display()
	}

	iconStr := resultIcon(entry.CmpRes)

	if opts.showColor {
		applyColor := getColorFn(entry, opts.short)
		var err error

		lhsText, err = applyColor(lhsText, entry)

		if err != nil {
			return false, err
		}

		rhsText, err = applyColor(rhsText, entry)

		if err != nil {
			return false, err
		}

		if iconStr != "" {
			iconStr = color.New(resultColor(entry.CmpRes)).Sprint(iconStr)
		}
	}

	fmt.Printf("%s%s%s%s%s\n", pre, lhsText, separator, iconStr, rhsText)

	return true, nil
}

func showFullResult(
	fullResult fullResult,
	aborted bool,
	startTime, endTime time.Time,
	opts options,
) {
	var duration string

	if opts.showDuration {
		duration = fmt.Sprintf(" (%s)", endTime.Sub(startTime))
	}

	if fullResult.totalLines == fullResult.correct {
		printfLnColor(
			color.FgGreen,
			opts.showColor,
			"Accepted %d/%d%s",
			fullResult.correct,
			fullResult.totalLines,
			duration,
		)
	} else {
		printfLnColor(
			color.FgRed,
			opts.showColor,
			"Wrong Answer %d/%d%s",
			fullResult.correct,
			fullResult.totalLines,
			duration,
		)
	}

	if aborted {
		printfLnColor(color.FgRed, opts.showColor, "Aborted")
	}

	if fullResult.hasRealNumbers {
		errType := "absolute"

		if opts.useRelativeError {
			errType = "relative"
		}

		printfLnColor(
			color.FgYellow,
			opts.showColor,
			"Max error found was %s (using %s error of %s)",
			fullResult.maxErr.String(),
			errType,
			opts.error.String(),
		)
	}
}

func readLinesToChannel(buf *bufio.Scanner, ch chan string, opts options) {
	for buf.Scan() {
		line := buf.Text()

		if opts.trim {
			line = strings.TrimSpace(line)
		}

		if shouldSkipLine(line, opts) {
			continue
		}

		ch <- line
	}

	close(ch)
}

type fullResult struct {
	totalLines     int
	correct        int
	approx         int
	hasRealNumbers bool
	maxErr         big.Decimal
}

func (v fullResult) putEntry(entry cmp.ComparisonEntry) fullResult {
	correct := v.correct
	approx := v.approx

	switch entry.CmpRes {
	case cmp.CmpRes.Approx:
		approx++
		correct++
	case cmp.CmpRes.Correct:
		correct++
	}

	return fullResult{
		totalLines:     v.totalLines + 1,
		correct:        correct,
		approx:         approx,
		hasRealNumbers: v.hasRealNumbers || entry.HasRealNumbers,
		maxErr:         big.Max(v.maxErr, entry.MaxErr),
	}
}

func newFullResult() fullResult {
	return fullResult{
		totalLines:     0,
		correct:        0,
		approx:         0,
		hasRealNumbers: false,
		maxErr:         big.Zero(),
	}
}

func getBothFiles(args []string) ([]*os.File, error) {
	files := make([]*os.File, 2)
	err := make([]error, 2)

	if len(args) == 2 {
		files[0], err[0] = os.Open(args[0])
		files[1], err[1] = os.Open(args[1])
	} else if len(args) == 1 {
		files[0] = os.Stdin
		files[1], err[0] = os.Open(args[0])
	} else {
		errMsg := "Specify 1 or 2 files (%d arguments were used)"
		return nil, fmt.Errorf(errMsg, len(args))
	}

	for _, e := range err {
		if e != nil {
			return nil, e
		}
	}

	return files, nil
}

func mainCommand(opts options, args []string) error {
	startTime := time.Now()

	files, err := getBothFiles(args)

	if err != nil {
		return err
	}

	defer files[0].Close()
	defer files[1].Close()

	lhs := bufio.NewScanner(files[0])
	rhs := bufio.NewScanner(files[1])

	const chSize = 100

	entries := make(chan cmp.ComparisonEntry, chSize)
	lhsCh := make(chan string, chSize)
	rhsCh := make(chan string, chSize)

	aborted := false

	go readLinesToChannel(lhs, lhsCh, opts)
	go readLinesToChannel(rhs, rhsCh, opts)
	go cmp.Process(
		lhsCh,
		rhsCh,
		opts.error,
		opts.useRelativeError,
		opts.numbers,
		entries,
	)

	fullResult := newFullResult()
	printedLines := false

	for entry := range entries {
		fullResult = fullResult.putEntry(entry)

		shown, err := showComparisonLine(entry, opts, fullResult.totalLines)

		if err != nil {
			return err
		}

		if shown {
			printedLines = true
		}

		if opts.abortEarly && entry.CmpRes == cmp.CmpRes.Incorrect {
			aborted = true
			break
		}
	}

	if err := lhs.Err(); err != nil {
		return err
	}

	if err := rhs.Err(); err != nil {
		return err
	}

	if printedLines {
		fmt.Println()
	}

	showFullResult(fullResult, aborted, startTime, time.Now(), opts)

	return nil
}
