package main

import (
	"advent-of-code-2015/day1"
	"advent-of-code-2015/day2"
	"advent-of-code-2015/day3"
	"advent-of-code-2015/day4"
	"advent-of-code-2015/day5"
	"advent-of-code-2015/day6"
)

type Solver interface {
	Solve()
}

func main() {
	solvers := []Solver{
		day1.Puzzle{}, day2.Puzzle{}, day3.Puzzle{}, day4.Puzzle{}, day5.Puzzle{},
		day6.Puzzle{},
	}
	for _, solver := range solvers {
		solver.Solve()
	}
}
