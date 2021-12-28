package day9

import (
	"log"
	"strings"
	"time"

	"advent-of-code-2015/utils"
)

type Puzzle struct{}

func (Puzzle) Solve() {
	input := utils.ReadLines("day9", "day-9-input.txt")
	solvePart1(input)
	solvePart2(input)
}

func solvePart1(lines []string) int {
	start := time.Now().UnixMilli()
	rp := NewRoutePlanner(lines, true)
	ans := rp.Route()
	end := time.Now().UnixMilli()
	log.Printf("Day 9, Part 1 (%dms): Shortest Route = %d", end-start, ans)
	return ans
}

func solvePart2(lines []string) int {
	start := time.Now().UnixMilli()
	rp := NewRoutePlanner(lines, false)
	ans := rp.Route()
	end := time.Now().UnixMilli()
	log.Printf("Day 9, Part 2 (%dms): Longest Route = %d", end-start, ans)
	return ans
}

func NewRoutePlanner(routes []string, shortest bool) *RoutePlanner {
	routeDistances := make(map[string]map[string]int, len(routes))
	smallest := 0
	largest := 0
	for _, r := range routes {
		pieces := strings.Split(r, " ")
		loc1 := pieces[0]
		loc2 := pieces[2]
		distance := utils.Number(pieces[4])

		toDistances, ok := routeDistances[loc1]
		if !ok {
			toDistances = make(map[string]int)
			routeDistances[loc1] = toDistances
		}
		toDistances[loc2] = distance

		toDistances, ok = routeDistances[loc2]
		if !ok {
			toDistances = make(map[string]int)
			routeDistances[loc2] = toDistances
		}
		toDistances[loc1] = distance

		if smallest == 0 || distance < smallest {
			smallest = distance
		}
		if distance > largest {
			largest = distance
		}
	}

	return &RoutePlanner{
		routeDistances:   routeDistances,
		locationCnt:      len(routeDistances),
		smallestDistance: smallest,
		largestDistance:  largest,
		shortest:         shortest,
	}
}

type RoutePlanner struct {
	routeDistances   map[string]map[string]int
	locationCnt      int
	smallestDistance int
	largestDistance  int
	shortest         bool
}

func (r *RoutePlanner) Route() int {
	best := 0
	for l := range r.routeDistances {
		search := utils.Search{Searcher: r}
		solution := search.Best(utils.SearchMove{
			Cost: r.cost(0),
			State: RouteState{
				current: l,
				visited: 1,
				path:    l,
			},
		})
		cost := solution.Path[len(solution.Path)-1].(RouteState).cost
		if best == 0 || (r.shortest && cost < best) || (!r.shortest && cost > best) {
			best = cost
		}
	}
	return best
}

func (r *RoutePlanner) Goal(state interface{}) bool {
	var routeState = state.(RouteState)
	return routeState.visited == r.locationCnt
}

func (r *RoutePlanner) PossibleNextMoves(state interface{}) []utils.SearchMove {
	var routeState = state.(RouteState)
	toDistances := r.routeDistances[routeState.current]

	var moves []utils.SearchMove
	for t, d := range toDistances {
		if strings.Contains(routeState.path, t) {
			continue
		}
		move := utils.SearchMove{
			Cost: r.cost(d),
			State: RouteState{
				current: t,
				visited: routeState.visited + 1,
				path:    routeState.path + "-" + t,
				cost:    routeState.cost + d,
			},
		}
		moves = append(moves, move)
	}
	return moves
}

func (r *RoutePlanner) DistanceFromGoal(state interface{}) int {
	var routeState = state.(RouteState)
	left := r.locationCnt - routeState.visited
	if r.shortest {
		return left * r.smallestDistance
	}
	return left * r.largestDistance
}

func (r *RoutePlanner) cost(distance int) int {
	if r.shortest {
		return distance
	}
	return r.largestDistance - distance
}

type RouteState struct {
	current string
	visited int
	path    string
	cost    int
}
