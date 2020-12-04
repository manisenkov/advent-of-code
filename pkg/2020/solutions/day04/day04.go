package day04

import (
	"strings"

	"github.com/manisenkov/advent-of-code/pkg/common"
)

// Solution contains solution for day 4
type Solution struct {
	passports []map[string]string
}

// Init initializes solution with input data
func (s *Solution) Init(input []string) {
	passports := make([]map[string]string, 0)
	currentPassport := make(map[string]string)
	for _, inp := range input {
		if inp == "" {
			passports = append(passports, currentPassport)
			currentPassport = make(map[string]string)
			continue
		}
		pairs := strings.Split(inp, " ")
		for _, pair := range pairs {
			keyValue := strings.Split(pair, ":")
			currentPassport[keyValue[0]] = keyValue[1]
		}
	}
	s.passports = passports
}

// Part1 .
func (s *Solution) Part1() common.Any {
	res := 0
	for _, p := range s.passports {
		_, cidPresent := p["cid"]
		if (len(p) == 7 && !cidPresent) || len(p) > 7 {
			res++
		}
	}
	return res
}

// Part2 .
func (s *Solution) Part2() common.Any {
	return 0
}
