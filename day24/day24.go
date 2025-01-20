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
	best := balanced(weights, NewConfiguration(weights, 3), nil)
	log.Printf("Best: %+v", *best)
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

func balanced(weights []int, c *Configuration, best *Configuration) *Configuration {
	if c.boxesRemaining == 0 {
		if c.balance == 0 && c.better(best) {
			best = c.copy()
		}
		return best
	}
	if !c.stillPossible(best) {
		return best
	}
	box := weights[c.boxesRemaining-1]
	for container := 0; container < 3; container++ {
		c.boxToContainer(box, container)
		best = balanced(weights, c, best)
		c.removeLastBoxFromContainer(container)
	}
	return best
}

func NewConfiguration(boxes []int, containerCount int) *Configuration {
	containers := make([][]int, containerCount)
	weights := make([]int, containerCount)
	c := &Configuration{boxesRemaining: len(boxes), containers: containers, weights: weights, quantum: 1}
	for _, b := range boxes {
		c.weightToDistribute += b
	}
	return c
}

type Configuration struct {
	weightToDistribute int
	boxesRemaining     int
	containers         [][]int
	weights            []int
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
	for i, weight := range c.weights {
		if weight > best.weights[i] {
			return false
		}
	}
	if len(c.containers[0]) > len(best.containers[0]) {
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
	if len(c.containers[0]) > len(best.containers[0]) {
		return false
	}
	if len(c.containers[0]) < len(best.containers[0]) {
		return true
	}
	return c.quantum < best.quantum
}

func (c *Configuration) boxToContainer(box int, container int) {
	c.boxesRemaining--
	c.weightToDistribute -= box
	c.containers[container] = append(c.containers[container], box)
	c.weights[container] += box
	if container == 0 {
		c.quantum *= box
	}
	c.updateBalance()
}

func (c *Configuration) removeLastBoxFromContainer(container int) {
	c.boxesRemaining++
	var weight int
	idx := len(c.containers[container]) - 1
	weight = c.containers[container][idx]
	c.containers[container] = c.containers[container][:idx]
	c.weights[container] -= weight
	if container == 0 {
		c.quantum /= weight
	}
	c.weightToDistribute += weight
	c.updateBalance()
}

func (c *Configuration) updateBalance() {
	c1Weight := c.weights[0]
	balance := 0
	for i := 1; i < len(c.weights); i++ {
		balance += int(math.Abs(float64(c1Weight - c.weights[i])))
	}
	c.balance = balance
}

func (c *Configuration) copy() *Configuration {
	cnt := len(c.containers)
	containers := make([][]int, cnt)
	weights := make([]int, cnt)
	for i := 0; i < cnt; i++ {
		containers[i] = make([]int, len(c.containers[i]))
		copy(containers[i], c.containers[i])
		weights[i] = c.weights[i]
	}
	return &Configuration{
		weightToDistribute: c.weightToDistribute,
		boxesRemaining:     c.boxesRemaining,
		containers:         containers,
		weights:            weights,
		balance:            c.balance,
		quantum:            c.quantum,
	}
}
