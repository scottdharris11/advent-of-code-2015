package day15

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"advent-of-code-2015/utils"
)

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, 18965440, solvePart1(utils.ReadLines("day15", "day-15-input.txt")))
}

func TestSolvePart2(t *testing.T) {
	assert.Equal(t, 15862900, solvePart2(utils.ReadLines("day15", "day-15-input.txt")))
}

func TestRecipe_Score(t *testing.T) {
	ingredients := []Ingredient{
		{"Butterscotch", -1, -2, 6, 3, 8},
		{"Cinnamon", 2, 3, -2, -1, 3},
	}
	tests := []struct {
		name    string
		weights []int
		score   int
	}{
		{"1", []int{44, 56}, 62842880},
		{"2", []int{67, 33}, 0},
		{"3", []int{25, 75}, 0},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			recipe := Recipe{ingredients: ingredients, weights: tt.weights}
			assert.Equal(t, tt.score, recipe.Score())
		})
	}
}

func TestRecipe_Calories(t *testing.T) {
	ingredients := []Ingredient{
		{"Butterscotch", -1, -2, 6, 3, 8},
		{"Cinnamon", 2, 3, -2, -1, 3},
	}
	tests := []struct {
		name     string
		weights  []int
		calories int
	}{
		{"1", []int{44, 56}, 520},
		{"2", []int{40, 60}, 500},
		{"3", []int{25, 75}, 425},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			recipe := Recipe{ingredients: ingredients, weights: tt.weights}
			assert.Equal(t, tt.calories, recipe.Calories())
		})
	}
}

func TestOptimizeRecipe(t *testing.T) {
	ingredients := []string{
		"Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8",
		"Cinnamon: capacity 2, durability 3, flavor -2, texture -1, calories 3",
	}
	recipe := Recipe{ingredients: parseIngredients(ingredients)}
	assert.Equal(t, 62842880, OptimizeRecipe(&recipe, 0))
}

func TestOptimizeRecipe_WithCalorieAmt(t *testing.T) {
	ingredients := []string{
		"Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8",
		"Cinnamon: capacity 2, durability 3, flavor -2, texture -1, calories 3",
	}
	recipe := Recipe{ingredients: parseIngredients(ingredients)}
	assert.Equal(t, 57600000, OptimizeRecipe(&recipe, 500))
}
