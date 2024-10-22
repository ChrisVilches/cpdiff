package integration

import (
	"testing"
)

var usageTestFlags = []string{"-c=0", "-p=0"}

func TestStdinArg(t *testing.T) {
	lines, statusCode := runCmd("-", "./data/1/ans", "hello\nworld\nbye", usageTestFlags...)
	expectEq(t, lines[0], "hello\t\tX  AAABBBCCC")
	expectEq(t, lines[1], "world\t\tX  123123 123123")
	expectEq(t, lines[2], "bye\t\tX  4455.555555    8888.9999")
	expectEq(t, statusCode, 1)
	expectEq(t, len(lines), 5)
}

func TestStdinArg2(t *testing.T) {
	lines, statusCode := runCmd("./data/1/ans", "-", "hello\nworld\nbye", usageTestFlags...)
	expectEq(t, lines[0], "AAABBBCCC\t\tX  hello")
	expectEq(t, lines[1], "123123 123123\t\tX  world")
	expectEq(t, lines[2], "4455.555555    8888.9999\t\tX  bye")
	expectEq(t, statusCode, 1)
	expectEq(t, len(lines), 5)
}

func TestStdinArgCorrectAnswer(t *testing.T) {
	stdin := "1 3    2.00001   3.00001 4"
	lines, statusCode := runCmd("-", "./data/3/ans", stdin, usageTestFlags...)
	expectEq(t, lines[0], "1 3    2.00001   3.00001 4\t\t≈  1 3 2.00001   3   4")
	expectEq(t, statusCode, 0)
	expectEq(t, len(lines), 4)
}

func TestStdinArgCorrectAnswer2(t *testing.T) {
	stdin := "1 3    2   3.00001 4"
	lines, statusCode := runCmd("./data/3/ans", "-", stdin, usageTestFlags...)
	expectEq(t, lines[0], "1 3 2.00001   3   4\t\t≈  1 3    2   3.00001 4")
	expectEq(t, statusCode, 0)
	expectEq(t, len(lines), 4)
}

func TestStdinArgError(t *testing.T) {
	lines, statusCode := runCmd("-", "-", "dummy", usageTestFlags...)
	expectEq(t, lines[0], "Error: Do not use - (standard input) for both sides")
	expectEq(t, statusCode, 1)
	expectEq(t, len(lines), 1)
}
