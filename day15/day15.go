package day15

import (
	"log"
	"regexp"
	"time"

	"advent-of-code-2015/utils"
)

type Puzzle struct{}

func (Puzzle) Solve() {
	input := utils.ReadLines("day15", "day-15-input.txt")
	solvePart1(input)
	solvePart2(input)
}

func solvePart1(lines []string) int {
	start := time.Now().UnixMilli()
	recipe := Recipe{ingredients: parseIngredients(lines)}
	ans := OptimizeRecipe(&recipe, 0)
	end := time.Now().UnixMilli()
	log.Printf("Day 15, Part 1 (%dms): Score = %d", end-start, ans)
	return ans
}

func solvePart2(lines []string) int {
	start := time.Now().UnixMilli()
	recipe := Recipe{ingredients: parseIngredients(lines)}
	ans := OptimizeRecipe(&recipe, 500)
	end := time.Now().UnixMilli()
	log.Printf("Day 15, Part 2 (%dms): Score = %d", end-start, ans)
	return ans
}

func parseIngredients(lines []string) []Ingredient {
	matcher := regexp.MustCompile(`^(.+): capacity ([-\d]+), durability ([-\d]+), flavor ([-\d]+), texture ([-\d]+), calories ([-\d]+)$`)
	var ingredients []Ingredient
	for _, line := range lines {
		matches := matcher.FindStringSubmatch(line)
		ingredients = append(ingredients, Ingredient{
			name:       matches[1],
			capacity:   utils.Number(matches[2]),
			durability: utils.Number(matches[3]),
			flavor:     utils.Number(matches[4]),
			texture:    utils.Number(matches[5]),
			calories:   utils.Number(matches[6]),
		})
	}
	return ingredients
}

type Ingredient struct {
	name       string
	capacity   int
	durability int
	flavor     int
	texture    int
	calories   int
}

type Recipe struct {
	ingredients []Ingredient
	weights     []int
}

func (r *Recipe) Score() int {
	categories := [4]int{}
	for i, ingredient := range r.ingredients {
		categories[0] += ingredient.capacity * r.weights[i]
		categories[1] += ingredient.durability * r.weights[i]
		categories[2] += ingredient.flavor * r.weights[i]
		categories[3] += ingredient.texture * r.weights[i]
	}

	score := 1
	for _, c := range categories {
		if c <= 0 {
			return 0
		}
		score *= c
	}
	return score
}

func (r *Recipe) Calories() int {
	calories := 0
	for i, ingredient := range r.ingredients {
		calories += ingredient.calories * r.weights[i]
	}
	return calories
}

func OptimizeRecipe(r *Recipe, calories int) int {
	iCnt := len(r.ingredients)
	weights := make([]int, iCnt)
	return runScenarios(r, weights, 0, 0, 100-iCnt+1, calories)
}

func runScenarios(r *Recipe, weights []int, best int, idx int, max int, calories int) int {
	b := best
	for i := 1; i <= max; i++ {
		weights[idx] = i
		if idx == len(weights)-2 {
			weights[idx+1] = max - i + 1
			r.weights = weights
			if calories == 0 || r.Calories() == calories {
				s := r.Score()
				if s > b {
					b = s
				}
			}
			continue
		}

		t := runScenarios(r, weights, b, idx+1, max-(i-1), calories)
		if t > b {
			b = t
		}
	}
	return b
}
