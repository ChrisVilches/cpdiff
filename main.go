package main

import (
	"cpdiff/cli"
	"log"
)

func main() {
	if err := cli.App(); err != nil {
		log.Fatal(err)
	}
}

// TODO: This project needs a massive code review.
// TODO: Do some benchmarks. Execute a program with a lot of data
//       (for example Jumping Grasshopper or Cutting Inequality Down)
//       see if it's too slow compared to executing
//       the program raw (both cases with/without
// cpdiff printing all output to STDOUT).
