package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

const testInput = `
Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400

Button A: X+26, Y+66
Button B: X+67, Y+21
Prize: X=12748, Y=12176

Button A: X+17, Y+86
Button B: X+84, Y+37
Prize: X=7870, Y=6450

Button A: X+69, Y+23
Button B: X+27, Y+71
Prize: X=18641, Y=10279`

type Day13TestSuite struct {
	suite.Suite
	sol *Solution
}

func (st *Day13TestSuite) SetupTest() {
	st.sol = new(Solution)
}

func (st *Day13TestSuite) Test1() {
	input := strings.Split(strings.Trim(testInput, " \n"), "\n")
	st.sol.Init(input)
	st.Equal(480, st.sol.Part1())
	st.Equal(875318608908, st.sol.Part2())
}

func TestDay13(t *testing.T) {
	st := new(Day13TestSuite)
	suite.Run(t, st)
}
