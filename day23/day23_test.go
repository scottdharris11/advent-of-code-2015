package day23

import (
	"advent-of-code-2015/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testLines = []string{
	"inc a",
	"jio a, +2",
	"tpl a",
	"inc a",
}

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, 2, solvePart1(testLines, "a"))
	assert.Equal(t, 170, solvePart1(utils.ReadLines("day23", "day-23-input.txt"), "b"))
}

func TestSolvePart2(t *testing.T) {
	assert.Equal(t, 247, solvePart2(utils.ReadLines("day23", "day-23-input.txt"), "b"))
}
