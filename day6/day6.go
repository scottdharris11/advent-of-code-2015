package day6

import (
	"log"
	"strings"
	"time"

	"advent-of-code-2015/utils"
)

type Puzzle struct{}

func (Puzzle) Solve() {
	input := utils.ReadLines("day6", "day-6-input.txt")
	solvePart1(input)
	solvePart2(input)
}

func solvePart1(lines []string) int {
	houseLights := HouseLights{}
	start := time.Now().UnixMilli()
	for _, line := range lines {
		houseLights.ProcessInstruction(line)
	}
	ans := houseLights.On
	end := time.Now().UnixMilli()
	log.Printf("Day 6, Part 1 (%dms): Lights On = %d", end-start, ans)
	return ans
}

func solvePart2(lines []string) int {
	houseLights := HouseLights{}
	start := time.Now().UnixMilli()
	for _, line := range lines {
		houseLights.ProcessInstruction(line)
	}
	ans := houseLights.Brightness
	end := time.Now().UnixMilli()
	log.Printf("Day 6, Part 2 (%dms): Light Brightness = %d", end-start, ans)
	return ans
}

type HouseLights struct {
	On             int
	Brightness     int
	gridOn         [1000][1000]bool
	gridBrightness [1000][1000]int
}

func (h *HouseLights) ProcessInstruction(i string) {
	pieces := strings.Split(i, " ")
	brCoordinatePieces := strings.Split(pieces[len(pieces)-1], ",")
	tlCoordinatePieces := strings.Split(pieces[len(pieces)-3], ",")
	top := utils.Number(tlCoordinatePieces[0])
	left := utils.Number(tlCoordinatePieces[1])
	bottom := utils.Number(brCoordinatePieces[0])
	right := utils.Number(brCoordinatePieces[1])

	for x := left; x <= right; x++ {
		for y := top; y <= bottom; y++ {
			if pieces[0] == "turn" {
				if pieces[1] == "on" {
					h.turnOn(x, y)
				} else {
					h.turnOff(x, y)
				}
			} else {
				h.toggle(x, y)
			}
		}
	}
}

func (h *HouseLights) turnOn(x int, y int) {
	if !h.gridOn[y][x] {
		h.On++
		h.gridOn[y][x] = true
	}
	h.gridBrightness[y][x]++
	h.Brightness++
}

func (h *HouseLights) turnOff(x int, y int) {
	if h.gridOn[y][x] {
		h.On--
		h.gridOn[y][x] = false
	}
	if h.gridBrightness[y][x] > 0 {
		h.gridBrightness[y][x]--
		h.Brightness--
	}
}

func (h *HouseLights) toggle(x int, y int) {
	h.gridOn[y][x] = !h.gridOn[y][x]
	if h.gridOn[y][x] {
		h.On++
	} else {
		h.On--
	}
	h.gridBrightness[y][x] += 2
	h.Brightness += 2
}
