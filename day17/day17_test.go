package day17

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"advent-of-code-2015/utils"
)

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, 4, solvePart1([]int{20, 15, 10, 5, 5}, 25))
	assert.Equal(t, 4372, solvePart1(utils.ReadIntegers("day17", "day-17-input.txt"), 150))
}

func TestSolvePart2(t *testing.T) {
	assert.Equal(t, 3, solvePart2([]int{20, 15, 10, 5, 5}, 25))
	assert.Equal(t, 4, solvePart2(utils.ReadIntegers("day17", "day-17-input.txt"), 150))
}

func TestCountCombinations(t *testing.T) {
	tests := []struct {
		containers     []int
		target         int
		prevUsed       int
		comboMap       map[int]int
		combos         int
		expectedCombos map[int]int
	}{
		{[]int{20, 15, 10, 5, 5}, 25, 0, nil, 4, nil},
		{[]int{15, 10, 5, 5}, 5, 1, make(map[int]int), 2, map[int]int{2: 2}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("Target %d with containers %v", tt.target, tt.containers), func(t *testing.T) {
			assert.Equal(t, tt.combos, CountCombinations(tt.containers, tt.target, tt.prevUsed, tt.comboMap))
			assert.Equal(t, tt.comboMap, tt.expectedCombos)
		})
	}
}
