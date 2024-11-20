package main

import (
	"regexp"

	"github.com/manisenkov/advent-of-code/pkg/numbers"
	"github.com/manisenkov/advent-of-code/pkg/problem"
)

var inputRegex = regexp.MustCompile(`^(\d+)-(\d+) ([a-z]{1}): ([a-z]+)$`)

type policy struct {
	pos1, pos2 int
	letter     rune
}

func (pol *policy) validatePasswordOld(password string) bool {
	letterBuckets := map[rune]int{}
	for _, c := range password {
		letterBuckets[c]++
	}
	return letterBuckets[pol.letter] >= pol.pos1 && letterBuckets[pol.letter] <= pol.pos2
}

func (pol *policy) validatePasswordNew(password string) bool {
	sPassword := []rune(password)
	return (sPassword[pol.pos1-1] == pol.letter && sPassword[pol.pos2-1] != pol.letter) ||
		(sPassword[pol.pos1-1] != pol.letter && sPassword[pol.pos2-1] == pol.letter)
}

// Solution contains solution for day 2
type Solution struct {
	policies  []policy
	passwords []string
}

// Init initializes solution with input data
func (sol *Solution) Init(input []string) {
	sol.policies = make([]policy, len(input))
	sol.passwords = make([]string, len(input))
	for i, inp := range input {
		m := inputRegex.FindAllStringSubmatch(inp, -1)
		pol := policy{
			pos1:   numbers.MustAtoi[int](m[0][1]),
			pos2:   numbers.MustAtoi[int](m[0][2]),
			letter: ([]rune(m[0][3]))[0],
		}
		password := m[0][4]
		sol.policies[i] = pol
		sol.passwords[i] = password
	}
}

// Part1 .
func (sol *Solution) Part1() any {
	result := 0
	for i, p := range sol.passwords {
		if sol.policies[i].validatePasswordOld(p) {
			result++
		}
	}
	return result
}

// Part2 .
func (sol *Solution) Part2() any {
	result := 0
	for i, p := range sol.passwords {
		if sol.policies[i].validatePasswordNew(p) {
			result++
		}
	}
	return result
}

func main() {
	problem.Solve(new(Solution))
}
