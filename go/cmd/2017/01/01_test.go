package main

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type testCase struct {
	input string
	part1 int
	part2 int
}

var cases []testCase = []testCase{
	{input: "1122", part1: 3, part2: 0},
	{input: "1111", part1: 4, part2: 4},
	{input: "1234", part1: 0, part2: 0},
	{input: "91212129", part1: 9, part2: 6},
	{input: "1212", part1: 0, part2: 6},
	{input: "1221", part1: 3, part2: 0},
	{input: "123123", part1: 0, part2: 12},
	{input: "12131415", part1: 0, part2: 4},
}

type Day01TestSuite struct {
	suite.Suite
	sol *Solution
}

func (st *Day01TestSuite) SetupTest() {
	st.sol = new(Solution)
}

func (st *Day01TestSuite) TestSolution() {
	for _, c := range cases {
		st.sol.Init([]string{c.input})
		st.Equal(c.part1, st.sol.Part1())
		st.Equal(c.part2, st.sol.Part2())
	}
}

func TestDay01(t *testing.T) {
	st := new(Day01TestSuite)
	suite.Run(t, st)
}
