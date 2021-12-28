package day7

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"advent-of-code-2015/utils"
)

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, 3176, solvePart1(utils.ReadLines("day7", "day-7-input.txt")))
}

func TestSolvePart2(t *testing.T) {
	assert.Equal(t, 14710, solvePart2(utils.ReadLines("day7", "day-7-input.txt")))
}

func TestGate_ApplyInstructions(t *testing.T) {
	instructions := []string{
		"123 -> x",
		"456 -> y",
		"x AND y -> d",
		"x OR y -> e",
		"x LSHIFT 2 -> f",
		"y RSHIFT 2 -> g",
		"NOT x -> h",
		"NOT y -> i",
	}
	c := NewCircuit()
	c.ApplyInstructions(instructions)

	assert.Equal(t, uint16(72), *c.Wire("d").Value)
	assert.Equal(t, uint16(507), *c.Wire("e").Value)
	assert.Equal(t, uint16(492), *c.Wire("f").Value)
	assert.Equal(t, uint16(114), *c.Wire("g").Value)
	assert.Equal(t, uint16(65412), *c.Wire("h").Value)
	assert.Equal(t, uint16(65079), *c.Wire("i").Value)
	assert.Equal(t, uint16(123), *c.Wire("x").Value)
	assert.Equal(t, uint16(456), *c.Wire("y").Value)
}
