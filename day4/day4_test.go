package day4

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"

	"advent-of-code-2015/utils"
)

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, 346386, solvePart1(utils.ReadLines("day4", "day-4-input.txt")))
}

func TestSolvePart2(t *testing.T) {
	assert.Equal(t, 9958218, solvePart2(utils.ReadLines("day4", "day-4-input.txt")))
}

func TestLowestWithPrefix(t *testing.T) {
	tests := []struct {
		name     string
		key      string
		prefix   string
		expected int
	}{
		{"1", "abcdef", "00000", 609043},
		{"2", "pqrstuv", "00000", 1048970},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, LowestWithPrefix(tt.key, tt.prefix, 0))
		})
	}
}

func TestLeadingPrefix(t *testing.T) {
	tests := []struct {
		name     string
		key      string
		number   int
		prefix   string
		expected bool
	}{
		{"1", "abcdef", 609043, "00000", true},
		{"2", "abcdef", 609042, "00000", false},
		{"3", "pqrstuv", 1048970, "00000", true},
		{"4", "pqrstuv", 1048969, "00000", false},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			prefix := []byte(tt.prefix)
			buffer := make([]byte, hex.EncodedLen(len(prefix)))
			assert.Equal(t, tt.expected, LeadingPrefix(tt.key, tt.number, prefix, buffer))
		})
	}
}
