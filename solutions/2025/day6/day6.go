package aoc2025day6

import (
	_ "embed"
	"strings"

	"github.com/ysmilda/Advent-of-code/foundation/aocstrconv"
	"github.com/ysmilda/Advent-of-code/foundation/solver"
)

//go:embed input.txt
var inputFile string

type puzzle struct {
	input string
}

type problem struct {
	numbers  []int
	multiply bool
}

func (p problem) solve() int {
	solution := p.numbers[0]
	for _, number := range p.numbers[1:] {
		if p.multiply {
			solution *= number
		} else {
			solution += number
		}
	}
	return solution
}

func GetSolver() solver.Solver {
	return &puzzle{}
}

func (s puzzle) GetTestInput() string {
	return inputFile
}

func (s puzzle) GetDay() int {
	return 6
}

func (s puzzle) Part1() (int, error) {
	sum := 0

	for _, problem := range parseNormal(s.input) {
		sum += problem.solve()
	}

	return sum, nil
}

func (s puzzle) Part2() (int, error) {
	sum := 0

	for _, problem := range parseRightToLeft(s.input) {
		sum += problem.solve()
	}

	return sum, nil
}

func (s *puzzle) Parse(input string) {
	s.input = input
}

func parseNormal(input string) []problem {
	removeEmpty := func(input []string) []string {
		out := make([]string, 0, len(input))
		for i := range input {
			if input[i] == "" {
				continue
			}
			out = append(out, input[i])
		}
		return out
	}

	lines := strings.Split(input, "\n")
	numbers := [][]int{}
	operations := removeEmpty(strings.Split(lines[len(lines)-1], " ")) // The last line contains the operations
	for _, line := range lines[:len(lines)-1] {
		numbers = append(numbers, aocstrconv.MustAtoiSlice(removeEmpty(strings.Split(line, " ")))) // The other lines the numbers
	}

	result := make([]problem, 0, len(operations))
	for i := range operations {
		problem := problem{
			multiply: operations[i] == "*",
		}
		for j := range numbers {
			problem.numbers = append(problem.numbers, numbers[j][i])
		}

		result = append(result, problem)
	}

	return result
}

func parseRightToLeft(input string) []problem {
	lines := strings.Split(input, "\n")
	result := []problem{{}}
	last := &result[len(result)-1]

	// Traverse backwards through the input (right to left) and read the numbers from top to bottom.
	for i := len(lines[0]) - 1; i >= 0; i-- {
		number := make([]byte, 0, len(lines))
		for j := 0; j < len(lines)-1; j++ {
			number = append(number, lines[j][i])
		}

		// Parse the found number and add it to the problem
		last.numbers = append(last.numbers, aocstrconv.MustAtoi(strings.TrimSpace(string(number))))

		if lines[len(lines)-1][i] != ' ' {
			last.multiply = lines[len(lines)-1][i] == '*' // Store the operation
			result = append(result, problem{})            // Add a new problem to the output
			last = &result[len(result)-1]                 // Make the last entry easy to target
			i--                                           // Skip the line with spaces
		}
	}

	return result[:len(result)-1] // An extra problem was appended, remove that
}
