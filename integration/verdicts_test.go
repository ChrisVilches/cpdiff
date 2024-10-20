package integration

import (
	"github.com/fatih/color"
	"testing"
)

func TestAccepted(t *testing.T) {
	lines := getLines(1)

	test(t, lines[len(lines)-2], color.GreenString("Accepted 3/3"))
	test(t, lines[len(lines)-1], color.YellowString("Max error found was 0 (using absolute error of 0.0001)"))
}

func TestWrongAnswer(t *testing.T) {
	lines := getLines(2)

	test(t, lines[len(lines)-1], color.RedString("Wrong Answer 0/2"))
}

func TestApprox(t *testing.T) {
	lines := getLines(3)

	test(t, lines[len(lines)-2], color.GreenString("Accepted 1/1"))
	test(t, lines[len(lines)-1], color.YellowString("Max error found was 0.00001 (using absolute error of 0.0001)"))
}
