package day6

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"advent-of-code-2015/utils"
)

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, 400410, solvePart1(utils.ReadLines("day6", "day-6-input.txt")))
}

func TestSolvePart2(t *testing.T) {
	assert.Equal(t, 15343601, solvePart2(utils.ReadLines("day6", "day-6-input.txt")))
}
