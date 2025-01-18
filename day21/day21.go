package day21

import (
	"log"
	"sort"
	"time"
)

type Puzzle struct{}

var Boss = Attacker{HitPoints: 100, Damage: 8, Armor: 2}

func (Puzzle) Solve() {
	solvePart1(Boss)
	solvePart2(Boss)
}

func solvePart1(boss Attacker) int {
	start := time.Now().UnixMilli()
	wins, _ := battles(boss)
	end := time.Now().UnixMilli()
	log.Printf("Day 21, Part 1 (%dms): Lowest Cost = %d", end-start, wins[0])
	return wins[0]
}

func solvePart2(boss Attacker) int {
	start := time.Now().UnixMilli()
	_, losses := battles(boss)
	end := time.Now().UnixMilli()
	log.Printf("Day 21, Part 2 (%dms): Highest Loss = %d", end-start, losses[len(losses)-1])
	return losses[len(losses)-1]
}

var Weapons = []Item{
	{8, 4, 0},
	{10, 5, 0},
	{25, 6, 0},
	{40, 7, 0},
	{74, 8, 0},
}

var Armor = []Item{
	{13, 0, 1},
	{31, 0, 2},
	{53, 0, 3},
	{75, 0, 4},
	{102, 0, 5},
}

var Rings = []Item{
	{25, 1, 0},
	{50, 2, 0},
	{100, 3, 0},
	{20, 0, 1},
	{40, 0, 2},
	{80, 0, 3},
}

type Item struct {
	Cost   int
	Damage int
	Armor  int
}

type Attacker struct {
	HitPoints int
	Damage    int
	Armor     int
}

func (a Attacker) damageTo(b Attacker) int {
	d := a.Damage - b.Armor
	if d < 1 {
		d = 1
	}
	return d
}

func ringCombinations() [][]Item {
	var combos [][]Item
	combos = append(combos, nil)
	for _, ring := range Rings {
		combo := []Item{ring}
		combos = append(combos, combo)
	}
	for a := 0; a < len(Rings)-1; a++ {
		for b := a + 1; b < len(Rings); b++ {
			combo := []Item{Rings[a], Rings[b]}
			combos = append(combos, combo)
		}
	}
	return combos
}

func checkPlayerWins(player Attacker, boss Attacker) bool {
	pd := player.damageTo(boss)
	ph := boss.HitPoints / pd
	if boss.HitPoints%pd > 0 {
		ph++
	}
	bd := boss.damageTo(player)
	bh := player.HitPoints / bd
	if player.HitPoints%bd > 0 {
		bh++
	}
	return ph <= bh
}

func battles(boss Attacker) ([]int, []int) {
	ringCombos := ringCombinations()
	var wins []int
	var losses []int
	for w := 0; w < len(Weapons); w++ {
		for a := -1; a < len(Armor); a++ {
			for _, rc := range ringCombos {
				player := Attacker{HitPoints: 100}
				cost := 0
				player.Damage += Weapons[w].Damage
				cost += Weapons[w].Cost
				if a >= 0 {
					player.Armor += Armor[a].Armor
					cost += Armor[a].Cost
				}
				for _, ring := range rc {
					player.Damage += ring.Damage
					player.Armor += ring.Armor
					cost += ring.Cost
				}
				if checkPlayerWins(player, boss) {
					wins = append(wins, cost)
				} else {
					losses = append(losses, cost)
				}
			}
		}
	}
	sort.Ints(wins)
	sort.Ints(losses)
	return wins, losses
}
