package integration

import (
	"testing"
)

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
	expectEq(t, len(getLines(1, "-w")), 2)
	expectEq(t, len(getLines(1, "--wrong")), 2)
	expectEq(t, len(getLines(1)), 6)
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
	expectLinesContain(t, getLines(2, "-l"), "1\t")
	expectLinesContain(t, getLines(2, "-l"), "2\t")
	expectLinesContain(t, getLines(2, "--linenum"), "1\t")
	expectLinesContain(t, getLines(2, "--linenum"), "2\t")
	expectLinesNotContain(t, getLines(2), "1 ")
	expectLinesNotContain(t, getLines(2), "2 ")
}
