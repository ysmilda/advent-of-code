package aoc2025day7

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
	return 7
}

func (s puzzle) Part1() (int, error) {
	start, _ := s.grid.Find('S')
	beams := map[grid.Coordinate]bool{
		start: true,
	}

	splitters := map[grid.Coordinate]bool{}
	for range s.grid.GetHeight() {
		newBeams := map[grid.Coordinate]bool{}
		for beam := range beams {
			new := beam.MoveInDirection(grid.South, 1)
			if v, ok := s.grid.GetValid(new); ok && v == '^' {
				splitters[new] = true
				newBeams[new.MoveInDirection(grid.East, 1)] = true
				newBeams[new.MoveInDirection(grid.West, 1)] = true
			} else {
				newBeams[new] = true
			}
		}
		beams = newBeams
	}

	return len(splitters), nil
}

func (s puzzle) Part2() (int, error) {
	start, _ := s.grid.Find('S')
	beams := map[grid.Coordinate]int{
		start: 1,
	}

	for range s.grid.GetHeight() {
		newBeams := map[grid.Coordinate]int{}
		for beam := range beams {
			new := beam.MoveInDirection(grid.South, 1)
			if v, ok := s.grid.GetValid(new); ok && v == '^' {
				newBeams[new.MoveInDirection(grid.East, 1)] += beams[beam]
				newBeams[new.MoveInDirection(grid.West, 1)] += beams[beam]
			} else {
				newBeams[new] += beams[beam]
			}
		}
		beams = newBeams
	}

	sum := 0
	for _, i := range beams {
		sum += i
	}
	return sum, nil
}

func (s puzzle) follow(c grid.Coordinate) int {
	if c.Y == s.grid.GetHeight() {
		return 0
	}
	sum := 0

	for range s.grid.GetHeight() - c.Y {
		c = c.MoveInDirection(grid.South, 1)
		if v, ok := s.grid.GetValid(c); ok && v == '^' {
			sum += 1
			sum += s.follow(c.MoveInDirection(grid.East, 1))
			sum += s.follow(c.MoveInDirection(grid.West, 1))
			return sum
		}
	}

	return sum
}

func (s *puzzle) Parse(input string) {
	s.grid = grid.NewGridFromLines(strings.Split(input, "\n"), func(r rune) byte { return byte(r) })
}
