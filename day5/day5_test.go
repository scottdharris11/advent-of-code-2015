package day5

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"advent-of-code-2015/utils"
)

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, 255, solvePart1(utils.ReadLines("day5", "day-5-input.txt")))
}

func TestSolvePart2(t *testing.T) {
	assert.Equal(t, 55, solvePart2(utils.ReadLines("day5", "day-5-input.txt")))
}

func TestNice1(t *testing.T) {
	tests := []struct {
		name     string
		line     string
		expected bool
	}{
		{"1", "ugknbfddgicrmopn", true},
		{"2", "aaa", true},
		{"3", "jchzalrnumimnmhp", false},
		{"4", "haegwjzuvuyypxyu", false},
		{"5", "dvszwmarrgswjxmb", false},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, Nice1(tt.line))
		})
	}
}

func TestNice2(t *testing.T) {
	tests := []struct {
		name     string
		line     string
		expected bool
	}{
		{"1", "qjhvhtzxzqqjkmpb", true},
		{"2", "xxyxx", true},
		{"3", "uurcxstgmygtbstg", false},
		{"4", "ieodomkazucvgmuy", false},
		{"5", "xyxyobo", true},
		{"6", "aabcdefgaaobo", true},
		{"7", "aaaobo", false},
		{"8", "aabcdefeghiaa", true},
		{"9", "xyyxobo", false},
		{"10", "xilodxfuxphuiiii", true},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, Nice2(tt.line))
		})
	}
}
