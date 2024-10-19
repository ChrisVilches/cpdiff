package main

import (
	"cpdiff/cli"
)

func main() {
	cli.App()
}

// TODO: This project needs a massive code review.
// TODO: Also use a linter to find coding mistakes
// TODO: Do some benchmarks. Execute a program with a lot of data
//       (for example Jumping Grasshopper or Cutting Inequality Down)
//       see if it's too slow compared to executing
//       the program raw (both cases with/without cpdiff printing all output to STDOUT).
