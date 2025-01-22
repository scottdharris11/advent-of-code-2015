package day25

import (
	"log"
	"time"
)

type Puzzle struct{}

func (Puzzle) Solve() {
	solvePart1(2978, 3083)
}

func solvePart1(row int, col int) int {
	start := time.Now().UnixMilli()
	code := code(codeNumber(row, col))
	end := time.Now().UnixMilli()
	log.Printf("Day 25, Part 1 (%dms): Code = %d", end-start, code)
	return code
}

func codeNumber(row int, col int) int {
	n := 1
	for i := 2; i < col+1; i++ {
		n += i
	}
	for i := 2; i < row+1; i++ {
		n += i + col - 2
	}
	return n
}

func code(codeNumber int) int {
	c := 20151125
	for i := 1; i < codeNumber; i++ {
		c = nextCode(c)
	}
	return c
}

func nextCode(prev int) int {
	c := prev * 252533
	return c % 33554393
}
