package integration

import (
	"testing"

	"github.com/fatih/color"
)

func TestAccepted(t *testing.T) {
	lines := getLines(1)
	res := "\033[32;1mAccepted\033[0;22m\033[32;1;22m 3/3\033[0;22;0m"
	expectEq(t, lines[len(lines)-2], res)
	expectEq(t, lines[len(lines)-1], color.YellowString("Max error found was 0 (using absolute error of 0.0001)"))
}

func TestWrongAnswer(t *testing.T) {
	lines := getLines(2)
	res := "\033[31;1mWrong Answer\033[0;22m\033[31;1;22m 0/2\033[0;22;0m"

	expectEq(t, lines[len(lines)-1], res)
}

func TestApprox(t *testing.T) {
	lines := getLines(3)
	res := "\033[32;1mAccepted\033[0;22m\033[32;1;22m 1/1\033[0;22;0m"
	expectEq(t, lines[len(lines)-2], res)
	expectEq(t, lines[len(lines)-1], color.YellowString("Max error found was 0.00001 (using absolute error of 0.0001)"))
}
