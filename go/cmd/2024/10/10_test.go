package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

const testInput = `
89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`

type Day10TestSuite struct {
	suite.Suite
	sol *Solution
}

func (st *Day10TestSuite) SetupTest() {
	st.sol = new(Solution)
}

func (st *Day10TestSuite) Test1() {
	input := strings.Split(strings.Trim(testInput, " \n"), "\n")
	st.sol.Init(input)
	st.Equal(36, st.sol.Part1())
	st.Equal(81, st.sol.Part2())
}

func TestDay10(t *testing.T) {
	st := new(Day10TestSuite)
	suite.Run(t, st)
}
