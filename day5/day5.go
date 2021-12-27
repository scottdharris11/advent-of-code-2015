package day5

import (
	"log"
	"time"

	"advent-of-code-2015/utils"
)

type Puzzle struct{}

func (Puzzle) Solve() {
	input := utils.ReadLines("day5", "day-5-input.txt")
	solvePart1(input)
	solvePart2(input)
}

func solvePart1(lines []string) int {
	start := time.Now().UnixMilli()
	ans := 0
	for _, line := range lines {
		if Nice1(line) {
			ans++
		}
	}
	end := time.Now().UnixMilli()
	log.Printf("Day 5, Part 1 (%dms): Nice1 Strings = %d", end-start, ans)
	return ans
}

func solvePart2(lines []string) int {
	start := time.Now().UnixMilli()
	ans := 0
	for _, line := range lines {
		if Nice2(line) {
			ans++
		}
	}
	end := time.Now().UnixMilli()
	log.Printf("Day 5, Part 2 (%dms): Nice2 Strings = %d", end-start, ans)
	return ans
}

var badCombos = map[rune]rune{'b': 'a', 'd': 'c', 'q': 'p', 'y': 'x'}

func Nice1(s string) bool {
	vowels := 0
	repeat := false
	var prevR rune
	for _, r := range s {
		switch r {
		case 'a', 'e', 'i', 'o', 'u':
			vowels++
		}
		if prevR == r {
			repeat = true
		}
		if badR, ok := badCombos[r]; ok && badR == prevR {
			return false
		}
		prevR = r
	}
	return repeat && vowels >= 3
}

func Nice2(s string) bool {
	// capture character combinations and indexes
	var prevR rune
	var prevR2 rune
	combos := make(map[string][]int, len(s))
	repeat := false
	for i, r := range s {
		if i > 0 {
			cs := string([]rune{prevR, r})
			if _, ok := combos[cs]; !ok {
				combos[cs] = make([]int, 0, 5)
			}
			combos[cs] = append(combos[cs], i)

			if i > 1 && r == prevR2 {
				repeat = true
			}
		}

		prevR2 = prevR
		prevR = r
	}

	// look for pairs that don't overlap
	pairs := false
PairCheck:
	for _, indexes := range combos {
		for i, idx1 := range indexes {
			for j, idx2 := range indexes {
				if i == j {
					continue
				}
				if idx2 > idx1+1 {
					pairs = true
					break PairCheck
				}
			}
		}
	}

	return repeat && pairs
}
