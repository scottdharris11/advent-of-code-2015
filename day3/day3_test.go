package day3

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"advent-of-code-2015/utils"
)

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, 2592, solvePart1(utils.ReadLines("day3", "day-3-input.txt")))
}

func TestSolvePart2(t *testing.T) {
	assert.Equal(t, 2360, solvePart2(utils.ReadLines("day3", "day-3-input.txt")))
}

func TestCountHouses(t *testing.T) {
	tests := []struct {
		name     string
		line     string
		expected int
	}{
		{"1", ">", 2},
		{"2", "^>v<", 4},
		{"3", "^v^v^v^v^v", 2},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, CountHouses(tt.line))
		})
	}
}

func TestCountHousesWithRobo(t *testing.T) {
	tests := []struct {
		name     string
		line     string
		expected int
	}{
		{"1", "^v", 3},
		{"2", "^>v<", 3},
		{"3", "^v^v^v^v^v", 11},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, CountHousesWithRobo(tt.line))
		})
	}
}
