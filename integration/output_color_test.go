package integration

import (
	"github.com/fatih/color"
	"testing"
)

func TestAcceptedColor(t *testing.T) {
	lines := getLines(1)
	test(t, lines[0], color.GreenString("AAABBBCCC")+"\t\t"+color.GreenString("AAABBBCCC"))
}

func TestWrongAnswerStringColor(t *testing.T) {
	lines := getLines(2)
	line := lines[0]

	expected := color.GreenString("YYYYYN") + "\t\t"
	expected += color.RedString("X ") + color.GreenString("YYYYYN") + color.RedString("NN")
	test(t, line, expected)
}

func TestWrongAnswerStringColor2(t *testing.T) {
	lines := getLines(2)
	line := lines[1]

	expected := color.GreenString("YYYYYNNN") + color.RedString("NNN") + "\t\t"
	expected += color.RedString("X ") + color.GreenString("YYYYYNNN")
	test(t, line, expected)
}

func TestApproxColor(t *testing.T) {
	lines := getLines(3)
	line := lines[0]

	expected := color.GreenString("1 3 ") + color.YellowString("2.00002 ") + color.GreenString("3 4")
	expected += "\t\t"
	expected += color.YellowString("≈ ") + color.GreenString("1 3 ") + color.YellowString("2.00001   ") + color.GreenString("3   4")
	test(t, line, expected)
}

func TestWrongAnswerNumsColor(t *testing.T) {
	lines := getLines(4)
	line := lines[0]

	expected := color.GreenString("1 2 3 4 5 6 7 ") + color.RedString("8 9")
	expected += "\t\t"
	expected += color.RedString("X ") + color.GreenString("1 2  3 4  5 6  7")
	test(t, line, expected)
}

func TestWrongAnswerNumsColor2(t *testing.T) {
	lines := getLines(4)
	line := lines[1]

	expected := color.GreenString("1 2 3 4 5 6 7 8 9")
	expected += "\t\t"
	expected += color.RedString("X ") + color.GreenString("1 2  3 4  5 6  7 8  9 ") + color.RedString("10  11 12")
	test(t, line, expected)
}

func TestWrongAnswerNumsColor3(t *testing.T) {
	lines := getLines(4)
	line := lines[2]

	expected := color.GreenString("1 ") + color.RedString("2 ") + color.GreenString("3 ") + color.RedString("4 ")
	expected += color.GreenString("5 ") + color.RedString("6 ") + color.GreenString("7 ") + color.RedString("8 ")
	expected += color.GreenString("9")
	expected += "\t\t"
	expected += color.RedString("X ")
	expected += color.GreenString("1 ") + color.RedString("3 ") + color.GreenString("3 ") + color.RedString("5 ")
	expected += color.GreenString("5 ") + color.RedString("7 ") + color.GreenString("7 ") + color.RedString("9 ")
	expected += color.GreenString("9")

	test(t, line, expected)
}
