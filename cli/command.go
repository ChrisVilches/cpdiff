package cli

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/ChrisVilches/cpdiff/cmp"
	"os"
	"strings"
	"time"
)

const (
	sep    = "\t\t"
	chSize = 100
)

func shouldSkipLine(line string, opts options) bool {
	return opts.skipEmptyLines && len(line) == 0
}

func resultIcon(res cmp.Verdict) string {
	switch res {
	case cmp.Verdicts.Correct:
		return " "
	case cmp.Verdicts.Approx:
		return "≈"
	default:
		return "X"
	}
}

// A rather complex logic that picks the coloring function for each case.
// e.g. Number sequences have ranges related to each item, not each
// character in the raw string. so they are handled differently from strings.
func getColorFn(
	entry cmp.ComparisonEntry,
	short bool,
) func(s string, entry cmp.ComparisonEntry) (string, error) {
	if short || !cmp.SameType(entry.LHS, entry.RHS) {
		return colorAll
	}

	switch entry.LHS.(type) {
	case cmp.NumArray:
		return colorFields
	case cmp.RawString:
		return colorSubstrings
	}

	return colorAll
}

func getPrefix(currLine int, opts options) string {
	if opts.showLineNum {
		return fmt.Sprintf("%d\t", currLine)
	}

	return ""
}

func getColMaxLengthHeuristic(opts options) int {
	return opts.leftExtraPadding
}

func getPadding(strlen, maxPad int) string {
	diff := maxPad - strlen
	if diff <= 0 {
		return ""
	}
	return strings.Repeat(" ", diff)
}

func showComparisonEntry(
	entry cmp.ComparisonEntry,
	opts options, line int,
) (bool, error) {
	correctAns := entry.Verdict != cmp.Verdicts.Incorrect

	if opts.quiet || (opts.showOnlyWrong && correctAns) {
		return false, nil
	}

	pre := getPrefix(line, opts)
	var lhsText, rhsText string

	if opts.short {
		maxLength := getColMaxLengthHeuristic(opts)
		lhsText = entry.LHS.ShortDisplay(maxLength)
		rhsText = entry.RHS.ShortDisplay(maxLength)
	} else {
		lhsText = entry.LHS.Display()
		rhsText = entry.RHS.Display()
	}

	iconStr := resultIcon(entry.Verdict)
	padding := getPadding(len(lhsText), opts.leftExtraPadding)

	applyColor := getColorFn(entry, opts.short)
	var err error

	if lhsText, err = applyColor(lhsText, entry); err != nil {
		return false, err
	}

	if rhsText, err = applyColor(rhsText, entry); err != nil {
		return false, err
	}

	if iconStr != " " {
		iconStr = resultColor(entry.Verdict).Sprint(iconStr)
	}

	fmt.Printf("%s%s%s%s%s  %s\n", pre, lhsText, padding, sep, iconStr, rhsText)

	return true, nil
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

func getFile(path string) (*os.File, error) {
	if path == "-" {
		return os.Stdin, nil
	}

	return os.Open(path)
}

func getBothFiles(args []string) ([]*os.File, error) {
	files := make([]*os.File, 2)
	err := make([]error, 2)

	if len(args) == 2 {
		if args[0] == "-" && args[1] == "-" {
			msg := "Do not use - (standard input) for both sides"
			return nil, errors.New(msg)
		}
		files[0], err[0] = getFile(args[0])
		files[1], err[1] = getFile(args[1])
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

func listenEntries(
	entries chan cmp.ComparisonEntry,
	opts options,
) (*fullResult, error) {
	res := newFullResult()

	for entry := range entries {
		res = res.putEntry(entry)

		shown, err := showComparisonEntry(entry, opts, res.totalLines)

		if err != nil {
			return nil, err
		}

		if shown {
			res.printedLines++
		}

		if opts.abortEarly && entry.Verdict == cmp.Verdicts.Incorrect {
			res.aborted = true
			break
		}
	}

	return &res, nil
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

	entries := make(chan cmp.ComparisonEntry, chSize)
	lhsCh := make(chan string, chSize)
	rhsCh := make(chan string, chSize)

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

	fullResult, err := listenEntries(entries, opts)

	if err != nil {
		return err
	}

	if err := lhs.Err(); err != nil {
		return err
	}

	if err := rhs.Err(); err != nil {
		return err
	}

	if fullResult.printedLines > 0 {
		fmt.Println()
	}

	fullResult.print(startTime, time.Now(), opts)

	return fullResult.toError()
}
