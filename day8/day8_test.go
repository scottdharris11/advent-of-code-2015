package day8

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"advent-of-code-2015/utils"
)

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, 1371, solvePart1(utils.ReadLines("day8", "day-8-input.txt")))
}

func TestSolvePart2(t *testing.T) {
	assert.Equal(t, 2117, solvePart2(utils.ReadLines("day8", "day-8-input.txt")))
}

func TestCountCharacters(t *testing.T) {
	tests := []struct {
		name      string
		line      string
		codeCnt   int
		charCnt   int
		encodeCnt int
	}{
		{"1", "\"\"", 2, 0, 6},
		{"2", "\"abc\"", 5, 3, 9},
		{"3", "\"aaa\\\"aaa\"", 10, 7, 16},
		{"4", "\"\\x27\"", 6, 1, 11},
		{"5", "\"\\xa8br\\x8bjr\\\"\"", 16, 7, 24},
		{"6", "\"daz\\\\zyyxddpwk\"", 16, 13, 22},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			code, char, encoded := CountCharacters(tt.line)
			assert.Equal(t, tt.codeCnt, code)
			assert.Equal(t, tt.charCnt, char)
			assert.Equal(t, tt.encodeCnt, encoded)
		})
	}
}
