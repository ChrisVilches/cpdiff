package util

import "strings"

func IsEmptyLine(line string) bool {
	return len(strings.TrimSpace(line)) == 0
}
