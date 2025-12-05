package aoc2023day7

import (
	"cmp"
	_ "embed"
	"slices"
	"strings"

	"github.com/ysmilda/Advent-of-code/foundation/aocstrconv"
	"github.com/ysmilda/Advent-of-code/foundation/solver"
)

type hand struct {
	strength int
	value    int
	bid      int
}

//go:embed input.txt
var inputFile string

type puzzle struct {
	input string
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
	sum := 0
	input := sort(parse(s.input, false, "23456789TJQKA"))
	for i, in := range input {
		sum += (in.bid * (i + 1))
	}

	return sum, nil
}

func (s puzzle) Part2() (int, error) {
	sum := 0
	input := sort(parse(s.input, true, "J23456789TQKA"))
	for i, in := range input {
		sum += (in.bid * (i + 1))
	}

	return sum, nil
}

func (s *puzzle) Parse(input string) {
	s.input = input
}

func parse(input string, useJokers bool, values string) []hand {
	lines := strings.Split(input, "\n")

	output := []hand{}
	for _, line := range lines {
		output = append(output, parseHand(line, useJokers, values))
	}
	return output
}

func parseHand(line string, useJokers bool, values string) hand {
	parts := strings.Fields(line)

	cards := make(map[byte]int)
	cardValue := 0
	cardStrength := 0
	highest := 0

	for _, c := range parts[0] {
		cards[byte(c)]++
		cardValue = (cardValue + strings.IndexByte(values, byte(c))) * 100

		val := cards[byte(c)]
		cardStrength += val
		if val > highest && c != 'J' {
			highest = val
		}
	}

	if useJokers {
		jokers := cards['J']
		cardStrength += (jokers * highest)
	}

	return hand{
		bid:      aocstrconv.MustAtoi(parts[1]),
		strength: cardStrength,
		value:    cardValue,
	}
}

func sort(input []hand) []hand {
	slices.SortFunc(input, func(a, b hand) int {
		i := cmp.Compare(a.strength, b.strength)
		if i == 0 {
			return cmp.Compare(a.value, b.value)
		}
		return i
	})
	return input
}
