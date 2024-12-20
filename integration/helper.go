package integration

import (
	"bufio"
	"fmt"
	"github.com/ChrisVilches/cpdiff/cli"
	"github.com/fatih/color"
	"io"
	"os"
	"os/exec"
	"strings"
	"testing"
)

const noPaddingFlag = "-p=0"

func expectEq[T comparable](t *testing.T, a, b T) {
	if a != b {
		t.Fatalf("expected '%v' to be equal to '%v'", a, b)
	}
}

func expectLinesNotContain(t *testing.T, lines []string, s string) {
	for _, line := range lines {
		if strings.Contains(line, s) {
			t.Fatalf("expected lines not to contain text %s", s)
		}
	}
}

func expectLinesContainPrefix(t *testing.T, lines []string, s string) {
	for _, line := range lines {
		if strings.HasPrefix(line, s) {
			return
		}
	}

	t.Fatalf("expected lines to contain line with prefix %s", s)
}

func expectLinesContain(t *testing.T, lines []string, s string) {
	for _, line := range lines {
		if strings.Contains(line, s) {
			return
		}
	}

	t.Fatalf("expected lines to contain line with content %s", s)
}

func bytesIntoLines(bytes []byte) []string {
	trimmed := strings.TrimSpace(string(bytes))

	if len(trimmed) == 0 {
		return nil
	}

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

func readLines(r io.ReadCloser) []string {
	buf := bufio.NewScanner(r)
	res := []string{}
	for buf.Scan() {
		res = append(res, buf.Text())
	}
	return res
}

func passStdin(stdin string, cmd *exec.Cmd) {
	stdinPipe, err := cmd.StdinPipe()
	if err != nil {
		panic(err)
	}

	// Don't add newline, to verify the behavior when
	// the last line ends with EOF instead of newline.
	fmt.Fprint(stdinPipe, stdin)
	_ = stdinPipe.Close()
}

func runCmd(inFile, ansFile, stdin string, flags ...string) ([]string, int) {
	color.NoColor = false

	flags = append(flags, inFile)
	flags = append(flags, ansFile)

	cmd := exec.Command(getExecutableName(), flags...)
	cmd.Env = append(cmd.Env, fmt.Sprintf("%s=1", cli.ForceColorFlag))

	passStdin(stdin, cmd)
	out, err := cmd.CombinedOutput()

	statusCode := 0

	if exitError, ok := err.(*exec.ExitError); ok {
		statusCode = exitError.ExitCode()
	} else if err != nil {
		fmt.Fprintf(os.Stderr, "failed: %s\n", flags)
		panic(err)
	}

	return bytesIntoLines(out), statusCode
}

func getLines(testCaseNum int, flags ...string) []string {
	inFile := fmt.Sprintf("./data/%d/in", testCaseNum)
	ansFile := fmt.Sprintf("./data/%d/ans", testCaseNum)
	lines, _ := runCmd(inFile, ansFile, "", flags...)
	return lines
}

func getStatusCode(testCaseNum int, flags ...string) int {
	inFile := fmt.Sprintf("./data/%d/in", testCaseNum)
	ansFile := fmt.Sprintf("./data/%d/ans", testCaseNum)
	_, statusCode := runCmd(inFile, ansFile, "", flags...)
	return statusCode
}
