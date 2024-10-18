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

func getConfigError(errorString string) (res *big.Float) {
	res = new(big.Float)

	res.SetString(defaultError)

	if util.IsEmptyLine(errorString) {
		return
	}

	parsedVal := new(big.Float)

	if _, ok := parsedVal.SetString(errorString); !ok || util.BigFloatOutsideRange(parsedVal, 0, 1) {
		warn(fmt.Sprintf("Allowed error value is incorrect. Using default value %s\n", defaultError))
		return
	}

	res.Set(parsedVal)
	return
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

func showComparisonLine(entry comparison.ComparisonEntry, showLineNum, useColor, short, wrongOnly bool, line int) {
	if wrongOnly && entry.CmpRes != comparison.ComparisonResults.Incorrect {
		return
	}

	pre := ""

	if showLineNum {
		pre = `#{}`
		pre = fmt.Sprintf("%d\t", line)
	}

	var lhsText, rhsText string

	if short {
		lhsText = entry.Lhs.ShortDisplay()
		rhsText = entry.Rhs.ShortDisplay()
	} else {
		lhsText = entry.Lhs.Display()
		rhsText = entry.Rhs.Display()
	}

	colorAttribute, iconStr := iconColor(entry)

	printfColor(colorAttribute, useColor, "%s%s%s%s%s\n", pre, lhsText, separator, iconStr, rhsText)
}

func showVerdict(verdict Verdict, useColor, showDuration, aborted, useRelative bool, startTime, endTime time.Time, error *big.Float) {
	var duration string

	if showDuration {
		duration = fmt.Sprintf(" (%s)", endTime.Sub(startTime))
	}

	if verdict.totalLines == verdict.correct {
		printfColor(color.FgGreen, useColor, "Accepted %d/%d%s\n", verdict.correct, verdict.totalLines, duration)
	} else {
		printfColor(color.FgRed, useColor, "Wrong Answer %d/%d%s\n", verdict.correct, verdict.totalLines, duration)
	}

	if aborted {
		printfColor(color.FgRed, useColor, "Aborted\n")
	}

	if verdict.hasRealNumbers {
		errType := "absolute"

		if useRelative {
			errType = "relative"
		}

		printfColor(color.FgYellow, useColor, "Max error found was %s (using %s error of %s)\n", verdict.maxErr.String(), errType, error.String())
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

type Verdict struct {
	totalLines     int
	correct        int
	approx         int
	hasRealNumbers bool
	maxErr         *big.Float
}

func (v Verdict) putEntry(entry comparison.ComparisonEntry) Verdict {
	correct := v.correct
	approx := v.approx

	switch entry.CmpRes {
	case comparison.ComparisonResults.Approx:
		approx++
		correct++
	case comparison.ComparisonResults.Correct:
		correct++
	}

	return Verdict{
		totalLines:     v.totalLines + 1,
		correct:        correct,
		approx:         approx,
		hasRealNumbers: v.hasRealNumbers || entry.HasRealNumbers,
		maxErr:         util.BigMax(v.maxErr, entry.MaxErr),
	}
}

func NewVerdict() Verdict {
	return Verdict{
		totalLines:     0,
		correct:        0,
		approx:         0,
		hasRealNumbers: false,
		maxErr:         big.NewFloat(0),
	}
}

func mainCommand(short, useColor, showDuration, showLineNum, useRelative, abortEarly, wrongOnly bool, errorString string, args []string) error {
	startTime := time.Now()
	error := getConfigError(errorString)

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

	// TODO: should these channels have size???? This is very important to get right.
	entries := make(chan comparison.ComparisonEntry)
	lhsCh := make(chan string)
	rhsCh := make(chan string)

	// TODO: This ain't necessary for now, so remove it when I'm sure it's not necessary.
	// but i still need the boolean value to print the "aborted" message
	aborted := false

	go readLinesToChannel(input, lhsCh)
	go readLinesToChannel(target, rhsCh)
	go comparison.Process(lhsCh, rhsCh, error, useRelative, entries)

	verdict := NewVerdict()

	for entry := range entries {
		verdict = verdict.putEntry(entry)

		showComparisonLine(entry, showLineNum, useColor, short, wrongOnly, verdict.totalLines)

		//        ↓↓↓↓this one is fixed I think. Verify once moar↓↓↓
		// TODO: Bug... run this command `cpdiff -a datalandscape`
		// and press CTRL+D after a second.
		// It will print 0/3. This number is wrong.
		// (refactor done) One solution is to collect the stats (final result) here instead of
		// inside the Process.
		if abortEarly && entry.CmpRes == comparison.ComparisonResults.Incorrect {
			// aborted.Store(true)
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

	// TODO: Untested. It shouldn't print a new line if there were no test cases.
	if verdict.totalLines > 0 {
		fmt.Println()
	}

	showVerdict(verdict, useColor, showDuration, aborted, useRelative, startTime, time.Now(), error)

	return nil
}
