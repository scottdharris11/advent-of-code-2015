package day19

import (
	"advent-of-code-2015/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testLines = []string{
	"H => HO",
	"H => OH",
	"O => HH",
	"",
	"HOH",
}

var testLines2 = []string{
	"H => HO",
	"H => OH",
	"O => HH",
	"",
	"HOHOHO",
}

var testLines3 = []string{
	"e => H",
	"e => O",
	"H => HO",
	"H => OH",
	"O => HH",
	"",
	"HOH",
}

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, 4, solvePart1(testLines))
	assert.Equal(t, 7, solvePart1(testLines2))
	assert.Equal(t, 576, solvePart1(utils.ReadLines("day19", "day-19-input.txt")))
}

func TestSolvePart2(t *testing.T) {
	assert.Equal(t, 3, solvePart2(testLines3))
	assert.Equal(t, 207, solvePart2(utils.ReadLines("day19", "day-19-input.txt")))
}
