package day10

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"advent-of-code-2015/utils"
)

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, 252594, solvePart1(utils.ReadLines("day10", "day-10-input.txt")))
}

func TestSolvePart2(t *testing.T) {
	assert.Equal(t, 3579328, solvePart2(utils.ReadLines("day10", "day-10-input.txt")))
}

func TestExpandSequence(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output string
	}{
		{"1", "1", "11"},
		{"2", "11", "21"},
		{"3", "21", "1211"},
		{"4", "1211", "111221"},
		{"5", "111221", "312211"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.output, ExpandSequence(tt.input))
		})
	}
}
