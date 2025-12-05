package aoc2024day1

import (
	_ "embed"
	"errors"
	"slices"
	"strings"

	"github.com/ysmilda/Advent-of-code/foundation/aocmath"
	"github.com/ysmilda/Advent-of-code/foundation/aocstrconv"
	"github.com/ysmilda/Advent-of-code/foundation/solver"
)

//go:embed input.txt
var inputFile string

type puzzle struct {
	left  []int
	right []int
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

	if len(s.left) != len(s.right) {
		return 0, errors.New("arrays should be equal length")
	}

	slices.Sort(s.left)
	slices.Sort(s.right)

	for i := range s.left {
		sum += aocmath.Abs(s.left[i] - s.right[i])
	}

	return sum, nil
}

func (s puzzle) Part2() (int, error) {
	sum := 0

	similarity := make(map[int]int)

	for _, value := range s.right {
		similarity[value]++
	}

	for _, value := range s.left {
		sum += similarity[value] * value
	}

	return sum, nil
}

func (s *puzzle) Parse(input string) {
	lines := strings.Split(input, "\n")
	left, right := []int{}, []int{}
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		elements := strings.Split(line, "  ")
		left = append(left, aocstrconv.MustAtoi(strings.TrimSpace(elements[0])))
		right = append(right, aocstrconv.MustAtoi(strings.TrimSpace(elements[1])))
	}

	s.left, s.right = left, right
}
