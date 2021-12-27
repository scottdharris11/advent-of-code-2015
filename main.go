package main

import "advent-of-code-2015/day1"

type Solver interface {
	Solve()
}

func main() {
	solvers := []Solver{
		day1.Puzzle{},
	}
	for _, solver := range solvers {
		solver.Solve()
	}
}
