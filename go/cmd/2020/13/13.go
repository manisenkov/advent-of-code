package main

import (
	"math"
	"strings"

	"github.com/manisenkov/advent-of-code/go/pkg/common"
)

// Solution contains solution for day 13
type Solution struct {
	ts     int64
	busIDs []int64
}

// Init initializes solution with input data
func (sol *Solution) Init(input []string) {
	sol.ts = common.MustParseInt(input[0], 10, 64)
	for _, inp := range strings.Split(input[1], ",") {
		if inp == "x" {
			sol.busIDs = append(sol.busIDs, -1)
			continue
		}
		sol.busIDs = append(sol.busIDs, common.MustParseInt(inp, 10, 64))
	}
}

// Part1 .
func (sol *Solution) Part1() common.Any {
	minTimeToWait := int64(math.MaxInt64)
	busIDToWait := int64(0)
	for _, busID := range sol.busIDs {
		if busID == -1 {
			continue
		}
		timeToWait := busID - sol.ts%busID
		if timeToWait < minTimeToWait {
			minTimeToWait = timeToWait
			busIDToWait = busID
		}
	}
	return minTimeToWait * busIDToWait
}

// Part2 .
func (sol *Solution) Part2() common.Any {
	// Can be solved as system of congruence equations. We're looking into *end* of the window (not the begin).
	// Algorithm is taken from https://math.stackexchange.com/questions/79282/solving-simultaneous-congruences
	// Calculate factors represented as [2]int64{r, bus_id} for congruent equations such as
	//
	//   N === r_0 mod bus_id_0
	//   N === r_1 mod bus_id_1
	//   ...
	//   N === r_k mod bus_id_k
	//
	// Keep in mind that r_k == 0 by definition
	factors := make([][2]int64, 0)
	for i := len(sol.busIDs) - 1; i >= 0; i-- {
		if sol.busIDs[i] == -1 {
			continue
		}
		factors = append(factors, [2]int64{int64(len(sol.busIDs) - i - 1), sol.busIDs[i]})
	}

	n := int64(1)
	for _, f := range factors {
		n *= f[1]
	}
	res := int64(0)
	for _, f := range factors {
		_, y, _ := common.ExtGCD(n/f[1], f[1])
		res += y * (n / f[1]) * f[0]
	}

	// Since we're calculated *end* of the window, we need to subtract length of the window to return
	// *begin* of it as requested by problem statement.
	return res%n - int64(len(sol.busIDs)) + 1
}

func main() {
	common.Run(new(Solution))
}
