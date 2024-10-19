package cli

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func warn(msg string) {
	fmt.Fprint(os.Stderr, "Warning: ")
	fmt.Fprintln(os.Stderr, msg)
}

func printfColor(c color.Attribute, useColor bool, s string, printArgs ...any) {
	if !useColor {
		fmt.Printf(s, printArgs...)

		return
	}

	color.New(c).Printf(s, printArgs...)
}
