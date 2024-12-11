package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

const testInput = `125 17`

type Day11TestSuite struct {
	suite.Suite
	sol *Solution
}

func (st *Day11TestSuite) SetupTest() {
	st.sol = new(Solution)
}

func (st *Day11TestSuite) Test1() {
	input := strings.Split(strings.Trim(testInput, " \n"), "\n")
	st.sol.Init(input)
	st.Equal(55312, st.sol.Part1())
	st.Equal(65601038650482, st.sol.Part2())
}

func TestDay11(t *testing.T) {
	st := new(Day11TestSuite)
	suite.Run(t, st)
}
