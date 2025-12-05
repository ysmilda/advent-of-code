package aoc2025day3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInput = `987654321111111
811111111111119
234234234234278
818181911112111`

func TestGetDay(t *testing.T) {
	solver := GetSolver()
	day := solver.GetDay()
	assert.Equal(t, 3, day)
}

func TestPart1(t *testing.T) {
	solver := puzzle{}
	solver.Parse(testInput)

	result, _ := solver.Part1()
	assert.Equal(t, 357, result)
}

func TestPart2(t *testing.T) {
	solver := puzzle{}
	solver.Parse(testInput)

	result, _ := solver.Part2()
	assert.Equal(t, 3121910778619, result)
}

func BenchmarkPart1(b *testing.B) {
	solver := puzzle{}
	solver.Parse(inputFile)

	for b.Loop() {
		res, _ := solver.Part1()
		_ = res
	}
}

func BenchmarkPart2(b *testing.B) {
	solver := puzzle{}
	solver.Parse(inputFile)

	for b.Loop() {
		res, _ := solver.Part2()
		_ = res
	}
}
