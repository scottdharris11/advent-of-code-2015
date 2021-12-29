package day11

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"advent-of-code-2015/utils"
)

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, "hepxxyzz", solvePart1(utils.ReadLines("day11", "day-11-input.txt")[0]))
}

func TestSolvePart2(t *testing.T) {
	assert.Equal(t, "heqaabcc", solvePart2(utils.ReadLines("day11", "day-11-input.txt")[0]))
}

func TestNextPassword(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output string
	}{
		{"1", "abcdefgh", "abcdffaa"},
		{"2", "ghijklmn", "ghjaabcc"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.output, NextPassword(tt.input))
		})
	}
}

func TestValidPassword(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output bool
	}{
		{"1", "hijklmmn", false},
		{"2", "abbceffg", false},
		{"3", "abbceffg", false},
		{"4", "abcdffaa", true},
		{"5", "ghjaabcc", true},
		{"6", "abcoffaa", false},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.output, ValidPassword(tt.input))
		})
	}
}
