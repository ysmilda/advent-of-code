package aoc2025day4

import (
	_ "embed"
	"strings"

	"github.com/ysmilda/Advent-of-code/foundation/grid"
	"github.com/ysmilda/Advent-of-code/foundation/solver"
)

//go:embed input.txt
var inputFile string

type puzzle struct {
	input grid.Grid[bool]
}

func MustGetSolver() solver.Solver {
	return puzzle{
		input: parse(inputFile),
	}
}

func (s puzzle) GetDay() int {
	return 4
}

func (s puzzle) Part1() (int, error) {
	return len(findRolls(s.input)), nil
}

func (s puzzle) Part2() (int, error) {
	sum := 0

	g := grid.CopyFrom(s.input)

	for {
		rolls := findRolls(g)
		if len(rolls) == 0 {
			break
		}

		for _, c := range rolls {
			g.Set(c, false)
		}

		sum += len(rolls)
	}

	return sum, nil
}

func findRolls(g grid.Grid[bool]) []grid.Coordinate {
	rolls := []grid.Coordinate{}
	for c, v := range g.Iterate {
		if !v {
			continue
		}

		// Check if it has at most 3 neighbouring rolls of paper
		neighbours := 0
		for _, direction := range grid.AllDirections {
			if next := c.MoveInDirection(direction, 1); g.Valid(next) && g.Get(next) {
				neighbours++
				if neighbours > 3 {
					goto skip
				}
			}
		}

		rolls = append(rolls, c)

	skip:
	}

	return rolls
}

func parse(input string) grid.Grid[bool] {
	return grid.NewGridFromLines(strings.Split(input, "\n"), func(i rune) bool {
		return i == '@'
	})
}
