package integration

import (
	"cpdiff/cli"
	"fmt"
	"github.com/fatih/color"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func expectLineCount(t *testing.T, lines []string, count int) {
	if len(lines) != count {
		msg := "expected output to have %d lines, but has %d"
		t.Fatalf(msg, count, len(lines))
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

func bytesIntoLines(bytes []byte) []string {
	trimmed := strings.TrimSpace(string(bytes))
	lines := strings.Split(trimmed, "\n")

	for i := range lines {
		lines[i] = strings.TrimSpace(lines[i])
	}

	return lines
}

func getExecutableName() string {
	executable := strings.TrimSpace(os.Getenv("INTEGRATION_TEST_EXECUTABLE"))

	if len(executable) == 0 {
		panic("Executable is not defined")
	}

	return executable
}

func getLines(testCaseNum int, flags ...string) []string {
	color.NoColor = false

	inFile := fmt.Sprintf("./data/%d/in", testCaseNum)
	ansFile := fmt.Sprintf("./data/%d/ans", testCaseNum)

	flags = append(flags, inFile)
	flags = append(flags, ansFile)

	cmd := exec.Command(getExecutableName(), flags...)
	cmd.Env = append(cmd.Env, fmt.Sprintf("%s=1", cli.ForceColorFlag))

	bytes, err := cmd.Output()

	if err != nil {
		fmt.Fprintf(os.Stderr, "failed: %s\n", flags)
		panic(err)
	}

	return bytesIntoLines(bytes)
}
