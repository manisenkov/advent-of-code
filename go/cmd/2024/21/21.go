package main

import (
	"strings"

	"github.com/manisenkov/advent-of-code/pkg/collections"
	"github.com/manisenkov/advent-of-code/pkg/numbers"
	"github.com/manisenkov/advent-of-code/pkg/problem"
)

var (
	robotPathsCache = map[[2]string][]string{}
	partsCache      = map[string]map[int]int{}
)

func findPaths(start string, target string, keyPad keyPadType) []string {
	if start == target {
		return []string{""}
	}
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

func findRobotPaths(start string, target string) []string {
	if res, ok := robotPathsCache[[2]string{start, target}]; ok {
		return res
	}
	res := findPaths(start, target, robotKeyPad)
	robotPathsCache[[2]string{start, target}] = res
	return res
}

func solveDoorCode(code string) []string {
	res := make(collections.Set[string])
	pos := "A"
	paths := []string{""}
	for i := 0; i < len(code); i++ {
		target := code[i : i+1]
		foundPaths := findPaths(pos, target, doorKeyPad)
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

func solveRobotCode(code string) []string {
	res := make(collections.Set[string])
	pos := "A"
	paths := []string{""}
	for i := 0; i < len(code); i++ {
		target := code[i : i+1]
		foundPaths := findRobotPaths(pos, target)
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

func solveParts(code string, steps int) int {
	if steps == 0 {
		return len(code)
	}
	if res, ok := partsCache[code][steps]; ok {
		return res
	}
	parts := strings.Split(code, "A")
	parts = parts[:len(parts)-1]
	res := 0
	for _, part := range parts {
		next := solveRobotCode(part + "A")
		res += numbers.Min(
			collections.MapTo(
				next,
				func(s string) int {
					return solveParts(s, steps-1)
				},
			)...,
		)
	}
	if c, ok := partsCache[code]; ok {
		c[steps] = res
	} else {
		partsCache[code] = map[int]int{steps: res}
	}
	return res
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
		doorPaths := solveDoorCode(doorCode)
		minVal := 0xFFFFFFFFFFFFF
		for _, p := range doorPaths {
			n := solveParts(p, 2)
			if n < minVal {
				minVal = n
			}
		}
		res += minVal * numbers.MustAtoi[int](doorCode[:len(doorCode)-1])
	}
	return res
}

// Part2 to solve a "gold" part (for a second star)
func (sol *Solution) Part2() any {
	res := 0
	for _, doorCode := range sol.doorCodes {
		doorPaths := solveDoorCode(doorCode)
		minVal := 0xFFFFFFFFFFFFF
		for _, p := range doorPaths {
			n := solveParts(p, 25)
			if n < minVal {
				minVal = n
			}
		}
		res += minVal * numbers.MustAtoi[int](doorCode[:len(doorCode)-1])
	}
	return res
}

func main() {
	problem.Solve(new(Solution))
}
