package aoc2025day5

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInput = `3-5
10-14
16-20
12-18

1
5
8
11
17
32`

func TestGetDay(t *testing.T) {
	solver := GetSolver()
	day := solver.GetDay()
	assert.Equal(t, 5, day)
}

func TestPart1(t *testing.T) {
	solver := puzzle{}
	solver.Parse(testInput)

	result, _ := solver.Part1()
	assert.Equal(t, 3, result)
}

func TestPart2(t *testing.T) {
	solver := puzzle{}
	solver.Parse(testInput)

	result, _ := solver.Part2()
	assert.Equal(t, 14, result)
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
