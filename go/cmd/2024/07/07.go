package main

import (
	"strconv"
	"strings"

	"github.com/manisenkov/advent-of-code/pkg/collections"
	"github.com/manisenkov/advent-of-code/pkg/numbers"
	"github.com/manisenkov/advent-of-code/pkg/problem"
)

type equation struct {
	res  int
	nums []int
}

// Solution contains a solution for day 7
type Solution struct {
	equations []equation
}

func solve(res int, nums []int, withConcat bool) bool {
	if len(nums) == 0 && res == 0 {
		return true
	}
	if len(nums) == 0 && res != 0 {
		return false
	}
	if len(nums) == 1 && res == nums[0] {
		return true
	}
	if len(nums) == 1 && res != nums[0] {
		return false
	}
	lastIdx := len(nums) - 1
	if res%nums[lastIdx] == 0 && solve(res/nums[lastIdx], nums[:lastIdx], withConcat) {
		return true
	}
	if res > nums[lastIdx] && solve(res-nums[lastIdx], nums[:lastIdx], withConcat) {
		return true
	}
	if !withConcat || len(nums) == 1 {
		return false
	}
	resStr := strconv.Itoa(res)
	lastNumStr := strconv.Itoa(nums[lastIdx])
	if resStr != lastNumStr &&
		strings.HasSuffix(resStr, lastNumStr) &&
		solve(numbers.MustAtoi[int](resStr[:len(resStr)-len(lastNumStr)]), nums[:lastIdx], withConcat) {
		return true
	}
	return false
}

// Init initializes the solution with the input data
func (sol *Solution) Init(input []string) {
	sol.equations = collections.MapTo(input, func(s string) equation {
		parts := strings.Split(s, ": ")
		return equation{
			res:  numbers.MustAtoi[int](parts[0]),
			nums: collections.MapTo(strings.Split(parts[1], " "), numbers.MustAtoi[int]),
		}
	})
}

// Part1 to solve a "silver" part (for a first star)
func (sol *Solution) Part1() any {
	return numbers.Sum(collections.MapTo(sol.equations, func(eq equation) int {
		if solve(eq.res, eq.nums, false) {
			return eq.res
		}
		return 0
	}))
}

// Part2 to solve a "gold" part (for a second star)
func (sol *Solution) Part2() any {
	return numbers.Sum(collections.MapTo(sol.equations, func(eq equation) int {
		if solve(eq.res, eq.nums, true) {
			return eq.res
		}
		return 0
	}))
}

func main() {
	problem.Solve(new(Solution))
}
