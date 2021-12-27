package day1

import (
	"log"
	"time"

	"advent-of-code-2015/utils"
)

type Puzzle struct{}

func (Puzzle) Solve() {
	input := utils.ReadLines("day1", "day-1-input.txt")
	solvePart1(input)
	solvePart2(input)
}

func solvePart1(lines []string) int {
	start := time.Now().UnixMilli()
	ans := FindFloor(lines[0])
	end := time.Now().UnixMilli()
	log.Printf("Day 1, Part 1 (%dms): Floor = %d", end-start, ans)
	return ans
}

func solvePart2(lines []string) int {
	start := time.Now().UnixMilli()
	ans := BasementIndex(lines[0])
	end := time.Now().UnixMilli()
	log.Printf("Day 1, Part 2 (%dms): Basement Index = %d", end-start, ans)
	return ans
}

func FindFloor(line string) int {
	floor := 0
	for _, r := range line {
		if r == '(' {
			floor++
		}
		if r == ')' {
			floor--
		}
	}
	return floor
}

func BasementIndex(line string) int {
	floor := 0
	for i, r := range line {
		if r == '(' {
			floor++
		}
		if r == ')' {
			floor--
		}
		if floor < 0 {
			return i + 1
		}
	}
	return -1
}
