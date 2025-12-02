package aoc2025day2

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
	input []idrange
}

type idrange struct {
	min, max int
}

func MustGetSolver() solver.Solver {
	return puzzle{
		input: parse(inputFile),
	}
}

func (s puzzle) GetDay() int {
	return 2
}

func (s puzzle) Part1() (int, error) {
	sum := 0
	for _, r := range s.input {
		for i := r.min; i <= r.max; i++ {
			n := aocmath.NumberOfDigits(i)
			if n%2 != 0 { // An uneven number of digits can't have been mirrored around the center.
				continue
			}

			tens := aocmath.Pow(10, n/2)
			a := i % tens
			b := i / tens
			if a == b {
				sum += i
			}
		}
	}

	return sum, nil
}

func (s puzzle) Part2() (int, error) {
	sum := 0

	for _, r := range s.input {
		for value := r.min; value <= r.max; value++ {
			digits := aocmath.NumberOfDigits(value)
			for partSize := 1; partSize <= digits/2; partSize++ {
				if digits%partSize != 0 { // Check if i can be split into parts of size n without remainder
					continue
				}

				tens := aocmath.Pow(10, partSize)
				first := value % tens
				temp := value / tens
				for range (digits / partSize) - 1 {
					current := temp % tens
					if current != first {
						goto skip
					}
					temp /= tens
				}

				sum += value
				break
			skip:
			}
		}
	}

	return sum, nil
}

func parse(input string) []idrange {
	ranges := strings.Split(input, ",")
	out := make([]idrange, 0, len(ranges))
	for _, r := range ranges {
		parts := strings.Split(r, "-")
		out = append(out, idrange{
			min: aocstrconv.MustAtoi(parts[0]),
			max: aocstrconv.MustAtoi(parts[1]),
		})
	}
	return out
}
