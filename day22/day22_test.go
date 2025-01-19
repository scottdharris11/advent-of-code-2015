package day22

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, 226, solvePart1(&Player{HitPoints: 13, Damage: 8}, &Player{HitPoints: 10, Mana: 250}))
	assert.Equal(t, 641, solvePart1(&Player{HitPoints: 14, Damage: 8}, &Player{HitPoints: 10, Mana: 250}))
	assert.Equal(t, 1269, solvePart1(&Player{HitPoints: 58, Damage: 9}, &Player{HitPoints: 50, Mana: 500}))
}

func TestSolvePart2(t *testing.T) {
	assert.Equal(t, 1309, solvePart2(&Player{HitPoints: 58, Damage: 9}, &Player{HitPoints: 50, Mana: 500}))
}
