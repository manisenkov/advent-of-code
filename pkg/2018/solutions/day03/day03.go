package day03

import (
	"fmt"
	"regexp"

	"github.com/manisenkov/advent-of-code/pkg/common"
)

var claimRegex = regexp.MustCompile(`^#(\d+) @ (\d+),(\d+): (\d+)x(\d+)$`)

// Solution contains solution for day 3
type Solution struct {
	ids    []string
	fabric map[string][]string
}

// Init initializes solution with input data
func (s *Solution) Init(input []string) {
	s.fabric = make(map[string][]string)
	s.ids = make([]string, len(input))
	for i, inp := range input {
		m := claimRegex.FindAllStringSubmatch(inp, -1)
		id := m[0][1]
		left := common.MustAtoi(m[0][2])
		top := common.MustAtoi(m[0][3])
		width := common.MustAtoi(m[0][4])
		height := common.MustAtoi(m[0][5])
		for x := left; x < left+width; x++ {
			for y := top; y < top+height; y++ {
				k := fmt.Sprintf("%v,%v", x, y)
				s.fabric[k] = append(s.fabric[k], id)
			}
		}
		s.ids[i] = id
	}
}

// Part1 .
func (s *Solution) Part1() common.Any {
	res := 0
	for _, ids := range s.fabric {
		if len(ids) > 1 {
			res++
		}
	}
	return res
}

// Part2 .
func (s *Solution) Part2() common.Any {
	idsLeft := map[string]bool{}
	for _, id := range s.ids {
		idsLeft[id] = true
	}
	for _, ids := range s.fabric {
		if len(ids) > 1 {
			for _, id := range ids {
				delete(idsLeft, id)
			}
		}
	}
	for id := range idsLeft {
		return id
	}
	return ""
}
