package cli

import (
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
)

func warn(msg string) {
	fmt.Fprint(os.Stderr, "Warning: ")
	fmt.Fprintln(os.Stderr, msg)
}

func openFileOrExit(filePath string) *os.File {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("File %s cannot be opened", filePath)
	}
	return file
}

func printfColor(c color.Attribute, useColor bool, s string, printArgs ...interface{}) {
	if !useColor {
		fmt.Printf(s, printArgs...)
		return
	}

	color.New(c).Printf(s, printArgs...)
}
