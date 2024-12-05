package main

import (
	"slices"
	"strings"

	"github.com/manisenkov/advent-of-code/pkg/collections"
	"github.com/manisenkov/advent-of-code/pkg/numbers"
	"github.com/manisenkov/advent-of-code/pkg/problem"
)

// Solution contains a solution for day 5
type Solution struct {
	rules            map[int]map[int]bool
	updates          [][]int
	incorrectUpdates [][]int
}

// Init initializes the solution with the input data
func (sol *Solution) Init(input []string) {
	rules := make(map[int]map[int]bool)
	idx := 0
	for i, s := range input {
		if s == "" {
			idx = i + 1
			break
		}
		parts := strings.Split(s, "|")
		a := numbers.MustAtoi[int](parts[0])
		b := numbers.MustAtoi[int](parts[1])
		if to, ok := rules[a]; ok {
			to[b] = true
		} else {
			rules[a] = map[int]bool{b: true}
		}
	}
	updates := collections.Map(input[idx:], func(s string) []int {
		return collections.Map(strings.Split(s, ","), numbers.MustAtoi[int])
	})
	sol.rules = rules
	sol.updates = updates
}

// Part1 to solve a "silver" part (for a first star)
func (sol *Solution) Part1() any {
	sol.incorrectUpdates = [][]int{}
	return collections.Reduce(
		collections.Map(sol.updates, func(upd []int) int {
			isCorrect := true
			for i, a := range upd[:len(upd)-1] {
				for _, b := range upd[i+1:] {
					bRule, ok := sol.rules[b]
					if ok && bRule[a] {
						isCorrect = false
						break
					}
				}
				if !isCorrect {
					break
				}
			}
			if isCorrect {
				return upd[(len(upd)-1)/2]
			} else {
				updCopy := make([]int, len(upd))
				copy(updCopy, upd)
				sol.incorrectUpdates = append(sol.incorrectUpdates, updCopy)
				return 0
			}
		}), func(a, b int) int {
			return a + b
		},
	)
}

// Part2 to solve a "gold" part (for a second star)
func (sol *Solution) Part2() any {
	for _, upd := range sol.incorrectUpdates {
		slices.SortStableFunc(upd, func(a, b int) int {
			aRule, ok := sol.rules[a]
			if ok && aRule[b] {
				return -1
			}
			bRule, ok := sol.rules[b]
			if ok && bRule[a] {
				return 1
			}
			return 0
		})
	}
	return collections.Reduce(
		collections.Map(
			sol.incorrectUpdates,
			func(upd []int) int {
				return upd[(len(upd)-1)/2]
			},
		),
		func(a, b int) int {
			return a + b
		},
	)
}

func main() {
	problem.Solve(new(Solution))
}
