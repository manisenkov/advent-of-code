package main

import (
	"strings"

	"github.com/manisenkov/advent-of-code/pkg/common"
)

// Solution contains solution for day 14
type Solution struct {
	template []byte
	rules    map[[2]byte]byte
}

// Init initializes solution with input data
func (sol *Solution) Init(input []string) {
	sol.template = []byte(input[0])
	sol.rules = map[[2]byte]byte{}
	for _, s := range input[2:] {
		xs := strings.Split(s, " -> ")
		sol.rules[[2]byte{xs[0][0], xs[0][1]}] = xs[1][0]
	}
}

// Part1 .
func (sol *Solution) Part1() any {
	return sol.polymerize(10)
}

// Part2 .
func (sol *Solution) Part2() any {
	return sol.polymerize(40)
}

func (sol *Solution) polymerize(numSteps int) int64 {
	pairBuckets := map[[2]byte]int64{}
	letterCount := map[byte]int64{}
	for i := 0; i < len(sol.template)-1; i++ {
		pair := [2]byte{sol.template[i], sol.template[i+1]}
		pairBuckets[pair]++
		letterCount[sol.template[i]]++
	}
	letterCount[sol.template[len(sol.template)-1]]++
	for i := 0; i < numSteps; i++ {
		bucketsToAdd := map[[2]byte]int64{}
		for pair, count := range pairBuckets {
			result, ok := sol.rules[pair]
			if !ok {
				continue
			}
			nextPair1 := [2]byte{pair[0], result}
			nextPair2 := [2]byte{result, pair[1]}
			bucketsToAdd[pair] -= count
			bucketsToAdd[nextPair1] += count
			bucketsToAdd[nextPair2] += count
			letterCount[result] += count
		}
		for pair, count := range bucketsToAdd {
			pairBuckets[pair] += count
			if pairBuckets[pair] == 0 {
				delete(pairBuckets, pair)
			}
		}
	}
	var maxCount int64 = -0x7fffffffffffffff
	var minCount int64 = 0x7fffffffffffffff
	for _, n := range letterCount {
		if n > maxCount {
			maxCount = n
		}
		if n < minCount {
			minCount = n
		}
	}
	return maxCount - minCount
}

func main() {
	common.Run(new(Solution))
}
