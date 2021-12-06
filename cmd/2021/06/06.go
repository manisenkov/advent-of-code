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
		nextBuckets := [9]int64{}
		for age, count := range ageBuckets {
			if age > 0 {
				nextBuckets[age-1] += count
			} else {
				nextBuckets[8] += count
				nextBuckets[6] += count
			}
		}
		ageBuckets = nextBuckets
	}
	res := int64(0)
	for _, count := range ageBuckets {
		res += count
	}
	return res
}

// Part1 .
func (sol *Solution) Part1() common.Any {
	return sol.solve(80)
}

// Part2 .
func (sol *Solution) Part2() common.Any {
	return sol.solve(256)
}

func main() {
	common.Run(new(Solution))
}
