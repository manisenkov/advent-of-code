package main

import (
	"regexp"
	"strings"

	"github.com/draffensperger/golp"
	"github.com/manisenkov/advent-of-code/pkg/collections"
	"github.com/manisenkov/advent-of-code/pkg/numbers"
	"github.com/manisenkov/advent-of-code/pkg/problem"
)

type deviceConfig struct {
	targetLights string
	buttons      [][]int
	joltages     []int
}

func (cfg deviceConfig) solveLights() int {
	lights := ""
	for range len(cfg.targetLights) {
		lights += "."
	}
	dists := map[string]int{lights: 0}
	queue := []string{lights}
	for len(queue) > 0 {
		lights = queue[0]
		queue = queue[1:]
		dist := dists[lights]
		for _, btns := range cfg.buttons {
			next := toggleLights(lights, btns)
			nextDist, ok := dists[next]
			if !ok || nextDist > dist+1 {
				dists[next] = dist + 1
				queue = append(queue, next)
			}
		}
	}
	return dists[cfg.targetLights]
}

func (cfg deviceConfig) solveJolts() int {
	numBtn := len(cfg.buttons)
	numJolts := len(cfg.joltages)

	constraints := make([][]float64, numJolts)
	obj := make([]float64, numBtn)
	for i := range constraints {
		constraints[i] = make([]float64, numBtn)
	}
	for buttonIndex, buttons := range cfg.buttons {
		for _, joltIndex := range buttons {
			constraints[joltIndex][buttonIndex] = 1
		}
		obj[buttonIndex] = 1
	}

	lp := golp.NewLP(0, numBtn)
	for i, c := range constraints {
		lp.AddConstraint(c, golp.EQ, float64(cfg.joltages[i]))
	}
	for i := range numBtn {
		lp.SetInt(i, true)
	}
	lp.SetObjFn(obj)
	lp.Solve()
	return int(numbers.Sum(lp.Variables()))
}

func toggleLights(lights string, btns []int) string {
	s := lights
	for _, i := range btns {
		if s[i] == '.' {
			s = s[:i] + "#" + s[i+1:]
		} else {
			s = s[:i] + "." + s[i+1:]
		}
	}
	return s
}

// Solution contains a solution for day 10
type Solution struct {
	configs []deviceConfig
}

// Init initializes the solution with the input data
func (sol *Solution) Init(input []string) {
	var lightsRegex, buttonsRegex, joltagesRegex *regexp.Regexp
	var err error
	if lightsRegex, err = regexp.Compile(`\[[\.#]*\]`); err != nil {
		panic(err)
	}
	if buttonsRegex, err = regexp.Compile(`\((\d+,?)*\)`); err != nil {
		panic(err)
	}
	if joltagesRegex, err = regexp.Compile(`\{(\d+,?)*\}`); err != nil {
		panic(err)
	}
	for _, s := range input {
		m := string(lightsRegex.Find([]byte(s)))
		btnInput := collections.MapTo(
			buttonsRegex.FindAll([]byte(s), -1),
			func(t []byte) string { return string(t) })
		buttons := [][]int{}
		for _, b := range btnInput {
			buttons = append(buttons,
				collections.MapTo(
					strings.Split(b[1:len(b)-1], ","),
					numbers.MustAtoi[int]))
		}
		j := string(joltagesRegex.Find([]byte(s)))
		joltages := collections.MapTo(
			strings.Split(j[1:len(j)-1], ","),
			numbers.MustAtoi[int])
		sol.configs = append(sol.configs, deviceConfig{
			targetLights: m[1 : len(m)-1],
			buttons:      buttons,
			joltages:     joltages,
		})
	}
}

// Part1 to solve a "silver" part (for a first star)
func (sol *Solution) Part1() any {
	res := 0
	for _, cfg := range sol.configs {
		res += cfg.solveLights()
	}
	return res
}

// Part2 to solve a "gold" part (for a second star)
func (sol *Solution) Part2() any {
	res := 0
	for _, cfg := range sol.configs {
		res += cfg.solveJolts()
	}
	return res
}

func main() {
	problem.Solve(new(Solution))
}
