package day25

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, 9380097, solvePart1(4, 4))
	assert.Equal(t, 27995004, solvePart1(6, 6))
	assert.Equal(t, 9250759, solvePart1(5, 5))
	assert.Equal(t, 2650453, solvePart1(2978, 3083))
}

func TestCodeNumber(t *testing.T) {
	tests := []struct {
		row      int
		col      int
		expected int
	}{
		{1, 1, 1},
		{2, 1, 2},
		{3, 2, 8},
		{4, 3, 18},
		{6, 1, 16},
		{1, 5, 15},
		{2, 5, 20},
		{1, 6, 21},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("%d/%d", tt.row, tt.col), func(t *testing.T) {
			assert.Equal(t, tt.expected, codeNumber(tt.row, tt.col))
		})
	}
}

func TestCode(t *testing.T) {
	tests := []struct {
		codeNumber int
		expected   int
	}{
		{1, 20151125},
		{2, 31916031},
		{6, 17289845},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("%d", tt.codeNumber), func(t *testing.T) {
			assert.Equal(t, tt.expected, code(tt.codeNumber))
		})
	}
}

func TestNextCode(t *testing.T) {
	tests := []struct {
		prev     int
		expected int
	}{
		{20151125, 31916031},
		{31916031, 18749137},
		{18749137, 16080970},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("%d", tt.prev), func(t *testing.T) {
			assert.Equal(t, tt.expected, nextCode(tt.prev))
		})
	}
}
