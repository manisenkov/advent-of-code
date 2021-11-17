package main

import (
	"strings"

	"github.com/manisenkov/advent-of-code/go/pkg/common"
)

type rule interface {
	match(msg []rune, rules map[int]rule) ([]rune, bool)
}

type branchRule [][]int

func (r branchRule) match(s []rune, rules map[int]rule) ([]rune, bool) {
	if len(s) == 0 {
		return s, false
	}
	for _, br := range r {
		left := s
		ok := true
		for _, nr := range br {
			left, ok = rules[nr].match(left, rules)
			if !ok {
				break
			}
		}
		if ok {
			return left, ok
		}
	}
	return s, false
}

type charRule rune

func (r charRule) match(s []rune, _ map[int]rule) ([]rune, bool) {
	if len(s) > 0 && rune(r) == s[0] {
		return s[1:], true
	}
	return s, false
}

// Solution contains solution for day 19
type Solution struct {
	rules    map[int]rule
	messages [][]rune
}

// Init initializes solution with input data
func (sol *Solution) Init(input []string) {
	i := 0

	rules := make(map[int]rule)
	for _, inp := range input {
		if inp == "" {
			i++
			break
		}
		parts := strings.Split(inp, ": ")
		ruleID := common.MustAtoi(parts[0])
		if strings.HasPrefix(parts[1], "\"") && strings.HasSuffix(parts[1], "\"") {
			rules[ruleID] = charRule([]rune(parts[1])[1])
		} else {
			parts = strings.Split(parts[1], " | ")
			branches := make([][]int, len(parts))
			for j, p := range parts {
				idxParts := strings.Split(p, " ")
				branches[j] = make([]int, len(idxParts))
				for k, rp := range idxParts {
					branches[j][k] = common.MustAtoi(rp)
				}
			}
			rules[ruleID] = branchRule(branches)
		}
		i++
	}
	sol.rules = rules

	messages := make([][]rune, 0)
	for i < len(input) {
		messages = append(messages, []rune(input[i]))
		i++
	}
	sol.messages = messages
}

// Part1 .
func (sol *Solution) Part1() common.Any {
	res := 0
	for _, msg := range sol.messages {
		left, ok := sol.rules[0].match(msg, sol.rules)
		if ok && len(left) == 0 {
			res++
		}
	}
	return res
}

// Part2 .
func (sol *Solution) Part2() common.Any {
	updRules := make(map[int]rule)
	for id, r := range sol.rules {
		updRules[id] = r
	}

	// Expand rule 0 to cover all possible cases
	rule0 := branchRule([][]int{})
	countOf42s := 10
	for i := countOf42s; i >= 2; i-- {
		for j := i - 1; j > 0; j-- {
			b := make([]int, 0)
			for k := 0; k < i; k++ {
				b = append(b, 42)
			}
			for k := 0; k < j; k++ {
				b = append(b, 31)
			}
			rule0 = append(rule0, b)
		}
	}
	updRules[0] = rule0

	res := 0
	for _, msg := range sol.messages {
		left, ok := updRules[0].match(msg, updRules)
		if ok && len(left) == 0 {
			res++
		}
	}
	return res
}

func main() {
	common.Run(new(Solution))
}
