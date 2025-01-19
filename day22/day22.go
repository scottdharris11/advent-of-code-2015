package day22

import (
	"log"
	"time"
)

type Puzzle struct{}

var Boss = Player{HitPoints: 58, Damage: 9}
var Me = Player{HitPoints: 50, Mana: 500}

func (Puzzle) Solve() {
	solvePart1(Boss, Me)
	solvePart2(Boss, Me)
}

func solvePart1(boss Player, player Player) int {
	start := time.Now().UnixMilli()
	result := 0
	end := time.Now().UnixMilli()
	log.Printf("Day 21, Part 1 (%dms): Result = %d", end-start, result)
	return result
}

func solvePart2(boss Player, player Player) int {
	start := time.Now().UnixMilli()
	result := 0
	end := time.Now().UnixMilli()
	log.Printf("Day 21, Part 2 (%dms): Result = %d", end-start, result)
	return result
}

type Player struct {
	HitPoints int
	Damage    int
	Armor     int
	Mana      int
}

func (a Player) damageTo(b Player) int {
	d := a.Damage - b.Armor
	if d < 1 {
		d = 1
	}
	return d
}
