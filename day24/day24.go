package day24

import (
	"advent-of-code-2015/utils"
	"log"
	"math"
	"time"
)

type Puzzle struct{}

func (Puzzle) Solve() {
	input := utils.ReadIntegers("day24", "day-24-input.txt")
	solvePart1(input)
	solvePart2(input)
}

func solvePart1(weights []int) int {
	start := time.Now().UnixMilli()
	best := balanced(NewConfiguration(weights), nil)
	end := time.Now().UnixMilli()
	log.Printf("Day 24, Part 1 (%dms): Result = %d", end-start, best.quantum)
	return best.quantum
}

func solvePart2(weights []int) int {
	start := time.Now().UnixMilli()
	result := 0
	end := time.Now().UnixMilli()
	log.Printf("Day 24, Part 2 (%dms): Result = %d", end-start, result)
	return result
}

func balanced(c *Configuration, best *Configuration) *Configuration {
	if c.boxesRemaining == 0 {
		if c.balance == 0 && c.better(best) {
			best = c.copy()
		}
		return best
	}
	if !c.stillPossible(best) {
		return best
	}
	box := c.boxesRemaining - 1
	for container := 1; container < 4; container++ {
		c.boxToContainer(box, container)
		best = balanced(c, best)
		c.removeLastBoxFromContainer(container)
	}
	return best
}

func NewConfiguration(boxes []int) *Configuration {
	c := &Configuration{boxes: boxes, boxesRemaining: len(boxes), quantum: 1}
	for _, b := range boxes {
		c.weightToDistribute += b
	}
	return c
}

type Configuration struct {
	boxes              []int
	weightToDistribute int
	boxesRemaining     int
	c1                 []int
	c1Weight           int
	c2                 []int
	c2Weight           int
	c3                 []int
	c3Weight           int
	balance            int
	quantum            int
}

func (c *Configuration) stillPossible(best *Configuration) bool {
	if c.weightToDistribute < c.balance {
		return false
	}
	if best == nil {
		return true
	}
	if c.c1Weight > best.c1Weight || c.c2Weight > best.c2Weight || c.c3Weight > best.c3Weight {
		return false
	}
	if len(c.c1) > len(best.c1) {
		return false
	}
	if c.quantum > best.quantum {
		return false
	}
	return true
}

func (c *Configuration) better(best *Configuration) bool {
	if best == nil {
		return true
	}
	if len(c.c1) > len(best.c1) {
		return false
	}
	if len(c.c1) < len(best.c1) {
		return true
	}
	return c.quantum < best.quantum
}

func (c *Configuration) boxToContainer(box int, container int) {
	c.boxesRemaining--
	weight := c.boxes[box]
	c.weightToDistribute -= weight
	switch container {
	case 1:
		c.c1 = append(c.c1, weight)
		c.c1Weight += weight
		c.quantum *= weight
	case 2:
		c.c2 = append(c.c2, weight)
		c.c2Weight += weight
	case 3:
		c.c3 = append(c.c3, weight)
		c.c3Weight += weight
	}
	c.updateBalance()
}

func (c *Configuration) removeLastBoxFromContainer(container int) {
	c.boxesRemaining++
	var weight int
	switch container {
	case 1:
		idx := len(c.c1) - 1
		weight = c.c1[idx]
		c.c1 = c.c1[:idx]
		c.c1Weight -= weight
		c.quantum /= weight
	case 2:
		idx := len(c.c2) - 1
		weight = c.c2[idx]
		c.c2 = c.c2[:idx]
		c.c2Weight -= weight
	case 3:
		idx := len(c.c3) - 1
		weight = c.c3[idx]
		c.c3 = c.c3[:idx]
		c.c3Weight -= weight
	}
	c.weightToDistribute += weight
	c.updateBalance()
}

func (c *Configuration) updateBalance() {
	c.balance = int(math.Abs(float64(c.c1Weight-c.c2Weight)) + math.Abs(float64(c.c1Weight-c.c3Weight)))
}

func (c *Configuration) copy() *Configuration {
	c1 := make([]int, len(c.c1))
	copy(c1, c.c1)
	c2 := make([]int, len(c.c2))
	copy(c2, c.c2)
	c3 := make([]int, len(c.c3))
	copy(c3, c.c3)
	return &Configuration{
		boxes:              c.boxes,
		weightToDistribute: c.weightToDistribute,
		boxesRemaining:     c.boxesRemaining,
		c1:                 c1,
		c1Weight:           c.c1Weight,
		c2:                 c2,
		c2Weight:           c.c2Weight,
		c3:                 c3,
		c3Weight:           c.c3Weight,
		balance:            c.balance,
		quantum:            c.quantum,
	}
}
