package day4

import (
	"crypto/md5"
	"encoding/hex"
	"log"
	"strconv"
	"time"

	"advent-of-code-2015/utils"
)

type Puzzle struct{}

func (Puzzle) Solve() {
	input := utils.ReadLines("day4", "day-4-input.txt")
	solvePart1(input)
	solvePart2(input)
}

func solvePart1(lines []string) int64 {
	start := time.Now().UnixMilli()
	ans := LowestWithPrefix(lines[0], "00000", 0)
	end := time.Now().UnixMilli()
	log.Printf("Day 4, Part 1 (%dms): Number = %d", end-start, ans)
	return ans
}

func solvePart2(lines []string) int64 {
	start := time.Now().UnixMilli()
	ans := LowestWithPrefix(lines[0], "000000", 346386)
	end := time.Now().UnixMilli()
	log.Printf("Day 4, Part 2 (%dms): Number = %d", end-start, ans)
	return ans
}

func LowestWithPrefix(key string, pre string, start int) int64 {
	d := int64(start)
	prefix := []byte(pre)
	buffer := make([]byte, hex.EncodedLen(len(prefix)))
	keyBytes := []byte(key)
	for {
		if LeadingPrefix(keyBytes, d, prefix, buffer) {
			break
		}
		d++
	}
	return d
}

func LeadingPrefix(key []byte, d int64, prefix []byte, buffer []byte) bool {
	bytes := md5.Sum(strconv.AppendInt(key, d, 10))
	hex.Encode(buffer, bytes[:len(prefix)])
	for i, b := range prefix {
		if buffer[i] != b {
			return false
		}
	}
	return true
}
