package aoc2024day10

import (
	_ "embed"
	"strings"

	"github.com/ysmilda/Advent-of-code/foundation/grid"
	"github.com/ysmilda/Advent-of-code/foundation/solver"
)

//go:embed input.txt
var inputFile string

type puzzle struct {
	grid grid.Grid[byte]
}

func GetSolver() solver.Solver {
	return &puzzle{}
}

func (s puzzle) GetTestInput() string {
	return inputFile
}

func (s puzzle) GetDay() int {
	return 10
}

func (s puzzle) Part1() (int, error) {
	sum := 0

	for y, row := range s.grid {
		for x, c := range row {
			if c != 0 {
				continue
			}

			ends := make(map[grid.Coordinate]bool)
			s.check(grid.NewCoordinate(x, y), func(coord grid.Coordinate) {
				ends[coord] = true
			})

			sum += len(ends)
		}
	}

	return sum, nil
}

func (s puzzle) Part2() (int, error) {
	sum := 0

	for y, row := range s.grid {
		for x, c := range row {
			if c != 0 {
				continue
			}

			s.check(grid.NewCoordinate(x, y), func(coord grid.Coordinate) {
				sum++
			})
		}
	}

	return sum, nil
}

func (s *puzzle) Parse(input string) {
	lines := strings.Split(input, "\n")
	g := grid.NewGrid[byte](uint(len(lines[0])), uint(len(lines)))

	for y, line := range lines {
		for x, c := range line {
			g.Set(grid.NewCoordinate(x, y), byte(c)-'0')
		}
	}

	s.grid = g
}

func (s puzzle) check(coord grid.Coordinate, found func(grid.Coordinate)) {
	if s.grid.Get(coord) == 9 {
		found(coord)
		return
	}

	for _, direction := range grid.CardinalDirections {
		if next := coord.MoveInDirection(direction, 1); s.grid.Valid(next) && s.grid.Get(next) == s.grid.Get(coord)+1 {
			s.check(next, found)
		}
	}
}
