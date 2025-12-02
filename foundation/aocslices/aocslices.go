package aocslices

import (
	"iter"

	"github.com/ysmilda/Advent-of-code/foundation/aocmath"
)

// Remove removes the element at index from a slice.
func Remove[T []E, E any](slice T, index int) T {
	return append(slice[:index], slice[index+1:]...)
}

// Iterate returns an iterater which iterates a number in a given base and returns each digit as part of a slice.
func Iterate(base, length int) iter.Seq2[int, []int] {
	return func(yield func(int, []int) bool) {
		steps := make([]int, length)
		for i := range aocmath.Pow(base, length) {
			if !yield(i, steps) {
				return
			}
			for j := range steps {
				steps[j]++
				if steps[j] > base-1 {
					steps[j] = 0
				} else {
					break
				}
			}
		}
	}
}

// AllEqual checks if all entries in the slice are equal.
// If the given slice is nil or has less than two entries it returns true.
func AllEqual[T []E, E comparable](slice T) bool {
	if len(slice) < 2 {
		return true
	}

	last := slice[0]
	for _, item := range slice[1:] {
		if last != item {
			return false
		}
		last = item
	}

	return true
}
