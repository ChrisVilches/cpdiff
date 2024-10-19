package cli

import (
	"bufio"
	"cpdiff/comparison"
	"cpdiff/util"
	"fmt"
	"github.com/fatih/color"
	"log"
	"math/big"
	"os"
	"time"
)

const separator = "\t\t"
const defaultError = "0.0001"

func shouldSkipLine(line string) bool {
	return util.IsEmptyLine(line)
}

func iconColor(entry comparison.ComparisonEntry) (color.Attribute, string) {
	switch entry.CmpRes {
	case comparison.ComparisonResults.Correct:
		return color.FgGreen, ""
	case comparison.ComparisonResults.Approx:
		return color.FgYellow, "≈ "
	default:
		return color.FgRed, "X "
	}
}

func showComparisonLine(entry comparison.ComparisonEntry, opts options, line int) bool {
	if opts.showOnlyWrong && entry.CmpRes != comparison.ComparisonResults.Incorrect {
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

	colorAttribute, iconStr := iconColor(entry)

	printfColor(colorAttribute, opts.showColor, "%s%s%s%s%s\n", pre, lhsText, separator, iconStr, rhsText)
	return true
}

func showFullResult(fullResult FullResult, aborted bool, startTime, endTime time.Time, opts options) {
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

func readLinesToChannel(buf *bufio.Scanner, ch chan string) {
	for buf.Scan() {
		line := buf.Text()

		if shouldSkipLine(line) {
			continue
		}

		ch <- line
	}

	close(ch)
}

// TODO: CHange the name of this struct and related methods/parameters because it should be more like
//
//	"full result" or something like that.
type FullResult struct {
	totalLines     int
	correct        int
	approx         int
	hasRealNumbers bool
	maxErr         *big.Float
}

func (v FullResult) putEntry(entry comparison.ComparisonEntry) FullResult {
	correct := v.correct
	approx := v.approx

	switch entry.CmpRes {
	case comparison.ComparisonResults.Approx:
		approx++
		correct++
	case comparison.ComparisonResults.Correct:
		correct++
	}

	return FullResult{
		totalLines:     v.totalLines + 1,
		correct:        correct,
		approx:         approx,
		hasRealNumbers: v.hasRealNumbers || entry.HasRealNumbers,
		maxErr:         util.BigMax(v.maxErr, entry.MaxErr),
	}
}

func NewFullResult() FullResult {
	return FullResult{
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

	// TODO: Lol? it was possible to define constants inside functions.
	// I need to change many things I think.
	const chSize = 100

	// TODO: should these channels have size???? This is very important to get right.
	entries := make(chan comparison.ComparisonEntry, chSize)
	lhsCh := make(chan string, chSize)
	rhsCh := make(chan string, chSize)

	aborted := false

	go readLinesToChannel(input, lhsCh)
	go readLinesToChannel(target, rhsCh)
	go comparison.Process(lhsCh, rhsCh, opts.error, opts.useRelativeError, entries)

	fullResult := NewFullResult()
	printedLines := false

	for entry := range entries {
		fullResult = fullResult.putEntry(entry)

		if showComparisonLine(entry, opts, fullResult.totalLines) {
			printedLines = true
		}

		if opts.abortEarly && entry.CmpRes == comparison.ComparisonResults.Incorrect {
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
