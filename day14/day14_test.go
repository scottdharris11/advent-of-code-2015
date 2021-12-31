package day14

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"advent-of-code-2015/utils"
)

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, 2655, solvePart1(utils.ReadLines("day14", "day-14-input.txt")))
}

func TestSolvePart2(t *testing.T) {
	assert.Equal(t, 1059, solvePart2(utils.ReadLines("day14", "day-14-input.txt")))
}

func TestRunRace(t *testing.T) {
	input := []string{
		"Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.",
		"Dancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds.",
	}

	tests := []struct {
		name       string
		seconds    int
		r1Distance int
		r1Running  bool
		r2Distance int
		r2Running  bool
		distance   int
		points     int
	}{
		{"1", 1, 14, true, 16, true, 16, 1},
		{"2", 10, 140, false, 160, true, 160, 10},
		{"3", 11, 140, false, 176, false, 176, 11},
		{"4", 137, 140, true, 176, false, 176, 137},
		{"5", 147, 280, false, 176, false, 280, 139},
		{"6", 173, 280, false, 176, true, 280, 139},
		{"7", 184, 280, false, 352, false, 352, 144},
		{"8", 1000, 1120, false, 1056, false, 1120, 689},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			reindeer := parseInput(input)
			distance, points := RunRace(reindeer, tt.seconds)
			assert.Equal(t, tt.r1Distance, reindeer[0].distance)
			assert.Equal(t, tt.r1Running, reindeer[0].running)
			assert.Equal(t, tt.r2Distance, reindeer[1].distance)
			assert.Equal(t, tt.r2Running, reindeer[1].running)
			assert.Equal(t, tt.distance, distance)
			assert.Equal(t, tt.points, points)
		})
	}
}
