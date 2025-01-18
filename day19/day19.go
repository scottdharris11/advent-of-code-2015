package day19

import (
	"advent-of-code-2015/utils"
	"log"
	"strings"
	"time"
)

type Puzzle struct{}

func (Puzzle) Solve() {
	input := utils.ReadLines("day19", "day-19-input.txt")
	solvePart1(input)
	solvePart2(input)
}

func solvePart1(lines []string) int {
	start := time.Now().UnixMilli()
	input, replacements := parseInput(lines)
	nValues := make(map[string]bool)
	for i := 0; i < len(input); i++ {
		keys := []string{string(input[i])}
		if i+1 < len(input) {
			keys = append(keys, input[i:i+2])
		}
		for _, key := range keys {
			rs, ok := replacements[key]
			if !ok {
				continue
			}
			for _, r := range rs {
				s := input[:i] + r + input[i+len(key):]
				nValues[s] = true
			}
		}
	}
	end := time.Now().UnixMilli()
	distinct := len(nValues)
	log.Printf("Day 19, Part 1 (%dms): Distinct = %d", end-start, distinct)
	return distinct
}

func solvePart2(lines []string) int {
	start := time.Now().UnixMilli()
	goal, replacements := parseInput(lines)
	rr := make(map[string]string)
	for k, v := range replacements {
		for _, s := range v {
			rr[s] = k
		}
	}
	visited := make(map[string]int)
	steps := reduce(goal, "e", 0, -1, rr, visited)
	end := time.Now().UnixMilli()
	log.Printf("Day 19, Part 2 (%dms): Steps = %d", end-start, steps)
	return steps
}

func parseInput(lines []string) (string, map[string][]string) {
	replacements := make(map[string][]string)
	input := ""
	for i, line := range lines {
		if line == "" {
			input = lines[i+1]
			break
		}
		pieces := strings.Split(line, " => ")
		r := replacements[pieces[0]]
		if r == nil {
			r = []string{}
		}
		replacements[pieces[0]] = append(r, pieces[1])
	}
	return input, replacements
}

func reduce(input string, goal string, steps int, best int, rr map[string]string, v map[string]int) int {
	if input == goal {
		return steps
	}
	if p, ok := v[input]; ok {
		return p
	}
	if best > 0 && steps <= best-1 {
		return -1
	}
	l := len(input)
	for i := 0; i < l; i++ {
		for e, r := range rr {
			ei := i + len(e)
			if ei <= l && input[i:ei] == e {
				ns := input[:i] + r + input[ei:]
				s := reduce(ns, goal, steps+1, best, rr, v)
				v[ns] = s
				if s > 0 && (best == -1 || s < best) {
					best = s
				}
			}
		}
	}
	return best
}
