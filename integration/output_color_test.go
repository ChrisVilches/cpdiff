package integration

import (
	"github.com/fatih/color"
	"testing"
)

func TestAcceptedColor(t *testing.T) {
	lines := getLines(1, noPaddingFlag)
	expectEq(t, lines[0], color.GreenString("AAABBBCCC")+"\t\t   "+color.GreenString("AAABBBCCC"))
}

func TestWrongAnswerStringColor(t *testing.T) {
	lines := getLines(2, noPaddingFlag)
	line := lines[0]

	expected := color.GreenString("YYYYYN") + "\t\t"
	expected += color.RedString("X") + "  " + color.GreenString("YYYYYN") + color.RedString("NN")
	expectEq(t, line, expected)
}

func TestWrongAnswerStringColor2(t *testing.T) {
	lines := getLines(2, noPaddingFlag)
	line := lines[1]

	expected := color.GreenString("YYYYYNNN") + color.RedString("NNN") + "\t\t"
	expected += color.RedString("X") + "  " + color.GreenString("YYYYYNNN")
	expectEq(t, line, expected)
}

func TestApproxColor(t *testing.T) {
	lines := getLines(3, noPaddingFlag)
	line := lines[0]

	expected := color.GreenString("1 3 ") + color.YellowString("2.00002 ") + color.GreenString("3 4")
	expected += "\t\t"
	expected += color.YellowString("≈") + "  " + color.GreenString("1 3 ") + color.YellowString("2.00001   ") + color.GreenString("3   4")
	expectEq(t, line, expected)
}

func TestWrongAnswerNumsColor(t *testing.T) {
	lines := getLines(4, noPaddingFlag)
	line := lines[0]

	expected := color.GreenString("1 2 3 4 5 6 7 ") + color.RedString("8 9")
	expected += "\t\t"
	expected += color.RedString("X") + "  " + color.GreenString("1 2  3 4  5 6  7")
	expectEq(t, line, expected)
}

func TestWrongAnswerNumsColor2(t *testing.T) {
	lines := getLines(4, noPaddingFlag)
	line := lines[1]

	expected := color.GreenString("1 2 3 4 5 6 7 8 9")
	expected += "\t\t"
	expected += color.RedString("X") + "  " + color.GreenString("1 2  3 4  5 6  7 8  9 ") + color.RedString("10  11 12")
	expectEq(t, line, expected)
}

func TestWrongAnswerNumsColor3(t *testing.T) {
	lines := getLines(4, noPaddingFlag)
	line := lines[2]

	expected := color.GreenString("1 ") + color.RedString("2 ") + color.GreenString("3 ") + color.RedString("4 ")
	expected += color.GreenString("5 ") + color.RedString("6 ") + color.GreenString("7 ") + color.RedString("8 ")
	expected += color.GreenString("9")
	expected += "\t\t"
	expected += color.RedString("X")
	expected += "  "
	expected += color.GreenString("1 ") + color.RedString("3 ") + color.GreenString("3 ") + color.RedString("5 ")
	expected += color.GreenString("5 ") + color.RedString("7 ") + color.GreenString("7 ") + color.RedString("9 ")
	expected += color.GreenString("9")

	expectEq(t, line, expected)
}

func TestStringFallbackHeuristic(t *testing.T) {
	lines := getLines(6, noPaddingFlag)
	line := lines[0]

	expected := color.RedString("0") + color.GreenString("1110001") + color.RedString("10") + color.GreenString("01010")
	expected += "\t\t"
	expected += color.RedString("X")
	expected += "  "
	expected += color.RedString("1") + color.GreenString("1110001") + color.RedString("01") + color.GreenString("01010")

	expectEq(t, line, expected)
}
