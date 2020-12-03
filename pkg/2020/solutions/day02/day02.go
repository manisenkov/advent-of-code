package day02

import (
	"regexp"

	"github.com/manisenkov/advent-of-code/pkg/common"
)

var inputRegex = regexp.MustCompile(`^(\d+)-(\d+) ([a-z]{1}): ([a-z]+)$`)

type policy struct {
	pos1, pos2 int
	letter     rune
}

// Solution contains solution for day 2
type Solution struct {
	policies  []policy
	passwords []string
}

func validatePasswordOld(pol policy, password string) bool {
	letterBuckets := map[rune]int{}
	for _, c := range password {
		letterBuckets[c]++
	}
	return letterBuckets[pol.letter] >= pol.pos1 && letterBuckets[pol.letter] <= pol.pos2
}

func validatePasswordNew(pol policy, password string) bool {
	sPassword := []rune(password)
	return (sPassword[pol.pos1-1] == pol.letter && sPassword[pol.pos2-1] != pol.letter) ||
		(sPassword[pol.pos1-1] != pol.letter && sPassword[pol.pos2-1] == pol.letter)
}

// Init initializes solution with input data
func (s *Solution) Init(input []string) {
	policies := make([]policy, len(input))
	passwords := make([]string, len(input))
	for i, inp := range input {
		m := inputRegex.FindAllStringSubmatch(inp, -1)
		pol := policy{
			pos1:   common.MustAtoi(m[0][1]),
			pos2:   common.MustAtoi(m[0][2]),
			letter: ([]rune(m[0][3]))[0],
		}
		password := m[0][4]
		policies[i] = pol
		passwords[i] = password
	}
	s.policies = policies
	s.passwords = passwords
}

// Part1 .
func (s *Solution) Part1() common.Any {
	result := 0
	for i, p := range s.passwords {
		if validatePasswordOld(s.policies[i], p) {
			result++
		}
	}
	return result
}

// Part2 .
func (s *Solution) Part2() common.Any {
	result := 0
	for i, p := range s.passwords {
		if validatePasswordNew(s.policies[i], p) {
			result++
		}
	}
	return result
}
