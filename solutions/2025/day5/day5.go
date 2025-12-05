package aoc2025day5

import (
	"cmp"
	_ "embed"
	"slices"
	"strings"

	"github.com/ysmilda/Advent-of-code/foundation/aocstrconv"
	"github.com/ysmilda/Advent-of-code/foundation/solver"
)

//go:embed input.txt
var inputFile string

type puzzle struct {
	input database
}

type database struct {
	freshIDs []*freshIDrange
	IDs      []int
}

type freshIDrange struct {
	deleted    bool
	start, end int
}

func MustGetSolver() solver.Solver {
	return puzzle{
		input: parse(inputFile),
	}
}

func (s puzzle) GetDay() int {
	return 5
}

func (s puzzle) Part1() (int, error) {
	sum := 0

	for _, id := range s.input.IDs {
		for _, r := range s.input.freshIDs {
			if id >= r.start && id <= r.end {
				sum++
				break
			}
		}
	}

	return sum, nil
}

func (s puzzle) Part2() (int, error) {
	sum := 0

	slices.SortFunc(s.input.freshIDs, func(a, b *freshIDrange) int {
		return cmp.Compare(a.start, b.start)
	})

	for i, r1 := range s.input.freshIDs {
		if r1.deleted {
			continue
		}

		for _, r2 := range s.input.freshIDs[i+1:] {
			if r2.deleted {
				continue
			}

			if r2.start > r1.end {
				break
			}

			if r2.end > r1.end {
				r1.end = r2.end
			}

			r2.deleted = true
		}

		sum += 1 + r1.end - r1.start
	}

	return sum, nil
}

func parse(input string) database {
	lines := strings.Split(input, "\n")
	out := database{}

	for i, line := range lines {
		if line == "" {
			out.IDs = aocstrconv.MustAtoiSlice(lines[i+1:])
			break
		}

		parts := strings.Split(line, "-")
		out.freshIDs = append(out.freshIDs, &freshIDrange{
			start: aocstrconv.MustAtoi(parts[0]),
			end:   aocstrconv.MustAtoi(parts[1]),
		})
	}

	return out
}
