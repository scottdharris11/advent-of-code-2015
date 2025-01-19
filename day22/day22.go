package day22

import (
	"log"
	"time"
)

type Puzzle struct{}

func (Puzzle) Solve() {
	solvePart1(&Player{HitPoints: 58, Damage: 9}, &Player{HitPoints: 50, Mana: 500})
	solvePart2(&Player{HitPoints: 58, Damage: 9}, &Player{HitPoints: 50, Mana: 500})
}

func solvePart1(boss *Player, player *Player) int {
	start := time.Now().UnixMilli()
	lowest := playTurn(GameState{Boss: boss, Player: player, Cost: 0}, true, -1)
	end := time.Now().UnixMilli()
	log.Printf("Day 22, Part 1 (%dms): Lowest Cost = %d", end-start, lowest)
	return lowest
}

func solvePart2(boss *Player, player *Player) int {
	start := time.Now().UnixMilli()
	lowest := playTurn(GameState{Boss: boss, Player: player, Cost: 0, Hard: true}, true, -1)
	end := time.Now().UnixMilli()
	log.Printf("Day 22, Part 2 (%dms): Lowest Cost = %d", end-start, lowest)
	return lowest
}

func playTurn(state GameState, playerTurn bool, lowest int) int {
	// if hard, remove a hit from player
	if state.Hard {
		state.Player.HitPoints--
	}

	// check end state
	if state.Boss.HitPoints <= 0 {
		return state.Cost
	}
	if state.Player.HitPoints <= 0 {
		return -1
	}
	if lowest >= 0 && state.Cost >= lowest {
		return -1
	}

	// apply pre-effects
	state.Player.shieldEffect()
	state.Player.poisonEffect(state.Boss)
	state.Player.rechargeEffect()

	// check boss end state again to see if pre-effects have won
	if state.Boss.HitPoints <= 0 {
		return state.Cost
	}

	// boss or player turns
	if playerTurn {
		for _, move := range state.playerMoves() {
			cost := playTurn(move, false, lowest)
			if cost != -1 && (lowest == -1 || cost < lowest) {
				lowest = cost
			}
		}
	} else {
		state.Player.HitPoints -= state.Boss.damageTo(*state.Player)
		cost := playTurn(state, true, lowest)
		if cost != -1 && (lowest == -1 || cost < lowest) {
			lowest = cost
		}
	}
	return lowest
}

type Player struct {
	HitPoints     int
	Damage        int
	Armor         int
	Mana          int
	ShieldTurns   int
	PoisonTurns   int
	RechargeTurns int
}

func (p *Player) damageTo(b Player) int {
	d := p.Damage - b.Armor
	if d < 1 {
		d = 1
	}
	return d
}

func (p *Player) shieldEffect() {
	if p.ShieldTurns > 0 {
		p.ShieldTurns--
		if p.ShieldTurns == 0 {
			p.Armor = 0
		}
	}
}

func (p *Player) poisonEffect(b *Player) {
	if p.PoisonTurns > 0 {
		p.PoisonTurns--
		b.HitPoints -= 3
	}
}

func (p *Player) rechargeEffect() {
	if p.RechargeTurns > 0 {
		p.RechargeTurns--
		p.Mana += 101
	}
}

type GameState struct {
	Boss   *Player
	Player *Player
	Cost   int
	Moves  []string
	Hard   bool
}

func (g GameState) playerMoves() []GameState {
	var moves []GameState
	// Magic Missile
	if g.Player.Mana >= 53 {
		gs := g.copy()
		gs.Cost += 53
		gs.Player.Mana -= 53
		gs.Boss.HitPoints -= 4
		gs.Moves = append(gs.Moves, "Player casts Magic Missile")
		moves = append(moves, gs)
	}
	// Drain
	if g.Player.Mana >= 73 {
		gs := g.copy()
		gs.Cost += 73
		gs.Player.Mana -= 73
		gs.Boss.HitPoints -= 2
		gs.Player.HitPoints += 2
		gs.Moves = append(gs.Moves, "Player casts Drain")
		moves = append(moves, gs)
	}
	// Shield
	if g.Player.Mana >= 113 && g.Player.ShieldTurns == 0 {
		gs := g.copy()
		gs.Cost += 113
		gs.Player.Mana -= 113
		gs.Player.Armor += 7
		gs.Player.ShieldTurns = 6
		gs.Moves = append(gs.Moves, "Player casts Shield")
		moves = append(moves, gs)
	}
	// Poison
	if g.Player.Mana >= 173 && g.Player.PoisonTurns == 0 {
		gs := g.copy()
		gs.Cost += 173
		gs.Player.Mana -= 173
		gs.Player.PoisonTurns = 6
		gs.Moves = append(gs.Moves, "Player casts Poison")
		moves = append(moves, gs)
	}
	// Recharge
	if g.Player.Mana >= 229 && g.Player.RechargeTurns == 0 {
		gs := g.copy()
		gs.Cost += 229
		gs.Player.Mana -= 229
		gs.Player.RechargeTurns = 5
		gs.Moves = append(gs.Moves, "Player casts Recharge")
		moves = append(moves, gs)
	}
	return moves
}

func (g GameState) copy() GameState {
	b := Player{HitPoints: g.Boss.HitPoints, Damage: g.Boss.Damage}
	p := Player{
		HitPoints:     g.Player.HitPoints,
		Damage:        g.Player.Damage,
		Armor:         g.Player.Armor,
		Mana:          g.Player.Mana,
		ShieldTurns:   g.Player.ShieldTurns,
		PoisonTurns:   g.Player.PoisonTurns,
		RechargeTurns: g.Player.RechargeTurns,
	}
	m := make([]string, len(g.Moves))
	copy(m, g.Moves)
	return GameState{Boss: &b, Player: &p, Cost: g.Cost, Hard: g.Hard, Moves: m}
}
