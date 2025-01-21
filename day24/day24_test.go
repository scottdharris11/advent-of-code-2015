package day24

import (
	"advent-of-code-2015/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testWeights = []int{1, 2, 3, 4, 5, 7, 8, 9, 10, 11}

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, 99, solvePart1(testWeights))
	assert.Equal(t, 11846773891, solvePart1(utils.ReadIntegers("day24", "day-24-input.txt")))
}

func TestSolvePart2(t *testing.T) {
	assert.Equal(t, 0, solvePart2(utils.ReadIntegers("day24", "day-24-input.txt")))
}
