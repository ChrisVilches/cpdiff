package integration

import (
	"github.com/fatih/color"
	"math/rand"
	"strings"
	"testing"
)

func TestQuietFlag(t *testing.T) {
	flags := []string{"-q", "--quiet"}

	codes := []int{0, 1, 0, 1, 0, 1}

	for i := 0; i < 6; i++ {
		expectEq(t, getStatusCode(i+1), codes[i])
		lines := getLines(i+1, flags[rand.Int()%2])
		expectEq(t, len(lines), 0)
	}
}

func TestExitFlag(t *testing.T) {
	expectLinesContain(t, getLines(2, "-x"), "Aborted")
	expectLinesContain(t, getLines(2, "--exit"), "Aborted")
	expectLinesNotContain(t, getLines(2), "Aborted")
}

func TestDurationFlag(t *testing.T) {
	expectLinesContain(t, getLines(2, "-d"), "0/2 (")
	expectLinesContain(t, getLines(2, "--duration"), "s)")
	expectLinesNotContain(t, getLines(2), "(")
}

func TestWrongFlag(t *testing.T) {
	expectEq(t, len(getLines(1, "-w")), 1)
	expectEq(t, len(getLines(1, "--wrong")), 1)
	expectEq(t, len(getLines(1)), 5)
	expectEq(t, len(getLines(5, "-w")), 1)
	expectEq(t, len(getLines(5, "--wrong")), 1)
	expectEq(t, len(getLines(5)), 4)
}

func TestIgnoreEmptyLinesFlag(t *testing.T) {
	expectLinesContain(t, getLines(5), "Accepted")
	expectLinesContain(t, getLines(5, "-i=0"), "Wrong Answer")
	expectLinesContain(t, getLines(5, "--ignore-empty=0"), "Wrong Answer")
}

func TestLineNumFlag(t *testing.T) {
	expectLinesContainPrefix(t, getLines(2, "-l"), "1\t")
	expectLinesContainPrefix(t, getLines(2, "-l"), "2\t")
	expectLinesContainPrefix(t, getLines(2, "--linenum"), "1\t")
	expectLinesContainPrefix(t, getLines(2, "--linenum"), "2\t")
	expectLinesNotContain(t, getLines(2), "1 ")
	expectLinesNotContain(t, getLines(2), "2 ")
}

func TestPaddingFlagWithoutColor(t *testing.T) {
	expectLinesContainPrefix(t, getLines(2, "-c=0", "-p=0"), "YYYYYN\t\tX")
	expectLinesContainPrefix(t, getLines(2, "-c=0", "-p=1"), "YYYYYN\t\tX")
	expectLinesContainPrefix(t, getLines(2, "-c=0", "-p=5"), "YYYYYN\t\tX")
	expectLinesContainPrefix(t, getLines(2, "-c=0", "-p=6"), "YYYYYN\t\tX")
	expectLinesContainPrefix(t, getLines(2, "-c=0", "-p=7"), "YYYYYN \t\tX")
	expectLinesContainPrefix(t, getLines(2, "-c=0", "-p=8"), "YYYYYN  \t\tX")
	expectLinesContainPrefix(t, getLines(2, "-c=0", "-s", "-p=0"), "Y...\t\tX")
	expectLinesContainPrefix(t, getLines(2, "-c=0", "-s", "-p=1"), "Y...\t\tX")
	expectLinesContainPrefix(t, getLines(2, "-c=0", "-s", "-p=4"), "Y...\t\tX")
	expectLinesContainPrefix(t, getLines(2, "-c=0", "-s", "-p=5"), "YY...\t\tX")
	expectLinesContainPrefix(t, getLines(2, "-c=0", "-s", "-p=6"), "YYYYYN\t\tX")
	expectLinesContainPrefix(t, getLines(2, "-c=0", "-s", "-p=7"), "YYYYYN \t\tX")
	expectLinesContainPrefix(t, getLines(2, "-c=0", "-s", "-p=8"), "YYYYYN  \t\tX")
	expectLinesContainPrefix(t, getLines(2, "-c=0", "-s", "-p=8"), "YYYYY...\t\tX")
}

func TestPaddingFlagWithColor(t *testing.T) {
	build := func(correctPrefix, wrongSuffix string, spaceCount int) string {
		left := ""
		if len(correctPrefix) > 0 {
			left += color.GreenString(correctPrefix)
		}
		if len(wrongSuffix) > 0 {
			left += color.RedString(wrongSuffix)
		}
		return left + strings.Repeat(" ", spaceCount) + "\t\t" + color.RedString("X")
	}
	expectLinesContainPrefix(t, getLines(2, "-p=0"), build("YYYYYN", "", 0))
	expectLinesContainPrefix(t, getLines(2, "-p=6"), build("YYYYYN", "", 0))
	expectLinesContainPrefix(t, getLines(2, "-p=7"), build("YYYYYN", "", 1))
	expectLinesContainPrefix(t, getLines(2, "-p=11"), build("YYYYYN", "", 5))
	expectLinesContainPrefix(t, getLines(2, "-p=0"), build("YYYYYNNN", "NNN", 0))
	expectLinesContainPrefix(t, getLines(2, "-p=11"), build("YYYYYNNN", "NNN", 0))
	expectLinesContainPrefix(t, getLines(2, "-p=12"), build("YYYYYNNN", "NNN", 1))
	expectLinesContainPrefix(t, getLines(2, "-p=17"), build("YYYYYNNN", "NNN", 6))
	expectLinesContainPrefix(t, getLines(2, "-s", "-p=0"), build("", "Y...", 0))
	expectLinesContainPrefix(t, getLines(2, "-s", "-p=5"), build("", "YY...", 0))
	expectLinesContainPrefix(t, getLines(2, "-s", "-p=6"), build("", "YYY...", 0))
	expectLinesContainPrefix(t, getLines(2, "-s", "-p=7"), build("", "YYYYYN", 1))
	expectLinesContainPrefix(t, getLines(2, "-s", "-p=11"), build("", "YYYYYN", 5))
	expectLinesContainPrefix(t, getLines(2, "-s", "-p=11"), build("", "YYYYYNNNNNN", 0))
	expectLinesContainPrefix(t, getLines(2, "-s", "-p=12"), build("", "YYYYYNNNNNN", 1))
	expectLinesContainPrefix(t, getLines(2, "-s", "-p=15"), build("", "YYYYYNNNNNN", 4))
}
