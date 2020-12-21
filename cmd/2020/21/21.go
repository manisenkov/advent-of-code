package main

import (
	"sort"
	"strings"

	"github.com/manisenkov/advent-of-code/pkg/common"
)

type product struct {
	ingredients map[string]bool
	allergens   map[string]bool
}

// Solution contains solution for day 21
type Solution struct {
	products             []product
	allergenToProduct    map[string][]int
	ingredientToAllergen map[string]string
}

// Init initializes solution with input data
func (sol *Solution) Init(input []string) {
	sol.products = make([]product, len(input))
	for i, inp := range input {
		parts := strings.Split(inp, " (contains ")
		parts[1] = strings.TrimSuffix(parts[1], ")")
		ingredientParts := strings.Split(parts[0], " ")
		allergenParts := strings.Split(parts[1], ", ")
		ingredients := make(map[string]bool)
		allergens := make(map[string]bool)
		for _, p := range ingredientParts {
			ingredients[p] = true
		}
		for _, p := range allergenParts {
			allergens[p] = true
		}
		sol.products[i] = product{
			ingredients: ingredients,
			allergens:   allergens,
		}
	}
	sol.allergenToProduct = make(map[string][]int)
	for i, prod := range sol.products {
		for allergen := range prod.allergens {
			sol.allergenToProduct[allergen] = append(sol.allergenToProduct[allergen], i)
		}
	}
}

// Part1 .
func (sol *Solution) Part1() common.Any {
	ingredientToAllergen := make(map[string]string)
	for _, prod := range sol.products {
		for ingredient := range prod.ingredients {
			ingredientToAllergen[ingredient] = ""
		}
	}

	ingredientToAllergen, ok := sol.solve(ingredientToAllergen, map[string]bool{})
	if !ok {
		panic("Can't solve :(")
	}

	res := 0
	for _, prod := range sol.products {
		for ingredient := range prod.ingredients {
			if ingredientToAllergen[ingredient] == "" {
				res++
			}
		}
	}
	sol.ingredientToAllergen = ingredientToAllergen

	return res
}

// Part2 .
func (sol *Solution) Part2() common.Any {
	ingredients := make([]string, 0)
	for ingredient, allergen := range sol.ingredientToAllergen {
		if allergen == "" {
			continue
		}
		ingredients = append(ingredients, ingredient)
	}
	sort.Slice(ingredients, func(i, j int) bool {
		return sol.ingredientToAllergen[ingredients[i]] < sol.ingredientToAllergen[ingredients[j]]
	})
	return strings.Join(ingredients, ",")
}

func (sol *Solution) solve(ingredientToAllergen map[string]string, allergensTaken map[string]bool) (map[string]string, bool) {
	for _, prod := range sol.products {
		for ingredient := range prod.ingredients {
			if ingredientToAllergen[ingredient] != "" {
				continue
			}

			// Try to assign allergen to this ingredient
			for allergen := range prod.allergens {
				if allergensTaken[allergen] {
					continue
				}

				if sol.checkIngredientToAllergen(ingredient, allergen) {
					ingredientToAllergen[ingredient] = allergen
					allergensTaken[allergen] = true
					_, ok := sol.solve(ingredientToAllergen, allergensTaken)
					if !ok {
						ingredientToAllergen[ingredient] = ""
						delete(allergensTaken, allergen)
					}
				}
			}
		}
	}
	return ingredientToAllergen, len(allergensTaken) == len(sol.allergenToProduct)
}

func (sol *Solution) checkIngredientToAllergen(ingredient, allergen string) bool {
	for _, prodID := range sol.allergenToProduct[allergen] {
		if !sol.products[prodID].ingredients[ingredient] {
			return false
		}
	}
	return true
}

func main() {
	common.Run(new(Solution))
}
