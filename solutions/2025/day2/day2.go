package aoc2025day2

import (
	_ "embed"
	"strconv"
	"strings"

	"github.com/ysmilda/Advent-of-code/foundation/aocmath"
	"github.com/ysmilda/Advent-of-code/foundation/aocslices"
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

			if a, b := aocmath.Split(i); a == b {
				sum += i
			}
		}
	}

	return sum, nil
}

func (s puzzle) Part2() (int, error) {
	sum := 0
	parts := make([]string, 0, 100)

	for _, r := range s.input {
		for i := r.min; i <= r.max; i++ {
			s := strconv.Itoa(i)

			half := len(s) / 2
			for n := 1; n <= half; n++ {
				// split s into parts of size n and compare if all parts are equal
				if len(s)%n != 0 {
					continue
				}

				parts := parts[:0]
				for i := 0; i <= len(s)-n; i += n {
					parts = append(parts, s[i:i+n])
				}

				if aocslices.AllEntriesEqual(parts) {
					sum += i
					break
				}
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
