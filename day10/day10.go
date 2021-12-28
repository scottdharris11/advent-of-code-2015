package day10

import (
	"log"
	"strconv"
	"strings"
	"time"

	"advent-of-code-2015/utils"
)

type Puzzle struct{}

func (Puzzle) Solve() {
	input := utils.ReadLines("day10", "day-10-input.txt")
	solvePart1(input)
	solvePart2(input)
}

func solvePart1(lines []string) int {
	start := time.Now().UnixMilli()
	s := lines[0]
	for i := 0; i < 40; i++ {
		s = ExpandSequence(s)
	}
	ans := len(s)
	end := time.Now().UnixMilli()
	log.Printf("Day 10, Part 1 (%dms): Sequence Length = %d", end-start, ans)
	return ans
}

func solvePart2(lines []string) int {
	start := time.Now().UnixMilli()
	s := lines[0]
	for i := 0; i < 50; i++ {
		s = ExpandSequence(s)
	}
	ans := len(s)
	end := time.Now().UnixMilli()
	log.Printf("Day 10, Part 2 (%dms): Sequence Length = %d", end-start, ans)
	return ans
}

func ExpandSequence(s string) string {
	sb := strings.Builder{}
	current := '0'
	count := 0
	for _, r := range s {
		if current != '0' && r != current {
			sb.WriteString(strconv.Itoa(count))
			sb.WriteRune(current)
			count = 0
		}
		current = r
		count++
	}
	sb.WriteString(strconv.Itoa(count))
	sb.WriteRune(current)
	return sb.String()
}
