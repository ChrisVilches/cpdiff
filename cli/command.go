package cli

import (
	"bufio"
	"cpdiff/comparison"
	"cpdiff/util"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/fatih/color"
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

func showComparisonLine(entry comparison.ComparisonEntry, showLineNum, useColor, short bool, line int) {
	pre := ""

	if showLineNum {
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

	var c color.Attribute
	icon := ""

	switch entry.Verdict {
	case comparison.SingleComparisonCorrect:
		c = color.FgGreen
	case comparison.SingleComparisonApproxCorrect:
		c = color.FgYellow
		icon = "≈ "
	default:
		c = color.FgRed
		icon = "X "
	}

	printfColor(c, useColor, "%s%s%s%s%s\n", pre, lhsText, separator, icon, rhsText)
}

func showVerdict(result comparison.ProcessResult, useColor, showDuration bool, startTime, endTime time.Time, error *big.Float) {
	var duration string

	if showDuration {
		duration = fmt.Sprintf(" (%s)", endTime.Sub(startTime))
	}

	if result.TotalLines == result.Correct {
		printfColor(color.FgGreen, useColor, "Accepted %d/%d%s\n", result.Correct, result.TotalLines, duration)
	} else {
		printfColor(color.FgRed, useColor, "Wrong Answer %d/%d%s\n", result.Correct, result.TotalLines, duration)
	}

	if result.HasRealNumbers {
		printfColor(color.FgYellow, useColor, "Biggest difference found was %s (allowing %s)\n", result.BiggestDifference.String(), error.String())
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

func mainCommand(short, useColor, showDuration, showLineNum, useRelative bool, errorString string, args []string) error {
	startTime := time.Now()
	error := getConfigError(errorString)

	var input *bufio.Scanner
	var target *bufio.Scanner
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

	input = bufio.NewScanner(files[0])
	target = bufio.NewScanner(files[1])

	for _, f := range files {
		defer f.Close()
	}

	// TODO: should these channels have size???? This is very important to get right.
	lines := make(chan comparison.ComparisonEntry)
	stats := make(chan comparison.ProcessResult)
	lhsCh := make(chan string)
	rhsCh := make(chan string)

	go readLinesToChannel(input, lhsCh)
	go readLinesToChannel(target, rhsCh)
	go comparison.Process(lhsCh, rhsCh, *error, useRelative, lines, stats)

	currLine := 1

Select:
	for {
		select {
		case elem := <-lines:
			showComparisonLine(elem, showLineNum, useColor, short, currLine)
			currLine++

		case result := <-stats:
			showVerdict(result, useColor, showDuration, startTime, time.Now(), error)
			break Select
		}
	}

	if err := input.Err(); err != nil {
		return err
	}

	if err := target.Err(); err != nil {
		return err
	}

	close(lines)
	close(stats)

	return nil
}
