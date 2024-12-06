package integration

import (
	"testing"

	"github.com/fatih/color"
)

func TestAccepted(t *testing.T) {
	lines := getLines(1)
	res := "\033[32;1mAccepted\033[0;22m\033[32;1;22m 3/3\033[0;22;0m"
	expectEq(t, lines[len(lines)-1], res)
}

func TestAcceptedNewlineOrEOF(t *testing.T) {
	lines1, statusCode1 := runCmd("-", "./data/8/ans", "1 2 3")
	lines2, statusCode2 := runCmd("-", "./data/8/ans", "1 2 3\n")
	expectEq(t, statusCode1, 0)
	expectLinesContain(t, lines1, "Accepted")
	expectEq(t, statusCode2, 0)
	expectLinesContain(t, lines2, "Accepted")
	expectEq(t, len(lines1), 3)
	expectEq(t, len(lines2), 3)
}

func TestWrongAnswerNewlinesRemoved(t *testing.T) {
	lines1, statusCode1 := runCmd("-", "./data/8/ans", "1 2 3", "-t=0")
	lines2, statusCode2 := runCmd("-", "./data/8/ans", "1 2 3\n", "-t=0")
	expectEq(t, statusCode1, 0)
	expectLinesContain(t, lines1, "Accepted")
	expectEq(t, statusCode2, 0)
	expectLinesContain(t, lines2, "Accepted")
	expectEq(t, len(lines1), 3)
	expectEq(t, len(lines2), 3)
}

func TestWrongAnswer(t *testing.T) {
	lines := getLines(2)
	res := "\033[31;1mWrong Answer\033[0;22m\033[31;1;22m 0/2\033[0;22;0m"

	expectEq(t, lines[len(lines)-1], res)
}

func TestEmptyWrongAnswer(t *testing.T) {
	lines := getLines(7)
	res := "\033[31;1mWrong Answer\033[0;22m\033[31;1;22m 1/3\033[0;22;0m"

	expectEq(t, lines[len(lines)-1], res)
}

func TestApprox(t *testing.T) {
	lines := getLines(3)
	res := "\033[32;1mAccepted\033[0;22m\033[32;1;22m 1/1\033[0;22;0m"
	expectEq(t, lines[len(lines)-2], res)
	expectEq(t, lines[len(lines)-1], color.YellowString("Max error found was 0.00001 (using absolute error of 0.0001)"))
}
