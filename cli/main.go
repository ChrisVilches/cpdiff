package cli

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"
	"sync"

	"github.com/fatih/color"
)

func isEmptyLine(line string) bool {
	return len(strings.TrimSpace(line)) == 0
}

func shouldSkipLine(line string) bool {
	return isEmptyLine(line)
}

func openFile(filePath string) *os.File {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	return file
}

func readLines(buf *bufio.Scanner, ch chan string, wg *sync.WaitGroup) {
	for buf.Scan() {
		line := buf.Text()
		if shouldSkipLine(line) {
			continue
		}
		ch <- line
	}

	close(ch)
	wg.Done()
}

// TODO: The problem is that the code using this can think of this as "int", but it should be a type.
const (
	Correct = iota
	Incorrect
	Approx
)

func toNumArray(s string) ([]big.Float, bool) {
	// TODO: Should be any amount of spaces, and the string should also have trailing/leading spaces.
	parts := strings.Split(s, " ")
	res := []big.Float{}

	for _, part := range parts {
		num := new(big.Float)

		if _, ok := num.SetString(part); !ok {
			return nil, false
		}

		// TODO: Not sure about this *num usage.
		res = append(res, *num)
	}

	return res, true
}

var Error = new(big.Float)

func compareLines(lhs, rhs string) (int, big.Float) {
	lhsAsNums, ok1 := toNumArray(lhs)
	rhsAsNums, ok2 := toNumArray(rhs)
	// TODO: perhaps doing this in every line is a bit too expensive?
	// try to globalize the biggest diff?????
	biggestDifference := new(big.Float)

	approx := false

	if ok1 && ok2 {
		if len(lhsAsNums) != len(rhsAsNums) {
			return Incorrect, *biggestDifference
		}

		for i := range len(lhsAsNums) {
			a := lhsAsNums[i]
			b := rhsAsNums[i]

			if a.Cmp(&b) == 0 {
				continue
			}

			diff := new(big.Float)
			diff.Sub(&a, &b)
			diff.Abs(diff)

			if biggestDifference.Cmp(diff) == -1 {
				biggestDifference.Set(diff)
			}

			if diff.Cmp(Error) == -1 {
				approx = true
			} else {
				return Incorrect, *biggestDifference
			}
		}

		if approx {
			return Approx, *biggestDifference
		} else {
			return Correct, *biggestDifference
		}

	} else {
		if lhs == rhs {
			return Correct, *biggestDifference
		} else {
			return Incorrect, *biggestDifference
		}
	}
}

// TODO: Try this command. Basically I want to fix this by using strings for setting the value as Big Decimal,
// never converting them to float64
// So this puzzle2 with the very small error should get the correct error.
// c++ kattis/puzzle2.cpp < datapuzzleall| CPDIFF_ERROR=0.000000001 cpdiff anspuzzleall

func bigFloatOutsideRange(bigFloat *big.Float, lo, hi float64) bool {
	bad1 := bigFloat.Cmp(new(big.Float).SetFloat64(lo)) == -1
	bad2 := bigFloat.Cmp(new(big.Float).SetFloat64(hi)) == 1
	return bad1 || bad2
}

func setConfigError() {
	defaultVal := "0.0001"
	Error.SetString(defaultVal)

	envVar := "CPDIFF_ERROR"
	errorString := os.Getenv(envVar)

	if isEmptyLine(errorString) {
		return
	}

	envVal := new(big.Float)

	if _, ok := envVal.SetString(errorString); !ok || bigFloatOutsideRange(envVal, 0, 1) {
		fmt.Fprintf(os.Stderr, "Warning: Value of %s is incorrect. Using default value %s\n", envVar, defaultVal)
		return
	}

	Error.Set(envVal)
}

func App() {
	setConfigError()

	var input *bufio.Scanner
	var target *bufio.Scanner

	programArgs := os.Args[1:]

	if len(programArgs) >= 2 {
		inputFile := openFile(programArgs[0])
		targetFile := openFile(programArgs[1])
		input = bufio.NewScanner(inputFile)
		target = bufio.NewScanner(targetFile)

		defer inputFile.Close()
		defer targetFile.Close()
	} else if len(programArgs) == 1 {
		input = bufio.NewScanner(os.Stdin)
		targetFile := openFile(programArgs[0])
		target = bufio.NewScanner(targetFile)
		defer targetFile.Close()
	} else {
		panic("Should have at least one argument")
	}

	lhs := make(chan string)
	rhs := make(chan string)

	var wg sync.WaitGroup

	go readLines(input, lhs, &wg)
	go readLines(target, rhs, &wg)

	wg.Add(2)

	totalLines := 0
	correct := 0
	approx := 0
	biggestDifference := new(big.Float)

	separator := "\t\t"

	for {
		l, ok1 := <-lhs
		r, ok2 := <-rhs

		lhsText := "-"
		rhsText := "-"

		if ok1 {
			lhsText = l
		}

		if ok2 {
			rhsText = r
		}

		if !ok1 && !ok2 {
			break
		}

		cmp, diff := compareLines(l, r)

		if biggestDifference.Cmp(&diff) == -1 {
			biggestDifference.Set(&diff)
		}

		if cmp == Correct {
			correct++
			fmt.Println(color.GreenString("%s%s%s", lhsText, separator, rhsText))
		} else if cmp == Approx {
			correct++
			approx++
			fmt.Println(color.YellowString("%s%s≈ %s", lhsText, separator, rhsText))
		} else {
			fmt.Println(color.RedString("%s%sX %s", lhsText, separator, rhsText))
		}

		totalLines++
	}

	// TODO: Should this be before the scan loop?
	if err := input.Err(); err != nil {
		log.Println(err)
	}

	if err := target.Err(); err != nil {
		log.Println(err)
	}

	wg.Wait()

	fmt.Println()

	if totalLines == correct {
		fmt.Println(color.GreenString("Accepted %d/%d", totalLines, totalLines))
	} else {
		fmt.Println(color.RedString("Wrong Answer %d/%d", correct, totalLines))
	}

	if approx > 0 {
		fmt.Println(color.YellowString("Biggest difference was %s (allowing %s)", biggestDifference.String(), Error.String()))
	}
}

// TODO: Add flag to number lines
// TODO: Add flag to skip correct lines.
// TODO: Flag to stop when both channels have different length
// TODO: Add flag to show all lines including empty ones (this is hard because I need to format properly and know which ones to compare to which ones)
