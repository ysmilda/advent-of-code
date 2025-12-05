package aoc2025day3

import (
	_ "embed"
	"strings"

	"github.com/ysmilda/Advent-of-code/foundation/aocmath"
	"github.com/ysmilda/Advent-of-code/foundation/aocstrconv"
	"github.com/ysmilda/Advent-of-code/foundation/char"
	"github.com/ysmilda/Advent-of-code/foundation/solver"
)

//go:embed input.txt
var inputFile string

type puzzle struct {
	input []string
}

func GetSolver() solver.Solver {
	return &puzzle{}
}

func (s puzzle) GetTestInput() string {
	return inputFile
}

func (s puzzle) GetDay() int {
	return 3
}

func (s puzzle) Part1() (int, error) {
	joltage := 0

	for _, bank := range s.input {
		joltage += findLargestJoltage(bank, 2)
	}

	return joltage, nil
}

func (s puzzle) Part2() (int, error) {
	joltage := 0

	for _, bank := range s.input {
		joltage += findLargestJoltage(bank, 12)
	}

	return joltage, nil
}

type args struct {
	bank   string
	length int
}

var cache = map[args]int{}

func findLargestJoltage(bank string, length int) int {
	if result, ok := cache[args{bank, length}]; ok {
		return result
	}
	if length == 0 {
		return 0
	}
	if len(bank) == length {
		return aocstrconv.MustAtoi(bank)
	}

	a := (char.ToInt(bank[0]) * aocmath.Pow(10, length-1)) + findLargestJoltage(bank[1:], length-1)
	b := findLargestJoltage(bank[1:], length)

	result := aocmath.Max(a, b)
	cache[args{bank, length}] = result

	return result
}

func (s *puzzle) Parse(input string) {
	s.input = strings.Split(input, "\n")
}
