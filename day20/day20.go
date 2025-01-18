package day20

import (
	"log"
	"math"
	"time"
)

type Puzzle struct{}

func (Puzzle) Solve() {
	solvePart1(29000000)
	solvePart2(29000000)
}

func solvePart1(goal int) int {
	start := time.Now().UnixMilli()
	house := int(math.Sqrt(float64(goal / 10)))
	for {
		if presents(house) >= goal {
			break
		}
		house++
	}
	end := time.Now().UnixMilli()
	log.Printf("Day 20, Part 1 (%dms): House = %d", end-start, house)
	return house
}

func solvePart2(input int) int {
	start := time.Now().UnixMilli()
	result := 0
	end := time.Now().UnixMilli()
	log.Printf("Day 20, Part 2 (%dms): Results = %d", end-start, result)
	return result
}

func presents(house int) int {
	p := 0
	limit := int(math.Sqrt(float64(house))) + 1
	for n := 1; n < limit; n++ {
		if house%n == 0 {
			p += n * 10
			r := house / n
			if r != n {
				p += r * 10
			}
		}
	}
	return p
}
