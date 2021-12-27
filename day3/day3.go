package day3

import (
	"log"
	"strconv"
	"time"

	"advent-of-code-2015/utils"
)

type Puzzle struct{}

func (Puzzle) Solve() {
	input := utils.ReadLines("day3", "day-3-input.txt")
	solvePart1(input)
	solvePart2(input)
}

func solvePart1(lines []string) int {
	start := time.Now().UnixMilli()
	ans := CountHouses(lines[0])
	end := time.Now().UnixMilli()
	log.Printf("Day 3, Part 1 (%dms): Houses = %d", end-start, ans)
	return ans
}

func solvePart2(lines []string) int {
	start := time.Now().UnixMilli()
	ans := CountHousesWithRobo(lines[0])
	end := time.Now().UnixMilli()
	log.Printf("Day 3, Part 2 (%dms): Houses = %d", end-start, ans)
	return ans
}

func CountHouses(s string) int {
	x, y := 0, 0
	locations := make(map[string]int, 100)
	locations["0-0"]++
	for _, r := range s {
		loc := move(r, &x, &y)
		locations[loc]++
	}
	return len(locations)
}

func CountHousesWithRobo(s string) int {
	sX, sY, rX, rY := 0, 0, 0, 0
	santa := true
	locations := make(map[string]int, 100)
	locations["0-0"]++
	for _, r := range s {
		var loc string
		if santa {
			loc = move(r, &sX, &sY)
		} else {
			loc = move(r, &rX, &rY)
		}
		locations[loc]++
		santa = !santa
	}
	return len(locations)
}

func move(d rune, currX *int, currY *int) string {
	switch d {
	case '>':
		*currX++
	case '<':
		*currX--
	case '^':
		*currY++
	case 'v':
		*currY--
	}
	return strconv.Itoa(*currX) + "-" + strconv.Itoa(*currY)
}
