package main

import (
	"strings"

	"github.com/manisenkov/advent-of-code/pkg/common"
)

// Solution contains solution for day 6
type Solution struct {
	ages []int
}

// Init initializes solution with input data
func (sol *Solution) Init(input []string) {
	xs := strings.Split(input[0], ",")
	sol.ages = make([]int, len(xs))
	for i, s := range xs {
		sol.ages[i] = common.MustAtoi(s)
	}
}

func (sol *Solution) solve(days int) int64 {
	ageBuckets := [9]int64{}
	for _, age := range sol.ages {
		ageBuckets[age]++
	}
	for day := 0; day < days; day++ {
		count0 := ageBuckets[0]
		ageBuckets[0] = ageBuckets[1]
		ageBuckets[1] = ageBuckets[2]
		ageBuckets[2] = ageBuckets[3]
		ageBuckets[3] = ageBuckets[4]
		ageBuckets[4] = ageBuckets[5]
		ageBuckets[5] = ageBuckets[6]
		ageBuckets[6] = count0 + ageBuckets[7]
		ageBuckets[7] = ageBuckets[8]
		ageBuckets[8] = count0
	}
	res := int64(0)
	for _, count := range ageBuckets {
		res += count
	}
	return res
}

// Part1 .
func (sol *Solution) Part1() any {
	return sol.solve(80)
}

// Part2 .
func (sol *Solution) Part2() any {
	return sol.solve(256)
}

func main() {
	common.Run(new(Solution))
}
