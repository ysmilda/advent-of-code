package aoc2025day1

import (
	_ "embed"
	"strings"

	"github.com/ysmilda/Advent-of-code/foundation/aocmath"
	"github.com/ysmilda/Advent-of-code/foundation/aocstrconv"
	"github.com/ysmilda/Advent-of-code/foundation/solver"
)

//go:embed input.txt
var inputFile string

type puzzle struct {
	input []int
}

func GetSolver() solver.Solver {
	return &puzzle{}
}

func (s puzzle) GetTestInput() string {
	return inputFile
}

func (s puzzle) GetDay() int {
	return 1
}

func (s puzzle) Part1() (int, error) {
	sum := 0

	dial := 50
	for _, i := range s.input {
		dial = (dial + i) % 100

		if dial == 0 {
			sum++
		}
	}

	return sum, nil
}

func (s puzzle) Part2() (int, error) {
	sum := 0

	dial := 50
	for _, i := range s.input {
		if (dial > 0 && dial+i <= 0) || (dial < 0 && dial+i >= 0) {
			sum++
		}

		dial += i
		sum += aocmath.Abs(dial / 100)
		dial %= 100
	}

	return sum, nil
}

func (s *puzzle) Parse(input string) {
	lines := strings.Split(input, "\n")
	result := make([]int, 0, len(lines))
	for _, line := range lines {
		value := aocstrconv.MustAtoi(line[1:])
		if line[0] == 'L' {
			value *= -1
		}
		result = append(result, value)
	}
	s.input = result
}
