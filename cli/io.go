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

// TODO: This now prints new line on its own.
// This is to avoid coloring the newline, which kinda fucks up the output.
// visually it's ok but it's hard to test
func printfColor(c color.Attribute, useColor bool, s string, printArgs ...any) {
	if !useColor {
		fmt.Println(fmt.Sprintf(s, printArgs...))

		return
	}

	fmt.Println(color.New(c).Sprintf(s, printArgs...))
}
