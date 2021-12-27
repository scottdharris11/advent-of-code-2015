package day2

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"advent-of-code-2015/utils"
)

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, 1598415, solvePart1(utils.ReadLines("day2", "day-2-input.txt")))
}

func TestSolvePart2(t *testing.T) {
	assert.Equal(t, 3812909, solvePart2(utils.ReadLines("day2", "day-2-input.txt")))
}

func TestPaperRequired(t *testing.T) {
	tests := []struct {
		name     string
		line     string
		expected int
	}{
		{"1", "2x3x4", 58},
		{"2", "1x1x10", 43},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, PaperRequired(tt.line))
		})
	}
}

func TestRibbonRequired(t *testing.T) {
	tests := []struct {
		name     string
		line     string
		expected int
	}{
		{"1", "2x3x4", 34},
		{"2", "1x1x10", 14},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, RibbonRequired(tt.line))
		})
	}
}
