package day2

import (
	"log"
	"regexp"
	"sort"
	"time"

	"advent-of-code-2015/utils"
)

type Puzzle struct{}

func (Puzzle) Solve() {
	input := utils.ReadLines("day2", "day-2-input.txt")
	solvePart1(input)
	solvePart2(input)
}

func solvePart1(lines []string) int {
	start := time.Now().UnixMilli()
	ans := 0
	for _, line := range lines {
		ans += PaperRequired(line)
	}
	end := time.Now().UnixMilli()
	log.Printf("Day 2, Part 1 (%dms): Paper Required = %d", end-start, ans)
	return ans
}

func solvePart2(lines []string) int {
	start := time.Now().UnixMilli()
	ans := 0
	for _, line := range lines {
		ans += RibbonRequired(line)
	}
	end := time.Now().UnixMilli()
	log.Printf("Day 2, Part 2 (%dms): Ribbon Required = %d", end-start, ans)
	return ans
}

func PaperRequired(s string) int {
	d := dimensions(s)
	return (2 * d[0] * d[1]) + (2 * d[1] * d[2]) + (2 * d[0] * d[2]) + (d[0] * d[1])
}

func RibbonRequired(s string) int {
	d := dimensions(s)
	rp := (2 * d[0]) + (2 * d[1])
	rb := d[0] * d[1] * d[2]
	return rp + rb
}

func dimensions(s string) []int {
	matcher := regexp.MustCompile(`^(.+)x(.+)x(.+)$`)
	matches := matcher.FindStringSubmatch(s)
	d := []int{utils.Number(matches[1]), utils.Number(matches[2]), utils.Number(matches[3])}
	sort.Ints(d)
	return d
}
