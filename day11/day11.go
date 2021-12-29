package day11

import (
	"log"
	"time"

	"advent-of-code-2015/utils"
)

type Puzzle struct{}

func (Puzzle) Solve() {
	input := utils.ReadLines("day11", "day-11-input.txt")
	solvePart1(input[0])
	solvePart2(input[0])
}

func solvePart1(pwd string) string {
	start := time.Now().UnixMilli()
	ans := NextPassword(pwd)
	end := time.Now().UnixMilli()
	log.Printf("Day 11, Part 1 (%dms): Next Password = %s", end-start, ans)
	return ans
}

func solvePart2(pwd string) string {
	start := time.Now().UnixMilli()
	ans := NextPassword(pwd)
	ans = NextPassword(ans)
	end := time.Now().UnixMilli()
	log.Printf("Day 11, Part 2 (%dms): Next Password = %s", end-start, ans)
	return ans
}

func NextPassword(pwd string) string {
	// look for any invalid characters and if found, move all index
	// values after it to 'z' so that when incremented, will reset
	// to 'a' after adjusting the invalid character index
	work := []rune(pwd)
	length := len(work)
	for i, r := range work {
		if r == 'i' || r == 'o' || r == 'l' {
			for j := i + 1; j < length; j++ {
				work[j] = 'z'
			}
			break
		}
	}

	// start with end index and increase its value until we hit 'z',
	// then move up to the next index.  After each adjustment, determine
	// if valid
	idx := length - 1
	for {
		zAdjust := false
		tempIdx := idx
		for work[idx] == 'z' {
			work[idx] = 'a'
			idx--
			zAdjust = true
		}

		if idx < 0 {
			break
		}

		work[idx]++
		if work[idx] == 'i' || work[idx] == 'o' || work[idx] == 'l' {
			work[idx]++
		}

		if zAdjust {
			idx = tempIdx
		}

		s := string(work)
		if ValidPassword(s) {
			return s
		}
	}
	panic("something went really wrong...no next password found")
}

func ValidPassword(pwd string) bool {
	var prevChar rune
	var prevChar2 rune
	pairCnt := 0
	pairActive := false
	straightFound := false
	for _, r := range pwd {
		if r == 'i' || r == 'o' || r == 'l' {
			return false
		}

		if r == prevChar {
			if !pairActive {
				pairCnt++
				pairActive = true
			}
		} else {
			pairActive = false
		}

		if prevChar2+2 == r && prevChar+1 == r {
			straightFound = true
		}

		prevChar2 = prevChar
		prevChar = r
	}
	return pairCnt >= 2 && straightFound
}
