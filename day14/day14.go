package day14

import (
	"log"
	"regexp"
	"time"

	"advent-of-code-2015/utils"
)

type Puzzle struct{}

func (Puzzle) Solve() {
	input := utils.ReadLines("day14", "day-14-input.txt")
	solvePart1(input)
	solvePart2(input)
}

func solvePart1(lines []string) int {
	start := time.Now().UnixMilli()
	reindeer := parseInput(lines)
	distance, _ := RunRace(reindeer, 2503)
	ans := distance
	end := time.Now().UnixMilli()
	log.Printf("Day 14, Part 1 (%dms): Distance Leader = %d", end-start, ans)
	return ans
}

func solvePart2(lines []string) int {
	start := time.Now().UnixMilli()
	reindeer := parseInput(lines)
	_, points := RunRace(reindeer, 2503)
	ans := points
	end := time.Now().UnixMilli()
	log.Printf("Day 14, Part 2 (%dms): Point Leader = %d", end-start, ans)
	return ans
}

func parseInput(lines []string) []*RaceState {
	matcher := regexp.MustCompile(`^(.+) can fly (\d+) km/s for (\d+) seconds, but then must rest for (\d+) seconds\.$`)
	var reindeer []*RaceState
	for _, line := range lines {
		matches := matcher.FindStringSubmatch(line)
		r := &Reindeer{
			name:        matches[1],
			speed:       utils.Number(matches[2]),
			runningSecs: utils.Number(matches[3]),
			restingSecs: utils.Number(matches[4]),
		}
		reindeer = append(reindeer, &RaceState{
			reindeer:            r,
			distance:            0,
			running:             true,
			secsTillStateChange: r.runningSecs,
		})
	}
	return reindeer
}

type Reindeer struct {
	name        string
	speed       int
	runningSecs int
	restingSecs int
}

type RaceState struct {
	reindeer            *Reindeer
	distance            int
	points              int
	running             bool
	secsTillStateChange int
}

func RunRace(reindeer []*RaceState, seconds int) (distanceLeader int, pointLeader int) {
	for i := 0; i < seconds; i++ {
		for _, r := range reindeer {
			if r.running {
				r.distance += r.reindeer.speed
				if r.distance > distanceLeader {
					distanceLeader = r.distance
				}
			}
			r.secsTillStateChange--
			if r.secsTillStateChange == 0 {
				if r.running {
					r.secsTillStateChange = r.reindeer.restingSecs
				} else {
					r.secsTillStateChange = r.reindeer.runningSecs
				}
				r.running = !r.running
			}
		}

		for _, r := range reindeer {
			if r.distance == distanceLeader {
				r.points++
				if r.points > pointLeader {
					pointLeader = r.points
				}
			}
		}
	}
	return
}
