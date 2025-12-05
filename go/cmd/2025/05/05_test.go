package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

const testInput = `
3-5
10-14
16-20
12-18

1
5
8
11
17
32`

type Day05TestSuite struct {
	suite.Suite
	sol *Solution
}

func (st *Day05TestSuite) SetupTest() {
	st.sol = new(Solution)
}

func (st *Day05TestSuite) Test1() {
	input := strings.Split(strings.Trim(testInput, " \n"), "\n")
	st.sol.Init(input)
	st.Equal(3, st.sol.Part1())
	st.Equal(14, st.sol.Part2())
}

func TestDay05(t *testing.T) {
	st := new(Day05TestSuite)
	suite.Run(t, st)
}
