package day18

import (
	"advent-of-code-2015/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testLines = []string{
	".#.#.#",
	"...##.",
	"#....#",
	"..#...",
	"#.#..#",
	"####..",
}

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, 4, solvePart1(testLines, 4))
	assert.Equal(t, 768, solvePart1(utils.ReadLines("day18", "day-18-input.txt"), 100))
}

func TestSolvePart2(t *testing.T) {
	assert.Equal(t, 17, solvePart2(testLines, 5))
	assert.Equal(t, 781, solvePart2(utils.ReadLines("day18", "day-18-input.txt"), 100))
}

func TestGridInitialize(t *testing.T) {
	grid := Grid{}
	grid.initialize([]string{
		"#..#",
		".##.",
		"#..#",
	})

	assert.Equal(t, 3, grid.height)
	assert.Equal(t, 4, grid.width)
	assert.Equal(t, 6, grid.onCount)
	assert.Equal(t, [][]bool{
		{true, false, false, true},
		{false, true, true, false},
		{true, false, false, true},
	}, grid.lights)
}

func TestTransition(t *testing.T) {
	grid := Grid{}
	grid.height = 3
	grid.width = 4
	grid.lights = [][]bool{
		{true, false, false, true},
		{false, true, true, false},
		{true, false, false, true},
	}

	grid.transition()

	assert.Equal(t, 3, grid.height)
	assert.Equal(t, 4, grid.width)
	assert.Equal(t, 8, grid.onCount)
	assert.Equal(t, [][]bool{
		{false, true, true, false},
		{true, true, true, true},
		{false, true, true, false},
	}, grid.lights)
}

func TestGridOnNext(t *testing.T) {
	grid := Grid{}
	grid.height = 3
	grid.width = 4
	grid.lights = [][]bool{
		{true, false, false, true},
		{false, true, true, false},
		{true, false, false, true},
	}

	assert.False(t, grid.onNext(0, 0))
	assert.True(t, grid.onNext(0, 1))
	assert.True(t, grid.onNext(0, 2))
	assert.False(t, grid.onNext(0, 3))

	assert.True(t, grid.onNext(1, 0))
	assert.True(t, grid.onNext(1, 1))
	assert.True(t, grid.onNext(1, 2))
	assert.True(t, grid.onNext(1, 3))

	assert.False(t, grid.onNext(2, 0))
	assert.True(t, grid.onNext(2, 1))
	assert.True(t, grid.onNext(2, 2))
	assert.False(t, grid.onNext(2, 3))
}

func TestGridOn(t *testing.T) {
	grid := Grid{}
	grid.height = 3
	grid.width = 4
	grid.lights = [][]bool{
		{true, false, false, true},
		{false, true, true, false},
		{true, false, false, true},
	}

	assert.True(t, grid.on(0, 0))
	assert.False(t, grid.on(0, 1))
	assert.False(t, grid.on(-1, 0))
	assert.False(t, grid.on(0, -1))
	assert.False(t, grid.on(grid.height, 0))
	assert.False(t, grid.on(0, grid.width))
}

func TestGridCornerOn(t *testing.T) {
	grid := Grid{}
	grid.height = 3
	grid.width = 4
	grid.cornersOn = false

	assert.False(t, grid.cornerOn(0, 0))
	assert.False(t, grid.cornerOn(0, grid.width-1))
	assert.False(t, grid.cornerOn(grid.height-1, 0))
	assert.False(t, grid.cornerOn(grid.height-1, grid.width-1))

	grid.cornersOn = true

	assert.True(t, grid.cornerOn(0, 0))
	assert.True(t, grid.cornerOn(0, grid.width-1))
	assert.True(t, grid.cornerOn(grid.height-1, 0))
	assert.True(t, grid.cornerOn(grid.height-1, grid.width-1))

	assert.False(t, grid.cornerOn(0, 1))
	assert.False(t, grid.cornerOn(1, 0))
}
