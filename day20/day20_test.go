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
	assert.Equal(t, 84, solvePart2(2376))
	assert.Equal(t, 705600, solvePart2(29000000))
}

func TestPresents(t *testing.T) {
	tests := []struct {
		input  int
		output int
		per    int
		limit  int
	}{
		{1, 10, 10, 0},
		{2, 30, 10, 0},
		{3, 40, 10, 0},
		{4, 70, 10, 0},
		{5, 60, 10, 0},
		{6, 120, 10, 0},
		{7, 80, 10, 0},
		{8, 150, 10, 0},
		{9, 130, 10, 0},
		// 84x1 == yes == 924
		// 1x84 == no  == 0
		// 42x2 == yes == 462
		// 2x42 == yes == 22
		// 28x3 == yes == 308
		// 3x28 == yes == 33
		// 21x4 == yes == 231
		// 4*21 == yes == 44
		// 14*6 == yes == 154
		// 6*14 == yes == 66
		// 12*7 == yes == 132
		// 7*12 == yes == 77
		{84, 2453, 11, 50},
		// 100x1 == yes == 1100
		// 1x100 == no  == 0
		// 2x50  == yes == 22
		// 50x2  == yes == 550
		// 4x25  == yes == 44
		// 25x4  == yes == 275
		// 5*20  == yes == 55
		// 20*5  == yes == 220
		// 10x10 == yes == 110
		{100, 2376, 11, 50},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("%d", tt.input), func(t *testing.T) {
			assert.Equal(t, tt.output, presents(tt.input, tt.per, tt.limit))
		})
	}
}
