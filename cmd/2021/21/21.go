package main

import (
	"strings"

	"github.com/manisenkov/advent-of-code/pkg/common"
)

type state struct {
	pos    [2]int
	score  [2]int
	player int
}

type step struct {
	st    state
	input int64
}

// Solution contains solution for day 21
type Solution struct {
	initState state
}

// Init initializes solution with input data
func (sol *Solution) Init(input []string) {
	startPos1 := common.MustAtoi(strings.Split(input[0], ": ")[1])
	startPos2 := common.MustAtoi(strings.Split(input[1], ": ")[1])
	sol.initState = state{
		pos:    [2]int{startPos1, startPos2},
		score:  [2]int{0, 0},
		player: 0,
	}
}

// Part1 .
func (sol *Solution) Part1() common.Any {
	st := sol.initState
	dice := 0
	rolledTimes := 0
	for {
		rolledTimes += 3
		move := loop(dice+1, 100) + loop(dice+2, 100) + loop(dice+3, 100)
		dice = loop(dice+3, 100)
		st.pos[st.player] = loop(st.pos[st.player]+move, 10)
		st.score[st.player] += st.pos[st.player]
		if st.score[st.player] >= 1000 {
			break
		}
		st.player = (st.player + 1) % 2
	}
	if st.score[0] > st.score[1] {
		return st.score[1] * rolledTimes
	} else {
		return st.score[0] * rolledTimes
	}
}

// Part2 .
func (sol *Solution) Part2() common.Any {
	const goal = 21
	diceProbabilities := map[int]int64{
		3: 1,
		4: 3,
		5: 6,
		6: 7,
		7: 6,
		8: 3,
		9: 1,
	}
	stateCounts := map[state]int64{
		sol.initState: 1,
	}
	queue := []step{{sol.initState, 1}}
	for len(queue) > 0 {
		cur := queue[0]
		st := cur.st
		input := cur.input
		queue = queue[1:]
		for move, prob := range diceProbabilities {
			nextSt := st
			nextSt.pos[st.player] = loop(st.pos[st.player]+move, 10)
			nextSt.score[st.player] += nextSt.pos[st.player]
			nextSt.player = (st.player + 1) % 2
			toAdd := prob * input
			stateCounts[nextSt] += toAdd
			if nextSt.score[st.player] < goal {
				queue = append(queue, step{nextSt, toAdd})
			}
		}
	}
	totPlay1 := int64(0)
	totPlay2 := int64(0)
	for st, count := range stateCounts {
		if st.score[0] >= goal {
			totPlay1 += count
		} else if st.score[1] >= goal {
			totPlay2 += count
		}
	}
	if totPlay1 > totPlay2 {
		return totPlay1
	} else {
		return totPlay2
	}
}

func loop(val, max int) int {
	res := val
	for res > max {
		res -= max
	}
	return res
}

func main() {
	common.Run(new(Solution))
}
