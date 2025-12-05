package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/ysmilda/Advent-of-code/foundation/solver"
)

var solvers map[int][]solver.Solver

func main() {
	year := flag.Int("year", time.Now().Year(), "the year to run")
	day := flag.Int("day", 0, "the solver to run, when set to 0 all solvers will be run")
	flag.Parse()

	s, ok := solvers[*year]
	if !ok {
		panic("Year not implemented yet")
	}

	ran := false
	for _, solver := range s {
		if solver.GetDay() == *day || *day == 0 {
			executeSolver(solver)
			ran = true
		}
	}

	if !ran {
		fmt.Println("Day not implemented yet")
	}
}

func executeSolver(solver solver.Solver) {
	printDay(solver.GetDay())

	startParse := time.Now()
	solver.Parse(solver.GetTestInput())
	printParse(time.Since(startParse))

	startPart1 := time.Now()
	solution, err := solver.Part1()
	if err != nil {
		fmt.Printf("Failed to solve part 1 of day %d: %s\n", solver.GetDay(), err)
	}
	printPart(1, solution, time.Since(startPart1))

	startPart2 := time.Now()
	solution, err = solver.Part2()
	if err != nil {
		fmt.Printf("Failed to solve part 1 of day %d: %s\n", solver.GetDay(), err)
	}
	printPart(2, solution, time.Since(startPart2))
}

func printDay(day int) {
	fmt.Printf("Day %d:\n", day)
}

func printPart(part int, solution int, duration time.Duration) {
	fmt.Printf("\tPart %d:\n\t\tsolution:\t%d\n\t\truntime:\t%s\n", part, solution, duration.String())
}

func printParse(duration time.Duration) {
	fmt.Printf("\tParsing:\n\t\truntime:\t%s\n", duration.String())
}
