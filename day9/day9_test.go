package day9

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"advent-of-code-2015/utils"
)

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, 251, solvePart1(utils.ReadLines("day9", "day-9-input.txt")))
}

func TestSolvePart2(t *testing.T) {
	assert.Equal(t, 898, solvePart2(utils.ReadLines("day9", "day-9-input.txt")))
}

func TestRoutePlanner_Route_Shortest(t *testing.T) {
	rp := NewRoutePlanner([]string{
		"London to Dublin = 464",
		"London to Belfast = 518",
		"Dublin to Belfast = 141",
	}, true)
	assert.Equal(t, 605, rp.Route())
}

func TestRoutePlanner_Route_Longest(t *testing.T) {
	rp := NewRoutePlanner([]string{
		"London to Dublin = 464",
		"London to Belfast = 518",
		"Dublin to Belfast = 141",
	}, false)
	assert.Equal(t, 982, rp.Route())
}
