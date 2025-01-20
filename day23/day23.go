package day23

import (
	"advent-of-code-2015/utils"
	"log"
	"strconv"
	"strings"
	"time"
)

type Puzzle struct{}

func (Puzzle) Solve() {
	input := utils.ReadLines("day23", "day-23-input.txt")
	solvePart1(input, "b")
	solvePart2(input)
}

func solvePart1(lines []string, register string) int {
	start := time.Now().UnixMilli()
	instructions := parseInstructions(lines)
	registers := runProgram(instructions)
	result := registers[register]
	end := time.Now().UnixMilli()
	log.Printf("Day 23, Part 1 (%dms): Result = %d", end-start, result)
	return result
}

func solvePart2(lines []string) int {
	start := time.Now().UnixMilli()
	result := 0
	end := time.Now().UnixMilli()
	log.Printf("Day 23, Part 2 (%dms): Result = %d", end-start, result)
	return result
}

func parseInstructions(lines []string) []Instruction {
	var instructions []Instruction
	for _, line := range lines {
		pieces := strings.Split(line, " ")
		switch pieces[0] {
		case "hlf":
			instructions = append(instructions, Half{register: pieces[1]})
		case "tpl":
			instructions = append(instructions, Triple{register: pieces[1]})
		case "inc":
			instructions = append(instructions, Increment{register: pieces[1]})
		case "jmp":
			offset, err := strconv.Atoi(pieces[1])
			if err != nil {
				log.Fatal("Invalid jump parameter: " + line)
			}
			instructions = append(instructions, Jump{offset: offset})
		case "jie":
			offset, err := strconv.Atoi(pieces[2])
			if err != nil {
				log.Fatal("Invalid jump parameter: " + line)
			}
			register := strings.ReplaceAll(pieces[1], ",", "")
			instructions = append(instructions, JumpIfEven{register: register, offset: offset})
		case "jio":
			offset, err := strconv.Atoi(pieces[2])
			if err != nil {
				log.Fatal("Invalid jump parameter: " + line)
			}
			register := strings.ReplaceAll(pieces[1], ",", "")
			instructions = append(instructions, JumpIfOne{register: register, offset: offset})
		}
	}
	return instructions
}

func runProgram(instructions []Instruction) map[string]int {
	registers := map[string]int{"a": 0, "b": 0}
	idx := 0
	for {
		if idx < 0 || idx >= len(instructions) {
			break
		}
		idx = instructions[idx].Apply(registers, idx)
	}
	return registers
}

type Instruction interface {
	Apply(registers map[string]int, index int) int
}

type Half struct {
	register string
}

func (h Half) Apply(registers map[string]int, index int) int {
	before := registers[h.register]
	after := before / 2
	registers[h.register] = after
	return index + 1
}

type Triple struct {
	register string
}

func (t Triple) Apply(registers map[string]int, index int) int {
	before := registers[t.register]
	after := before * 3
	registers[t.register] = after
	return index + 1
}

type Increment struct {
	register string
}

func (i Increment) Apply(registers map[string]int, index int) int {
	before := registers[i.register]
	after := before + 1
	registers[i.register] = after
	return index + 1
}

type Jump struct {
	offset int
}

func (j Jump) Apply(_ map[string]int, index int) int {
	return index + j.offset
}

type JumpIfEven struct {
	register string
	offset   int
}

func (j JumpIfEven) Apply(registers map[string]int, index int) int {
	before := registers[j.register]
	if before%2 == 0 {
		return index + j.offset
	}
	return index + 1
}

type JumpIfOne struct {
	register string
	offset   int
}

func (j JumpIfOne) Apply(registers map[string]int, index int) int {
	before := registers[j.register]
	if before == 1 {
		return index + j.offset
	}
	return index + 1
}
