package main

import (
	"strings"

	"github.com/manisenkov/advent-of-code/pkg/common"
)

// Solution contains solution for day 15
type Solution struct {
	nums []int
}

// Init initializes solution with input data
func (sol *Solution) Init(input []string) {
	parts := strings.Split(input[0], ",")
	sol.nums = make([]int, len(parts))
	for i, p := range parts {
		sol.nums[i] = common.MustAtoi(p)
	}
}

// Part1 .
func (sol *Solution) Part1() common.Any {
	return sol.solve(2020)
}

// Part2 .
func (sol *Solution) Part2() common.Any {
	return sol.solve(30000000)
}

func (sol *Solution) solve(numTurns int) int {
	mem := make(map[int][2]int)
	for i, n := range sol.nums {
		mem[n] = [2]int{i + 1, 0}
	}
	turn := len(sol.nums) + 1
	lastSpoken := sol.nums[len(sol.nums)-1]
	for {
		var toSay int
		if mem[lastSpoken][1] == 0 { // Last number is spoken for first time
			toSay = 0
		} else {
			toSay = mem[lastSpoken][0] - mem[lastSpoken][1]
		}

		if turn == numTurns {
			return toSay
		}

		_, isSpoken := mem[toSay]
		if !isSpoken {
			mem[toSay] = [2]int{turn, 0}
		} else {
			mem[toSay] = [2]int{turn, mem[toSay][0]}
		}
		lastSpoken = toSay
		turn++
	}
}

func main() {
	common.Run(new(Solution))
}
