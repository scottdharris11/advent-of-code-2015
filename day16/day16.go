package day16

import (
	"log"
	"regexp"
	"time"

	"advent-of-code-2015/utils"
)

type Puzzle struct{}

func (Puzzle) Solve() {
	input := utils.ReadLines("day16", "day-16-input.txt")
	solvePart1(input)
	solvePart2(input)
}

func solvePart1(lines []string) int {
	start := time.Now().UnixMilli()
	sues := parseInput(lines)
	ans := FindMatchingSue(sues, true)
	end := time.Now().UnixMilli()
	log.Printf("Day 16, Part 1 (%dms): Sue = %d", end-start, ans)
	return ans
}

func solvePart2(lines []string) int {
	start := time.Now().UnixMilli()
	sues := parseInput(lines)
	ans := FindMatchingSue(sues, false)
	end := time.Now().UnixMilli()
	log.Printf("Day 16, Part 2 (%dms): Sue = %d", end-start, ans)
	return ans
}

func parseInput(lines []string) []Sue {
	matcher := regexp.MustCompile(`^Sue (.+): (.+): (\d+), (.+): (\d+), (.+): (\d+)$`)
	var sues []Sue
	for _, line := range lines {
		matches := matcher.FindStringSubmatch(line)
		sues = append(sues, Sue{
			id: utils.Number(matches[1]),
			attributes: map[string]int{
				matches[2]: utils.Number(matches[3]),
				matches[4]: utils.Number(matches[5]),
				matches[6]: utils.Number(matches[7]),
			},
		})
	}
	return sues
}

type Sue struct {
	id         int
	attributes map[string]int
}

var mfcsam = map[string]int{
	"children":    3,
	"cats":        7,
	"samoyeds":    2,
	"pomeranians": 3,
	"akitas":      0,
	"vizslas":     0,
	"goldfish":    5,
	"trees":       3,
	"cars":        2,
	"perfumes":    1,
}

func FindMatchingSue(sues []Sue, exact bool) int {
	for _, sue := range sues {
		if MatchSue(sue, exact) {
			return sue.id
		}
	}
	return -1
}

func MatchSue(sue Sue, exact bool) bool {
	for k, v := range sue.attributes {
		if exact {
			if mfcsam[k] != v {
				return false
			}
			continue
		}

		switch k {
		case "cats", "trees":
			if v <= mfcsam[k] {
				return false
			}
		case "pomeranians", "goldfish":
			if v >= mfcsam[k] {
				return false
			}
		default:
			if mfcsam[k] != v {
				return false
			}
		}
	}
	return true
}
