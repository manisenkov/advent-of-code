package main

import (
	"regexp"
	"strings"

	"github.com/manisenkov/advent-of-code/pkg/collections"
	"github.com/manisenkov/advent-of-code/pkg/numbers"
	"github.com/manisenkov/advent-of-code/pkg/problem"
)

type command int

const (
	mulCmd command = iota
	doCmd
	dontCmd
)

type instruction = struct {
	typ  command
	args [2]int
}

// Solution contains a solution for day 3
type Solution struct {
	insts []instruction
}

// Init initializes the solution with the input data
func (sol *Solution) Init(input []string) {
	inp := strings.Join(input, "")
	re := regexp.MustCompile(`don't|do|mul\((\d{1,3}),(\d{1,3})\)`)
	matches := re.FindAllSubmatch([]byte(inp), -1)
	sol.insts = collections.MapTo(matches, func(m [][]byte) instruction {
		switch string(m[0]) {
		case "do":
			return instruction{
				typ: doCmd,
			}
		case "don't":
			return instruction{
				typ: dontCmd,
			}
		default:
			return instruction{
				typ: mulCmd,
				args: [2]int{
					numbers.MustAtoi[int](string(m[1])),
					numbers.MustAtoi[int](string(m[2])),
				},
			}
		}
	})
}

// Part1 to solve a "silver" part (for a first star)
func (sol *Solution) Part1() any {
	res := 0
	for _, inst := range collections.Filter(sol.insts, func(inst instruction) bool {
		return inst.typ == mulCmd
	}) {
		res += inst.args[0] * inst.args[1]
	}
	return res
}

// Part2 to solve a "gold" part (for a second star)
func (sol *Solution) Part2() any {
	res := 0
	isEnabled := true
	for _, inst := range sol.insts {
		switch inst.typ {
		case doCmd:
			isEnabled = true
		case dontCmd:
			isEnabled = false
		case mulCmd:
			if isEnabled {
				res += inst.args[0] * inst.args[1]
			}
		}
	}
	return res
}

func main() {
	problem.Solve(new(Solution))
}
