package main

import (
	"fmt"

	"github.com/manisenkov/advent-of-code/pkg/common"
)

// Solution contains a solution for day 23
type Solution struct {
	initState state
}

// Init initializes the solution with the input data
func (sol *Solution) Init(input []string) {
	sol.initState = state{
		rooms: [4][4]byte{
			{input[2][3], input[3][3], 0, 0},
			{input[2][5], input[3][5], 0, 0},
			{input[2][7], input[3][7], 0, 0},
			{input[2][9], input[3][9], 0, 0},
		},
		depth: 2,
	}
}

// Part1 to solve a "silver" part (for a first star)
func (sol *Solution) Part1() any {
	res := map[state]int{sol.initState: 0}
	queue := []state{sol.initState}
	step := 0
	for len(queue) > 0 {
		step++
		if step == 10000 {
			fmt.Println(" -- depth: 2 | queue size:", len(queue), "| state cache size:", len(deriveCache[2]))
			step = 0
		}
		st := queue[0]
		queue = queue[1:]
		curEnergy := res[st]
		moves := st.deriveNext()
		for _, m := range moves {
			if e, ok := res[m.state]; !ok || curEnergy+m.energy < e {
				res[m.state] = curEnergy + m.energy
				if m.state != finalStateP1 {
					queue = append(queue, m.state)
				}
			}
		}
	}
	return res[finalStateP1]
}

// Part2 to solve a "gold" part (for a second star)
func (sol *Solution) Part2() any {
	p2InitState := state{
		rooms: [4][4]byte{
			{sol.initState.rooms[0][0], 'D', 'D', sol.initState.rooms[0][1]},
			{sol.initState.rooms[1][0], 'C', 'B', sol.initState.rooms[1][1]},
			{sol.initState.rooms[2][0], 'B', 'A', sol.initState.rooms[2][1]},
			{sol.initState.rooms[3][0], 'A', 'C', sol.initState.rooms[3][1]},
		},
		depth: 4,
	}
	res := map[state]int{p2InitState: 0}
	queue := []state{p2InitState}
	step := 0
	for len(queue) > 0 {
		step++
		if step == 1000000 {
			fmt.Println(" -- depth: 4 | queue size:", len(queue), "| state cache size:", len(deriveCache[4]))
			step = 0
		}
		st := queue[0]
		queue = queue[1:]
		curEnergy := res[st]
		moves := st.deriveNext()
		for _, m := range moves {
			if e, ok := res[m.state]; !ok || curEnergy+m.energy < e {
				res[m.state] = curEnergy + m.energy
				if m.state != finalStateP2 {
					queue = append(queue, m.state)
				}
			}
		}
	}
	return res[finalStateP2]
}

func main() {
	common.Run(new(Solution))
}
