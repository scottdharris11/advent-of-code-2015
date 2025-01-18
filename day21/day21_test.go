package day21

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, 91, solvePart1(Boss))
}

func TestSolvePart2(t *testing.T) {
	assert.Equal(t, 158, solvePart2(Boss))
}

func TestPlayerWins(t *testing.T) {
	tests := []struct {
		name     string
		player   Attacker
		boss     Attacker
		expected bool
	}{
		{
			"Player wins",
			Attacker{100, 10, 10},
			Attacker{50, 5, 5},
			true,
		},
		{
			"Boss wins",
			Attacker{50, 5, 5},
			Attacker{100, 10, 10},
			false,
		},
		{
			"Tie but Player wins",
			Attacker{100, 10, 10},
			Attacker{100, 10, 10},
			true,
		},
		{
			"Puzzle simulation",
			Attacker{8, 5, 5},
			Attacker{12, 7, 2},
			true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, checkPlayerWins(tt.player, tt.boss))
		})
	}
}
