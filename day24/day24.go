package day24

import (
	"advent-of-code-2015/utils"
	"log"
	"time"
)

type Puzzle struct{}

func (Puzzle) Solve() {
	input := utils.ReadIntegers("day24", "day-24-input.txt")
	solvePart1(input)
	solvePart2(input)
}

func solvePart1(weights []int) int {
	start := time.Now().UnixMilli()
	result := 0
	end := time.Now().UnixMilli()
	log.Printf("Day 24, Part 1 (%dms): Result = %d", end-start, result)
	return result
}

func solvePart2(weights []int) int {
	start := time.Now().UnixMilli()
	result := 0
	end := time.Now().UnixMilli()
	log.Printf("Day 24, Part 2 (%dms): Result = %d", end-start, result)
	return result
}
