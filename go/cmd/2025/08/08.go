package main

import (
	"slices"
	"strings"

	"github.com/manisenkov/advent-of-code/pkg/numbers"
	"github.com/manisenkov/advent-of-code/pkg/problem"
	"github.com/manisenkov/advent-of-code/pkg/vec"
)

type jboxDist struct {
	from, to int
	dist     float64
}

// Solution contains a solution for day 8
type Solution struct {
	jboxPositions []vec.Vec[int]
	numSteps      int
	ans1, ans2    int
}

// Init initializes the solution with the input data
func (sol *Solution) Init(input []string) {
	for _, s := range input {
		parts := strings.Split(s, ",")
		sol.jboxPositions = append(sol.jboxPositions, vec.New([]int{
			numbers.MustAtoi[int](parts[0]),
			numbers.MustAtoi[int](parts[1]),
			numbers.MustAtoi[int](parts[2]),
		}))
	}
	sol.numSteps = 1000
}

// Part1 to solve a "silver" part (for a first star)
func (sol *Solution) Part1() any {
	sz := len(sol.jboxPositions)
	distances := []jboxDist{}
	for i := range sz - 1 {
		for j := i + 1; j < sz; j++ {
			v := sol.jboxPositions[i]
			w := sol.jboxPositions[j]
			d := v.Sub(w)
			distances = append(distances, jboxDist{
				i, j,
				d.Abs(),
			})
		}
	}
	slices.SortFunc(distances, func(a, b jboxDist) int {
		if a.dist > b.dist {
			return 1
		}
		return -1
	})

	clusters := make([]int, sz)
	for i := range sz {
		clusters[i] = i
	}
	for step := 0; ; step++ {
		dist := distances[0]
		distances = distances[1:]
		p := clusters[dist.from]
		q := clusters[dist.to]
		isDone := false
		if p != q {
			updated := 0
			for i, c := range clusters {
				if c == p || c == q {
					clusters[i] = p
					updated++
				}
				if updated == sz {
					isDone = true
				}
			}
		}
		if step == sol.numSteps-1 {
			counters := make([]int, sz)
			for _, c := range clusters {
				counters[c]++
			}
			slices.SortFunc(counters, func(a, b int) int {
				if a > b {
					return -1
				}
				return 1
			})
			sol.ans1 = counters[0] * counters[1] * counters[2]
		}
		if isDone {
			sol.ans2 = sol.jboxPositions[dist.from].At(0) * sol.jboxPositions[dist.to].At(0)
			break
		}
	}
	return sol.ans1
}

// Part2 to solve a "gold" part (for a second star)
func (sol *Solution) Part2() any {
	return sol.ans2
}

func main() {
	problem.Solve(new(Solution))
}
