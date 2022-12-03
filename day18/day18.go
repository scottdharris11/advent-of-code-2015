package day18

import (
	"advent-of-code-2015/utils"
	"log"
	"time"
)

type Puzzle struct{}

func (Puzzle) Solve() {
	input := utils.ReadLines("day18", "day-18-input.txt")
	solvePart1(input, 100)
	solvePart2(input, 100)
}

func solvePart1(lines []string, steps int) int {
	start := time.Now().UnixMilli()
	grid := Grid{}
	grid.initialize(lines)
	for i := 0; i < steps; i++ {
		grid.transition()
	}
	end := time.Now().UnixMilli()
	log.Printf("Day 18, Part 1 (%dms): Lights On = %d", end-start, grid.onCount)
	return grid.onCount
}

func solvePart2(lines []string, steps int) int {
	start := time.Now().UnixMilli()
	grid := Grid{}
	grid.cornersOn = true
	grid.initialize(lines)
	for i := 0; i < steps; i++ {
		grid.transition()
	}
	end := time.Now().UnixMilli()
	log.Printf("Day 18, Part 2 (%dms): Lights On = %d", end-start, grid.onCount)
	return grid.onCount
}

const ON rune = '#'

type Grid struct {
	cornersOn bool
	height    int
	width     int
	onCount   int
	lights    [][]bool
}

func (g *Grid) initialize(lines []string) {
	g.height = len(lines)
	g.width = len(lines[0])
	g.lights = make([][]bool, g.height)
	g.onCount = 0
	for rowIdx, line := range lines {
		row := make([]bool, g.width)
		for colIdx, light := range []rune(line) {
			row[colIdx] = light == ON || g.cornerOn(rowIdx, colIdx)
			if row[colIdx] {
				g.onCount++
			}
		}
		g.lights[rowIdx] = row
	}
}

func (g *Grid) transition() {
	nextLights := make([][]bool, g.height)
	nextOnCount := 0
	for rowIdx := 0; rowIdx < g.height; rowIdx++ {
		row := make([]bool, g.width)
		for colIdx := 0; colIdx < g.width; colIdx++ {
			row[colIdx] = g.cornerOn(rowIdx, colIdx) || g.onNext(rowIdx, colIdx)
			if row[colIdx] {
				nextOnCount++
			}
		}
		nextLights[rowIdx] = row
	}
	g.onCount = nextOnCount
	g.lights = nextLights
}

func (g *Grid) onNext(row int, col int) bool {
	neighborsOn := 0
	for i := row - 1; i <= row+1; i++ {
		for j := col - 1; j <= col+1; j++ {
			if i == row && j == col {
				continue
			}
			if g.on(i, j) {
				neighborsOn++
			}
		}
	}

	onNow := g.lights[row][col]
	switch {
	case onNow && (neighborsOn == 2 || neighborsOn == 3):
		return true
	case !onNow && neighborsOn == 3:
		return true
	}
	return false
}

func (g *Grid) on(row int, col int) bool {
	if row < 0 || row >= g.height || col < 0 || col >= g.width {
		return false
	}
	return g.lights[row][col]
}

func (g *Grid) cornerOn(row int, col int) bool {
	switch {
	case !g.cornersOn:
		return false
	case row == 0 && col == 0,
		row == 0 && col == g.width-1,
		row == g.height-1 && col == 0,
		row == g.height-1 && col == g.width-1:
		return true
	}
	return false
}
