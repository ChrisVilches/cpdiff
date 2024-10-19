package cli

import (
	"bufio"
	"cpdiff/comparison"
	"cpdiff/util"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
)

const separator = "\t\t"
const defaultError = "0.0001"

func shouldSkipLine(line string, opts options) bool {
	return opts.removeWhitespace && len(line) == 0
}

func resultColor(res comparison.ComparisonResult) color.Attribute {
	switch res {
	case comparison.CmpRes.Correct:
		return color.FgGreen
	case comparison.CmpRes.Approx:
		return color.FgYellow
	default:
		return color.FgRed
	}
}

func resultIcon(res comparison.ComparisonResult) string {
	switch res {
	case comparison.CmpRes.Correct:
		return ""
	case comparison.CmpRes.Approx:
		return "≈ "
	default:
		return "X "
	}
}

// A rather complex logic that picks the coloring function for each case.
// e.g. Number sequences have ranges related to each item, not each character in the raw string. so they are handled differently from strings.
func getColorFn(entry comparison.ComparisonEntry, short bool) func(s string, entry comparison.ComparisonEntry) string {
	if short || entry.Lhs.Type() != entry.Rhs.Type() {
		return colorAll
	}

	if entry.Lhs.Type() == comparison.ComparableTypes.NumArray {
		return colorFields
	} else if entry.Lhs.Type() == comparison.ComparableTypes.RawString {
		return colorSubstrings
	}

	return colorAll
}

func showComparisonLine(entry comparison.ComparisonEntry, opts options, line int) bool {
	if opts.showOnlyWrong && entry.CmpRes != comparison.CmpRes.Incorrect {
		return false
	}

	pre := ""

	if opts.showLineNum {
		pre = `#{}`
		pre = fmt.Sprintf("%d\t", line)
	}

	var lhsText, rhsText string

	if opts.short {
		lhsText = entry.Lhs.ShortDisplay()
		rhsText = entry.Rhs.ShortDisplay()
	} else {
		lhsText = entry.Lhs.Display()
		rhsText = entry.Rhs.Display()
	}

	iconStr := resultIcon(entry.CmpRes)

	if opts.showColor {
		applyColor := getColorFn(entry, opts.short)

		lhsText = applyColor(lhsText, entry)
		rhsText = applyColor(rhsText, entry)
		iconStr = color.New(resultColor(entry.CmpRes)).Sprint(iconStr)
	}

	fmt.Printf("%s%s%s%s%s\n", pre, lhsText, separator, iconStr, rhsText)
	return true
}

func showFullResult(fullResult fullResult, aborted bool, startTime, endTime time.Time, opts options) {
	var duration string

	if opts.showDuration {
		duration = fmt.Sprintf(" (%s)", endTime.Sub(startTime))
	}

	if fullResult.totalLines == fullResult.correct {
		printfColor(color.FgGreen, opts.showColor, "Accepted %d/%d%s\n", fullResult.correct, fullResult.totalLines, duration)
	} else {
		printfColor(color.FgRed, opts.showColor, "Wrong Answer %d/%d%s\n", fullResult.correct, fullResult.totalLines, duration)
	}

	if aborted {
		printfColor(color.FgRed, opts.showColor, "Aborted\n")
	}

	if fullResult.hasRealNumbers {
		errType := "absolute"

		if opts.useRelativeError {
			errType = "relative"
		}

		printfColor(
			color.FgYellow,
			opts.showColor,
			"Max error found was %s (using %s error of %s)\n",
			fullResult.maxErr.String(),
			errType,
			opts.error.String(),
		)
	}
}

func readLinesToChannel(buf *bufio.Scanner, ch chan string, opts options) {
	for buf.Scan() {
		line := buf.Text()

		if opts.removeWhitespace {
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
	maxErr         *big.Float
}

func (v fullResult) putEntry(entry comparison.ComparisonEntry) fullResult {
	correct := v.correct
	approx := v.approx

	switch entry.CmpRes {
	case comparison.CmpRes.Approx:
		approx++
		correct++
	case comparison.CmpRes.Correct:
		correct++
	}

	return fullResult{
		totalLines:     v.totalLines + 1,
		correct:        correct,
		approx:         approx,
		hasRealNumbers: v.hasRealNumbers || entry.HasRealNumbers,
		maxErr:         util.BigMax(v.maxErr, entry.MaxErr),
	}
}

func NewFullResult() fullResult {
	return fullResult{
		totalLines:     0,
		correct:        0,
		approx:         0,
		hasRealNumbers: false,
		maxErr:         big.NewFloat(0),
	}
}

func mainCommand(opts options, args []string) error {
	startTime := time.Now()

	files := make([]*os.File, 2)

	if len(args) == 2 {
		files[0] = openFileOrExit(args[0])
		files[1] = openFileOrExit(args[1])
	} else if len(args) == 1 {
		files[0] = os.Stdin
		files[1] = openFileOrExit(args[0])
	} else {
		log.Fatal("Should have 1 or 2 arguments")
	}

	input := bufio.NewScanner(files[0])
	target := bufio.NewScanner(files[1])

	for _, f := range files {
		defer f.Close()
	}

	const chSize = 100

	entries := make(chan comparison.ComparisonEntry, chSize)
	lhsCh := make(chan string, chSize)
	rhsCh := make(chan string, chSize)

	aborted := false

	go readLinesToChannel(input, lhsCh, opts)
	go readLinesToChannel(target, rhsCh, opts)
	go comparison.Process(lhsCh, rhsCh, opts.error, opts.useRelativeError, entries)

	fullResult := NewFullResult()
	printedLines := false

	for entry := range entries {
		fullResult = fullResult.putEntry(entry)

		if showComparisonLine(entry, opts, fullResult.totalLines) {
			printedLines = true
		}

		if opts.abortEarly && entry.CmpRes == comparison.CmpRes.Incorrect {
			aborted = true
			break
		}
	}

	if err := input.Err(); err != nil {
		return err
	}

	if err := target.Err(); err != nil {
		return err
	}

	if printedLines {
		fmt.Println()
	}

	showFullResult(fullResult, aborted, startTime, time.Now(), opts)

	return nil
}
