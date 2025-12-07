package main

import (
	"flag"
	"fmt"
	"log"

	"advent2025/internal"
	"advent2025/pkg/util"
)

func main() {
	day := flag.Int("day", 1, "Which Advent of Code day to run")
	flag.Parse()

	inputPath := fmt.Sprintf("inputs/day%02d.txt", *day)
	input := util.ReadInput(inputPath)

	solver := internal.GetSolver(*day)
	if solver == nil {
		log.Fatalf("No solver registered for day %d. Use ctl to generate it first.", *day)
	}

	fmt.Println(solver(input))
}
