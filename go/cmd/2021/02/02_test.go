package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

const testInput = `
forward 5
down 5
forward 8
up 3
down 8
forward 2
`

type Day02TestSuite struct {
	suite.Suite
	sol *Solution
}

func (st *Day02TestSuite) SetupTest() {
	st.sol = new(Solution)
}

func (st *Day02TestSuite) Test1() {
	input := strings.Split(strings.Trim(testInput, " \n"), "\n")
	st.sol.Init(input)
	st.Equal(150, st.sol.Part1())
	st.Equal(900, st.sol.Part2())
}

func TestDay02(t *testing.T) {
	st := new(Day02TestSuite)
	suite.Run(t, st)
}
