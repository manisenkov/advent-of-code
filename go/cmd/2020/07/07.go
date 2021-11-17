package main

import (
	"strings"

	"github.com/manisenkov/advent-of-code/go/pkg/common"
)

const startBag = "shiny gold"

type bagCount struct {
	count int
	color string
}

type rule struct {
	color   string
	contain []bagCount
}

// Solution contains solution for day 7
type Solution struct {
	rules map[string]rule
}

// Init initializes solution with input data
func (sol *Solution) Init(input []string) {
	sol.rules = make(map[string]rule)
	for _, inp := range input {
		r := parseRule(inp)
		sol.rules[r.color] = r
	}
}

// Part1 .
func (sol *Solution) Part1() common.Any {
	// Create reverse mapping to -> from
	colorMapping := map[string][]string{}
	for _, r := range sol.rules {
		for _, bc := range r.contain {
			colorMapping[bc.color] = append(colorMapping[bc.color], r.color)
		}
	}

	alreadyVisited := map[string]bool{}
	stack := []string{startBag}
	for len(stack) > 0 {
		cur := stack[0]
		stack = stack[1:]
		if alreadyVisited[cur] {
			continue
		}
		alreadyVisited[cur] = true
		stack = append(stack, colorMapping[cur]...)
	}
	return len(alreadyVisited) - 1
}

// Part2 .
func (sol *Solution) Part2() common.Any {
	return sol.getNumberOfBags(startBag) - 1
}

func (sol *Solution) getNumberOfBags(color string) int {
	res := 0
	r := sol.rules[color]
	for _, bc := range r.contain {
		res += bc.count * sol.getNumberOfBags(bc.color)
	}
	return res + 1
}

func parseRule(inp string) rule {
	ruleInp := strings.Split(inp, " contain ")
	res := rule{
		color:   strings.Join(strings.Split(ruleInp[0], " ")[0:2], " "),
		contain: []bagCount{},
	}
	if ruleInp[1] != "no other bags." {
		for _, s := range strings.Split(ruleInp[1], ", ") {
			spl := strings.Split(s, " ")
			res.contain = append(res.contain, bagCount{
				count: common.MustAtoi(spl[0]),
				color: strings.Join(spl[1:3], " "),
			})
		}
	}
	return res
}

func main() {
	common.Run(new(Solution))
}
