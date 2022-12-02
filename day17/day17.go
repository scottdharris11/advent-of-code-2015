package day17

import (
	"log"
	"time"

	"advent-of-code-2015/utils"
)

type Puzzle struct{}

func (Puzzle) Solve() {
	input := utils.ReadIntegers("day17", "day-17-input.txt")
	solvePart1(input, 150)
	solvePart2(input, 150)
}

func solvePart1(containers []int, target int) int {
	start := time.Now().UnixMilli()
	combos := CountCombinations(containers, target, 0, nil)
	end := time.Now().UnixMilli()
	log.Printf("Day 17, Part 1 (%dms): Combinations = %d", end-start, combos)
	return combos
}

func solvePart2(containers []int, target int) int {
	start := time.Now().UnixMilli()
	combosByCnt := make(map[int]int)
	CountCombinations(containers, target, 0, combosByCnt)
	minCombos := 0
	i := 1
	for {
		if c, ok := combosByCnt[i]; ok {
			minCombos = c
			break
		}
		i++
	}
	end := time.Now().UnixMilli()
	log.Printf("Day 17, Part 2 (%dms): Min Container Combos = %d", end-start, minCombos)
	return minCombos
}

func CountCombinations(containers []int, goal int, prevUsed int, combosByCnt map[int]int) int {
	combos := 0
	for i, v := range containers {
		if v == goal {
			combos++
			if combosByCnt != nil {
				combosByCnt[prevUsed+1]++
			}
			continue
		}
		if v < goal {
			combos += CountCombinations(containers[i+1:], goal-v, prevUsed+1, combosByCnt)
		}
	}
	return combos
}
