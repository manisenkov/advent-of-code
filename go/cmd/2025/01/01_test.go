package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

const testInput = `
L68
L30
R48
L5
R60
L55
L1
L99
R14
L82
`

type Day01TestSuite struct {
	suite.Suite
	sol *Solution
}

func (st *Day01TestSuite) SetupTest() {
	st.sol = new(Solution)
}

func (st *Day01TestSuite) Test1() {
	input := strings.Split(strings.Trim(testInput, " \n"), "\n")
	st.sol.Init(input)
	st.Equal(3, st.sol.Part1())
	st.Equal(6, st.sol.Part2())
}

func TestDay01(t *testing.T) {
	st := new(Day01TestSuite)
	suite.Run(t, st)
}
