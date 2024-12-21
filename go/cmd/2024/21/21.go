package main

import (
	"github.com/manisenkov/advent-of-code/pkg/collections"
	"github.com/manisenkov/advent-of-code/pkg/numbers"
	"github.com/manisenkov/advent-of-code/pkg/problem"
)

func findPaths(start string, target string, keyPad keyPadType) []string {
	paths := map[string][]string{
		start: {""},
	}
	queue := []string{start}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		curPaths := paths[cur]
		for key, to := range keyPad[cur] {
			targetPaths, ok := paths[to]
			if !ok || len(targetPaths[0]) > len(curPaths[0])+1 {
				paths[to] = collections.MapTo(curPaths, func(p string) string {
					return p + key
				})
				queue = append(queue, to)
			} else if len(targetPaths[0]) >= len(curPaths[0])+1 {
				paths[to] = append(paths[to], collections.MapTo(curPaths, func(p string) string {
					return p + key
				})...)
				queue = append(queue, to)
			}
		}
	}
	return collections.Unique(paths[target])
}

func solveCode(code string, keyPay keyPadType) []string {
	res := make(collections.Set[string])
	pos := "A"
	paths := []string{""}
	for i := 0; i < len(code); i++ {
		target := code[i : i+1]
		foundPaths := findPaths(pos, target, keyPay)
		var nextPaths []string
		for _, sp := range paths {
			for _, fp := range foundPaths {
				nextPaths = append(nextPaths, sp+fp+"A")
			}
		}
		paths = nextPaths
		pos = target
	}
	for _, p := range paths {
		res[p] = true
	}
	return collections.SetToSlice[string, []string](res)
}

// Solution contains a solution for day 21
type Solution struct {
	doorCodes []string
}

// Init initializes the solution with the input data
func (sol *Solution) Init(input []string) {
	sol.doorCodes = append([]string{}, input...)
}

// Part1 to solve a "silver" part (for a first star)
func (sol *Solution) Part1() any {
	res := 0
	for _, doorCode := range sol.doorCodes {
		cur := solveCode(doorCode, doorKeyPad)
		for i := 0; i < 2; i++ {
			next := []string{}
			for _, c := range cur {
				next = append(next, solveCode(c, robotKeyPad)...)
			}
			cur = next
		}
		r := numbers.Min(collections.MapTo(cur, func(t string) int { return len(t) })...)
		res += r * numbers.MustAtoi[int](doorCode[:len(doorCode)-1])
	}
	return res
}

// Part2 to solve a "gold" part (for a second star)
func (sol *Solution) Part2() any {
	return 0
}

func main() {
	problem.Solve(new(Solution))
}
