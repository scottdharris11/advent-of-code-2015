package day13

import (
	"log"
	"strings"
	"time"

	"advent-of-code-2015/utils"
)

type Puzzle struct{}

func (Puzzle) Solve() {
	input := utils.ReadLines("day13", "day-13-input.txt")
	solvePart1(input)
	solvePart2(input)
}

func solvePart1(lines []string) int {
	start := time.Now().UnixMilli()
	so := NewSeatingOptimizer(lines, false)
	ans := so.Optimize()
	end := time.Now().UnixMilli()
	log.Printf("Day 13, Part 1 (%dms): Happiness = %d", end-start, ans)
	return ans
}

func solvePart2(lines []string) int {
	start := time.Now().UnixMilli()
	so := NewSeatingOptimizer(lines, true)
	ans := so.Optimize()
	end := time.Now().UnixMilli()
	log.Printf("Day 13, Part 2 (%dms): Happiness Change = %d", end-start, ans)
	return ans
}

func NewSeatingOptimizer(arrangements []string, addMe bool) *SeatingOptimizer {
	happiness := make(map[string]map[string]int, len(arrangements))

	firstGuest := ""
	for _, r := range arrangements {
		pieces := strings.Split(r, " ")
		name1 := pieces[0]
		name2 := strings.ReplaceAll(pieces[10], ".", "")
		happyUnits := utils.Number(pieces[3])
		if pieces[2] == "lose" {
			happyUnits *= -1
		}

		units, ok := happiness[name1]
		if !ok {
			units = make(map[string]int)
			happiness[name1] = units
		}
		units[name2] = happyUnits

		if firstGuest == "" {
			firstGuest = name1
		}
	}

	if addMe {
		happiness["Me"] = make(map[string]int)
		for k1 := range happiness {
			happiness["Me"][k1] = 0
			happiness[k1]["Me"] = 0
		}
	}

	worst := int(^uint(0) >> 1)
	best := -worst - 1
	for k1 := range happiness {
		for k2 := range happiness {
			if k1 == k2 {
				continue
			}
			happy := happiness[k1][k2]
			happy += happiness[k2][k1]
			if happy < worst {
				worst = happy
			}
			if happy > best {
				best = happy
			}
		}
	}

	return &SeatingOptimizer{
		firstGuest: firstGuest,
		happiness:  happiness,
		guestCnt:   len(happiness),
		leastHappy: worst,
		mostHappy:  best,
	}
}

type SeatingOptimizer struct {
	happiness  map[string]map[string]int
	firstGuest string
	guestCnt   int
	leastHappy int
	mostHappy  int
}

func (so *SeatingOptimizer) Optimize() int {
	best := 0
	for name := range so.happiness {
		search := utils.Search{Searcher: so}
		solution := search.Best(utils.SearchMove{
			Cost: 0,
			State: SeatingState{
				current: name,
				seated:  1,
				arrange: name,
			},
		})
		happy := solution.Path[len(solution.Path)-1].(SeatingState).happiness
		if happy > best {
			best = happy
		}
	}
	return best
}

func (so *SeatingOptimizer) Goal(state interface{}) bool {
	var seatingState = state.(SeatingState)
	return seatingState.seated == so.guestCnt
}

func (so *SeatingOptimizer) PossibleNextMoves(state interface{}) []utils.SearchMove {
	var seatingState = state.(SeatingState)
	happyUnits := so.happiness[seatingState.current]

	var moves []utils.SearchMove
	for t, h := range happyUnits {
		if strings.Contains(seatingState.arrange, t) {
			continue
		}

		happyChange := h
		happyChange += so.happiness[t][seatingState.current]
		moveCost := so.cost(happyChange)

		if seatingState.seated+1 == so.guestCnt {
			firstGuest := strings.Split(seatingState.arrange, "-")[0]
			additionalChange := so.happiness[t][firstGuest]
			additionalChange += so.happiness[firstGuest][t]
			happyChange += additionalChange
			moveCost += so.cost(additionalChange)
		}

		move := utils.SearchMove{
			Cost: moveCost,
			State: SeatingState{
				current:   t,
				seated:    seatingState.seated + 1,
				arrange:   seatingState.arrange + "-" + t,
				happiness: seatingState.happiness + happyChange,
			},
		}
		moves = append(moves, move)
	}
	return moves
}

func (so *SeatingOptimizer) DistanceFromGoal(state interface{}) int {
	var seatingState = state.(SeatingState)
	left := so.guestCnt - seatingState.seated
	return left * so.mostHappy
}

func (so *SeatingOptimizer) cost(happyUnits int) int {
	return so.mostHappy - happyUnits
}

type SeatingState struct {
	current   string
	seated    int
	arrange   string
	happiness int
}
