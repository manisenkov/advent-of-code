package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

const testInput = `
35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576
`

type Day09TestSuite struct {
	suite.Suite
	sol *Solution
}

func (st *Day09TestSuite) SetupTest() {
	st.sol = &Solution{
		preambleSize: 5,
	}
}

func (st *Day09TestSuite) Test1() {
	input := strings.Split(strings.Trim(testInput, " \n"), "\n")
	st.sol.Init(input)
	st.Equal(int64(127), st.sol.Part1())
	st.Equal(int64(62), st.sol.Part2())
}

func TestDay09(t *testing.T) {
	st := new(Day09TestSuite)
	suite.Run(t, st)
}
