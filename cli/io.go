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

func printfLnColor(c color.Attribute, useColor bool, s string, args ...any) {
	if !useColor {
		fmt.Println(fmt.Sprintf(s, args...))

		return
	}

	fmt.Println(color.New(c).Sprintf(s, args...))
}
