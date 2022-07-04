package main

import (
	"math"

	"github.com/manisenkov/advent-of-code/pkg/common"
)

// Solution contains solution for day 9
type Solution struct {
	preambleSize  int
	nums          []int64
	invalidNumber int64
}

// Init initializes solution with input data
func (sol *Solution) Init(input []string) {
	sol.nums = make([]int64, len(input))
	for i, inp := range input {
		sol.nums[i] = common.MustParseInt(inp, 10, 64)
	}
}

// Part1 .
func (sol *Solution) Part1() any {
	for i := sol.preambleSize; i < len(sol.nums); i++ {
		testNum := sol.nums[i]
		if !hasSumPair(sol.nums[i-sol.preambleSize:i], testNum) {
			sol.invalidNumber = testNum
			return testNum
		}
	}
	panic("Invalid input")
}

// Part2 .
func (sol *Solution) Part2() any {
	// Find partial sums
	partSums := make([]int64, len(sol.nums)+1)
	sum := int64(0)
	for i, n := range sol.nums {
		sum += n
		partSums[i+1] = sum
	}

	// Find contagious set
	sumFound := false
	minIdx := 0
	maxIdx := 0
	for ; minIdx < len(sol.nums)-1; minIdx++ {
		for maxIdx = minIdx + 1; maxIdx < len(sol.nums); maxIdx++ {
			sum = partSums[maxIdx+1] - partSums[minIdx]
			if sum == sol.invalidNumber {
				sumFound = true
				break
			}
		}
		if sumFound {
			break
		}
	}

	// Find min and max in the found set
	min := int64(math.MaxInt64)
	max := int64(0)
	for i := minIdx; i <= maxIdx; i++ {
		if sol.nums[i] < min {
			min = sol.nums[i]
		}
		if sol.nums[i] > max {
			max = sol.nums[i]
		}
	}
	return min + max
}

func hasSumPair(checkNums []int64, target int64) bool {
	for i := 0; i < len(checkNums)-1; i++ {
		for j := i + 1; j < len(checkNums); j++ {
			if checkNums[i]+checkNums[j] == target {
				return true
			}
		}
	}
	return false
}

func main() {
	common.Run(&Solution{preambleSize: 25})
}
