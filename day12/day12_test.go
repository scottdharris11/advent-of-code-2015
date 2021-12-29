package day12

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"advent-of-code-2015/utils"
)

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, 156366, solvePart1(utils.ReadLines("day12", "day-12-input.txt")))
}

func TestSolvePart2(t *testing.T) {
	assert.Equal(t, 96852, solvePart2(utils.ReadLines("day12", "day-12-input.txt")))
}

func TestSumNumbersIn(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		ignoreRed bool
		output    int
	}{
		{"1", "[1,2,3]", false, 6},
		{"2", "{\"a\":2,\"b\":4}", false, 6},
		{"3", "[[[3]]]", false, 3},
		{"4", "{\"a\":{\"b\":4},\"c\":-1}", false, 3},
		{"5", "{\"a\":[-1,1]}", false, 0},
		{"6", "[-1,{\"a\":1}]", false, 0},
		{"7", "[]", false, 0},
		{"8", "{}", false, 0},
		{"9", "[1,2,3]", true, 6},
		{"10", "[1,{\"c\":\"red\",\"b\":2},3]", false, 6},
		{"11", "[1,{\"c\":\"red\",\"b\":2},3]", true, 4},
		{"12", "{\"d\":\"red\",\"e\":[1,2,3,4],\"f\":5}", false, 15},
		{"13", "{\"d\":\"red\",\"e\":[1,2,3,4],\"f\":5}", true, 0},
		{"14", "[1,\"red\",5]", false, 6},
		{"15", "[1,\"red\",5]", true, 6},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			in := utils.ParseJSON(tt.input)
			assert.Equal(t, tt.output, SumNumbersIn(in, tt.ignoreRed))
		})
	}
}
