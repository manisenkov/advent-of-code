package main

import (
	"math"
	"strconv"
	"strings"

	"github.com/manisenkov/advent-of-code/pkg/numbers"
	"github.com/manisenkov/advent-of-code/pkg/problem"
)

type idRange struct {
	low  int
	high int
}

// Solution contains a solution for day 2
type Solution struct {
	ranges []idRange
}

// Split input range to a sequence of same length ranges (f.i. 15-1200 would become 15-99,100-999,1000-1200)
func getSubRanges(input idRange) []idRange {
	res := make([]idRange, 0)
	cur := input.low
	for cur < input.high {
		nextHigh := numbers.Min(input.high, int(math.Pow10(len(strconv.Itoa(cur))))-1)
		res = append(res, idRange{cur, nextHigh})
		cur = nextHigh + 1
	}
	return res
}

// Init initializes the solution with the input data
func (sol *Solution) Init(input []string) {
	rangeStrs := strings.Split(input[0], ",")
	ranges := make([]idRange, len(rangeStrs))
	for i, rs := range rangeStrs {
		parts := strings.Split(rs, "-")
		ranges[i] = idRange{
			numbers.MustAtoi[int](parts[0]),
			numbers.MustAtoi[int](parts[1]),
		}
	}
	sol.ranges = ranges
}

func repeatString(s string, times int) string {
	res := ""
	for ; times > 0; times-- {
		res += s
	}
	return res
}

// Part1 to solve a "silver" part (for a first star)
func (sol *Solution) Part1() any {
	res := 0
	for _, rang := range sol.ranges {
		subRanges := getSubRanges(rang)
		for _, subRang := range subRanges {
			rlStr := strconv.Itoa(subRang.low)
			if len(rlStr)%2 == 1 {
				continue
			}

			// Iterate over halves
			cur := numbers.MustAtoi[int](rlStr[:len(rlStr)/2])
			for {
				curStr := strconv.Itoa(cur)
				numToCheck := numbers.MustAtoi[int](curStr + curStr)
				if numToCheck <= subRang.high && len(curStr)*2 == len(rlStr) {
					if numToCheck >= subRang.low {
						res += numToCheck
					}
					cur++
				} else {
					break
				}
			}
		}
	}
	return res
}

// Part2 to solve a "gold" part (for a second star)
func (sol *Solution) Part2() any {
	res := 0
	for _, rang := range sol.ranges {
		for cur := rang.low; cur <= rang.high; cur++ {
			num := strconv.Itoa(cur)
			wrongID := false
			for i := 1; i <= len(num)/2; i++ {
				if len(num)%i != 0 {
					continue
				}
				check := repeatString(num[:i], len(num)/i)
				if check == num {
					wrongID = true
					break
				}
			}
			if wrongID {
				res += cur
			}
		}
	}
	return res
}

func main() {
	problem.Solve(new(Solution))
}
