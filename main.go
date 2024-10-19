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
// TODO: Here's a good benchmark to optimize it.
// time c++ urionlinejudge/3021-jumping_grasshopper.cpp
//    < ~/dev/data/J/input/J_29
// c++ urionlinejudge/3021-jumping_grasshopper.cpp
//    < ~/dev/data/J/input/J_29 | cpdiff -d ~/dev/data/J/output/J_29
// 698.490613ms with cpdiff, 0.17s without
// I should be able to get a better time (but it won't be exactly the same).
// However, before micro-optimizing, I need to test everything
// so that it doesn't break while changing it.
