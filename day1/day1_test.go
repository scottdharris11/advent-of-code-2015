package day1

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"advent-of-code-2015/utils"
)

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, 138, solvePart1(utils.ReadLines("day1", "day-1-input.txt")))
}

func TestSolvePart2(t *testing.T) {
	assert.Equal(t, 1771, solvePart2(utils.ReadLines("day1", "day-1-input.txt")))
}

func TestFindFloor(t *testing.T) {
	tests := []struct {
		name     string
		line     string
		expected int
	}{
		{"1", "(())", 0},
		{"2", "()()", 0},
		{"3", "(((", 3},
		{"4", "(()(()(", 3},
		{"5", "))(((((", 3},
		{"6", "())", -1},
		{"7", "))(", -1},
		{"8", ")))", -3},
		{"9", ")())())", -3},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, FindFloor(tt.line))
		})
	}
}

func TestBasementIndex(t *testing.T) {
	tests := []struct {
		name     string
		line     string
		expected int
	}{
		{"1", "(())", -1},
		{"2", ")", 1},
		{"3", "()())", 5},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, BasementIndex(tt.line))
		})
	}
}
