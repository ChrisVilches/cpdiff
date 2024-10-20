package integration

import (
	"fmt"
	"github.com/creack/pty"
	"github.com/fatih/color"
	"io"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func readFileLines(f *os.File) []string {
	buf := strings.Builder{}
	_, err := io.Copy(&buf, f)

	if err != nil {
		// TODO: It panics, but it still reads....
		// panic(err)
	}

	trimmed := strings.TrimSpace(buf.String())
	lines := strings.Split(trimmed, "\n")

	for i := range len(lines) {
		lines[i] = strings.TrimSpace(lines[i])
	}

	return lines
}

func expectLineCount(t *testing.T, lines []string, count int) {
	if len(lines) != count {
		t.Fatalf("expected output to have %d lines, but has %d", count, len(lines))
	}
}

func expectLinesNotContain(t *testing.T, lines []string, s string) {
	for _, line := range lines {
		if strings.Contains(line, s) {
			t.Fatalf("expected output not to contain text %s", s)
		}
	}

}

func expectLinesContain(t *testing.T, lines []string, s string) {
	for _, line := range lines {
		if strings.Contains(line, s) {
			return
		}
	}

	t.Fatalf("expected output to contain text %s", s)
}

func test(t *testing.T, actual, expected string) {
	if expected != actual {
		t.Fatalf("expected text to be '%s', but got '%s'", expected, actual)
	}
}

func getLines(testCaseNum int, flags ...string) []string {
	/// TODO: Verify this line (removed) breaks everything.
	color.NoColor = false
	executable := strings.TrimSpace(os.Getenv("CPDIFF_INTEGRATION_TEST_EXECUTABLE"))

	if len(executable) == 0 {
		panic("Executable is not defined")
	}

	inFile := fmt.Sprintf("./data/%d/in", testCaseNum)
	ansFile := fmt.Sprintf("./data/%d/ans", testCaseNum)

	flags = append(flags, inFile)
	flags = append(flags, ansFile)

	cmd := exec.Command(executable, flags...)
	f, err := pty.Start(cmd)

	if err != nil {
		panic(err)
	}

	lines := readFileLines(f)
	f.Close()
	return lines
}
