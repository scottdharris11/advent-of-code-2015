package day16

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"advent-of-code-2015/utils"
)

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, 373, solvePart1(utils.ReadLines("day16", "day-16-input.txt")))
}

func TestSolvePart2(t *testing.T) {
	assert.Equal(t, 260, solvePart2(utils.ReadLines("day16", "day-16-input.txt")))
}
