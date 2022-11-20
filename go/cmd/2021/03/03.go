package main

import (
	"github.com/manisenkov/advent-of-code/pkg/common"
)

// Solution contains solution for day 3
type Solution struct {
	input []string
}

// Init initializes solution with input data
func (sol *Solution) Init(input []string) {
	sol.input = input
}

// Part1 .
func (sol *Solution) Part1() any {
	bitsize := len(sol.input[0])
	gamma := ""
	epsilon := ""
	for i := 0; i < bitsize; i++ {
		sum1s := 0
		for j := range sol.input {
			if sol.input[j][i] == '1' {
				sum1s++
			}
		}
		if sum1s >= len(sol.input)/2 {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}
	return int(common.MustParseInt(gamma, 2, 32) * common.MustParseInt(epsilon, 2, 32))
}

// Part2 .
func (sol *Solution) Part2() any {
	bitsize := len(sol.input[0])

	// O2
	oxygenFilter := make([]string, len(sol.input))
	copy(oxygenFilter, sol.input)
	for i := 0; i < bitsize; i++ {
		if len(oxygenFilter) == 1 {
			break
		}

		sum1s := 0
		for j := range oxygenFilter {
			if oxygenFilter[j][i] == '1' {
				sum1s++
			}
		}
		var c byte
		if sum1s >= (len(oxygenFilter)/2)+(len(oxygenFilter)%2) {
			c = '1'
		} else {
			c = '0'
		}
		var nextOxygenFilter []string
		for _, b := range oxygenFilter {
			if b[i] == c {
				nextOxygenFilter = append(nextOxygenFilter, b)
			}
		}
		oxygenFilter = nextOxygenFilter
	}

	// CO2 scrubber
	scrubberFilter := make([]string, len(sol.input))
	copy(scrubberFilter, sol.input)
	for i := 0; i < bitsize; i++ {
		if len(scrubberFilter) == 1 {
			break
		}

		sum0s := 0
		for j := range scrubberFilter {
			if scrubberFilter[j][i] == '0' {
				sum0s++
			}
		}
		var c byte
		if sum0s <= (len(scrubberFilter)/2)-(len(scrubberFilter)%2) {
			c = '0'
		} else {
			c = '1'
		}
		var nextScrubberFilter []string
		for _, b := range scrubberFilter {
			if b[i] == c {
				nextScrubberFilter = append(nextScrubberFilter, b)
			}
		}
		scrubberFilter = nextScrubberFilter
	}

	return int(common.MustParseInt(oxygenFilter[0], 2, 32) * common.MustParseInt(scrubberFilter[0], 2, 32))
}

func main() {
	common.Run(new(Solution))
}
