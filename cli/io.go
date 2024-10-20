package cli

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func warn(w error) {
	fmt.Fprint(os.Stderr, "Warning: ")
	fmt.Fprintln(os.Stderr, w)
}

func printfColor(c color.Attribute, useColor bool, s string, printArgs ...any) {
	if !useColor {
		fmt.Printf(s, printArgs...)

		return
	}

	color.New(c).Printf(s, printArgs...)
}
