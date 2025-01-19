package day22

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, 0, solvePart1(Player{HitPoints: 13, Damage: 8}, Player{HitPoints: 10, Mana: 250}))
	assert.Equal(t, 0, solvePart1(Boss, Me))
}

func TestSolvePart2(t *testing.T) {
	assert.Equal(t, 0, solvePart2(Player{HitPoints: 13, Damage: 8}, Player{HitPoints: 10, Mana: 250}))
	assert.Equal(t, 0, solvePart2(Boss, Me))
}
