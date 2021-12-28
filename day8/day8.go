package day8

import (
	"log"
	"time"

	"advent-of-code-2015/utils"
)

type Puzzle struct{}

func (Puzzle) Solve() {
	input := utils.ReadLines("day8", "day-8-input.txt")
	solvePart1(input)
	solvePart2(input)
}

func solvePart1(lines []string) int {
	start := time.Now().UnixMilli()
	codeCnt, charCnt := 0, 0
	for _, line := range lines {
		code, char, _ := CountCharacters(line)
		codeCnt += code
		charCnt += char
	}
	ans := codeCnt - charCnt
	end := time.Now().UnixMilli()
	log.Printf("Day 8, Part 1 (%dms): Char Difference = %d", end-start, ans)
	return ans
}

func solvePart2(lines []string) int {
	start := time.Now().UnixMilli()
	codeCnt, encodeCnt := 0, 0
	for _, line := range lines {
		code, _, encode := CountCharacters(line)
		codeCnt += code
		encodeCnt += encode
	}
	ans := encodeCnt - codeCnt
	end := time.Now().UnixMilli()
	log.Printf("Day 8, Part 2 (%dms): Char Difference = %d", end-start, ans)
	return ans
}

func CountCharacters(s string) (codeCnt int, charCnt int, encodedCnt int) {
	escaped := false
	hex := false
	hexEscapeCnt := 0
	for _, r := range s {
		codeCnt++
		encodedCnt++
		if hex {
			if (r >= '0' && r <= '9') || (r >= 'A' && r <= 'F') || (r >= 'a' && r <= 'f') {
				hexEscapeCnt++
				if hexEscapeCnt == 2 {
					escaped = false
					hex = false
					hexEscapeCnt = 0
					charCnt++
				}
			} else {
				escaped = false
				hex = false
				charCnt += 2 + hexEscapeCnt
			}
			continue
		}

		if escaped {
			switch r {
			case '\\':
				charCnt++
				encodedCnt++
			case '"':
				charCnt++
				encodedCnt++
			case 'x':
				hex = true
			default:
				charCnt += 2
			}
			escaped = false
			continue
		}

		if r == '\\' {
			escaped = true
			encodedCnt++
			continue
		}

		if r == '"' {
			encodedCnt += 2
			continue
		}

		charCnt++
	}
	return codeCnt, charCnt, encodedCnt
}
