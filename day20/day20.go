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
		if presents(house, 10, 0) >= goal {
			break
		}
		house++
	}
	end := time.Now().UnixMilli()
	log.Printf("Day 20, Part 1 (%dms): House = %d", end-start, house)
	return house
}

func solvePart2(goal int) int {
	start := time.Now().UnixMilli()
	house := int(math.Sqrt(float64(goal / 11)))
	for {
		if presents(house, 11, 50) >= goal {
			break
		}
		house++
	}
	end := time.Now().UnixMilli()
	log.Printf("Day 20, Part 2 (%dms): House = %d", end-start, house)
	return house
}

func presents(house int, per int, limit int) int {
	p := 0
	max := int(math.Sqrt(float64(house))) + 1
	for n := 1; n < max; n++ {
		if house%n == 0 {
			if limit == 0 || limit*n >= house {
				p += n * per
			}
			r := house / n
			if r != n {
				if limit == 0 || limit*r >= house {
					p += r * per
				}
			}
		}
	}
	return p
}
