package util

func RemoveTrailingNewLine(line string) string {
	if len(line) > 0 && line[len(line)-1] == '\n' {
		return line[:len(line)-1]
	}

	return line
}
