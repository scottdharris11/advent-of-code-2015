package day20

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, 8, solvePart1(150))
	assert.Equal(t, 8, solvePart1(130))
	assert.Equal(t, 8, solvePart1(150))
	assert.Equal(t, 665280, solvePart1(29000000))
}

func TestSolvePart2(t *testing.T) {
	assert.Equal(t, 0, solvePart2(150))
	assert.Equal(t, 0, solvePart2(29000000))
}

func TestPresents(t *testing.T) {
	tests := []struct {
		input  int
		output int
	}{
		{1, 10},
		{2, 30},
		{3, 40},
		{4, 70},
		{5, 60},
		{6, 120},
		{7, 80},
		{8, 150},
		{9, 130},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("%d", tt.input), func(t *testing.T) {
			assert.Equal(t, tt.output, presents(tt.input))
		})
	}
}
