package aoc2025day2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInput = `11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124`

func TestGetDay(t *testing.T) {
	solver := MustGetSolver()
	day := solver.GetDay()
	assert.Equal(t, 2, day)
}

func TestPart1(t *testing.T) {
	solver := puzzle{
		input: parse(testInput),
	}

	result, _ := solver.Part1()
	assert.Equal(t, 1227775554, result)
}

func TestPart2(t *testing.T) {
	solver := puzzle{
		input: parse(testInput),
	}

	result, _ := solver.Part2()
	assert.Equal(t, 4174379265, result)
}

func BenchmarkPart1(b *testing.B) {
	solver := puzzle{
		input: parse(inputFile),
	}
	for b.Loop() {
		res, _ := solver.Part1()
		_ = res
	}
}
func BenchmarkPart2(b *testing.B) {
	solver := puzzle{
		input: parse(inputFile),
	}
	for b.Loop() {
		res, _ := solver.Part2()
		_ = res
	}
}
