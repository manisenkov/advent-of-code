package main

import (
	"strings"

	"github.com/manisenkov/advent-of-code/pkg/numbers"
	"github.com/manisenkov/advent-of-code/pkg/problem"
)

// Solution contains solution for day 16
type Solution struct {
	rules        map[string][][2]int
	yourTicket   []int
	otherTickets [][]int
	validTickets [][]int
}

// Init initializes solution with input data
func (sol *Solution) Init(input []string) {
	line := 0

	// Parse rules
	sol.rules = make(map[string][][2]int)
	for input[line] != "" {
		parts := strings.Split(input[line], ": ")
		name := parts[0]
		rangeParts := strings.Split(parts[1], " or ")
		ranges := make([][2]int, len(rangeParts))
		for i, part := range rangeParts {
			vals := strings.Split(part, "-")
			ranges[i] = [2]int{numbers.MustAtoi[int](vals[0]), numbers.MustAtoi[int]((vals[1]))}
		}
		sol.rules[name] = ranges
		line++
	}
	line += 2

	// Parse your ticket
	parts := strings.Split(input[line], ",")
	sol.yourTicket = make([]int, len(parts))
	for i, part := range parts {
		sol.yourTicket[i] = numbers.MustAtoi[int](part)
	}
	line += 3

	// Parse other tickets
	sol.otherTickets = make([][]int, len(input)-line)
	i := 0
	for line < len(input) {
		parts := strings.Split(input[line], ",")
		sol.otherTickets[i] = make([]int, len(parts))
		for j, part := range parts {
			sol.otherTickets[i][j] = numbers.MustAtoi[int](part)
		}
		i++
		line++
	}
}

// Part1 .
func (sol *Solution) Part1() any {
	res := 0
	sol.validTickets = [][]int{}
	for _, ticket := range sol.otherTickets {
		isValidTicket := true
		for _, field := range ticket {
			isValid := false
			for _, ranges := range sol.rules {
				if checkRanges(field, ranges) {
					isValid = true
					break
				}
			}
			if !isValid {
				res += field
				isValidTicket = false
			}
		}
		if isValidTicket {
			sol.validTickets = append(sol.validTickets, ticket)
		}
	}
	return res
}

// Part2 .
func (sol *Solution) Part2() any {
	possibleRules := make([][]string, len(sol.rules))
	for i := 0; i < len(sol.rules); i++ {
		possibleRules[i] = make([]string, 0)
		for name, ranges := range sol.rules {
			isPossibleRule := true
			for _, ticket := range sol.validTickets {
				if !checkRanges(ticket[i], ranges) {
					isPossibleRule = false
					break
				}
			}
			if isPossibleRule {
				possibleRules[i] = append(possibleRules[i], name)
			}
		}
	}
	smallestGroupIdx := 0
	smallestGroupSize := 100000
	for i, p := range possibleRules {
		if len(p) < smallestGroupSize {
			smallestGroupSize = len(p)
			smallestGroupIdx = i
		}
	}
	order := make([]string, len(sol.rules))
	findOrder(order, smallestGroupIdx, map[string]bool{}, possibleRules, 0)
	res := 1
	for i, name := range order {
		if strings.HasPrefix(name, "departure") {
			res *= sol.yourTicket[i]
		}
	}
	return res
}

func checkRanges(field int, ranges [][2]int) bool {
	for _, rng := range ranges {
		if field >= rng[0] && field <= rng[1] {
			return true
		}
	}
	return false
}

func findOrder(order []string, idx int, taken map[string]bool, possibleRules [][]string, total int) bool {
	if total == len(order) {
		return true
	}
	for _, p := range possibleRules[idx] {
		if taken[p] {
			continue
		}
		order[idx] = p
		taken[p] = true

		// Find next smallest group
		smallestGroupIdx := 0
		smallestGroupSize := 100000
		for i, p := range possibleRules {
			if order[i] == "" && len(p) < smallestGroupSize {
				smallestGroupSize = len(p)
				smallestGroupIdx = i
			}
		}

		res := findOrder(order, smallestGroupIdx, taken, possibleRules, total+1)
		if res {
			return true
		}
		taken[p] = false
		order[idx] = ""
	}
	return false
}

func main() {
	problem.Solve(new(Solution))
}
