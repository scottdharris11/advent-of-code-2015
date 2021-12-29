package day12

import (
	"log"
	"reflect"
	"time"

	"advent-of-code-2015/utils"
)

type Puzzle struct{}

func (Puzzle) Solve() {
	input := utils.ReadLines("day12", "day-12-input.txt")
	solvePart1(input)
	solvePart2(input)
}

func solvePart1(lines []string) int {
	start := time.Now().UnixMilli()
	ans := SumNumbersIn(utils.ParseJSON(lines[0]), false)
	end := time.Now().UnixMilli()
	log.Printf("Day 12, Part 1 (%dms): Number Sum = %d", end-start, ans)
	return ans
}

func solvePart2(lines []string) int {
	start := time.Now().UnixMilli()
	ans := SumNumbersIn(utils.ParseJSON(lines[0]), true)
	end := time.Now().UnixMilli()
	log.Printf("Day 12, Part 2 (%dms): Number Sum w/o Red = %d", end-start, ans)
	return ans
}

func SumNumbersIn(o interface{}, ignoreRed bool) int {
	sum := 0
	rt := reflect.TypeOf(o)
	switch rt.Kind() {
	case reflect.Slice, reflect.Array:
		s := o.([]interface{}) //nolint:errcheck
		for _, so := range s {
			sum += SumNumbersIn(so, ignoreRed)
		}
	case reflect.Map:
		m := o.(map[string]interface{}) //nolint:errcheck
		for _, mv := range m {
			if ignoreRed && reflect.TypeOf(mv).Kind() == reflect.String && mv.(string) == "red" {
				return 0
			}
			sum += SumNumbersIn(mv, ignoreRed)
		}
	case reflect.Int:
		return o.(int)
	case reflect.Float64:
		return int(o.(float64))
	}
	return sum
}
