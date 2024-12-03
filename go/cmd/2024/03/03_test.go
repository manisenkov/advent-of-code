package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

const testInput = `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`

type Day03TestSuite struct {
	suite.Suite
	sol *Solution
}

func (st *Day03TestSuite) SetupTest() {
	st.sol = new(Solution)
}

func (st *Day03TestSuite) Test1() {
	input := strings.Split(strings.Trim(testInput, " \n"), "\n")
	st.sol.Init(input)
	st.Equal(161, st.sol.Part1())
	st.Equal(48, st.sol.Part2())
}

func TestDay03(t *testing.T) {
	st := new(Day03TestSuite)
	suite.Run(t, st)
}
