package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

const (
	testInput1 = `0,3,6`
	testInput2 = `1,3,2`
	testInput3 = `2,1,3`
)

type Day15TestSuite struct {
	suite.Suite
	sol *Solution
}

func (st *Day15TestSuite) SetupTest() {
	st.sol = new(Solution)
}

func (st *Day15TestSuite) Test1() {
	input := strings.Split(strings.Trim(testInput1, " \n"), "\n")
	st.sol.Init(input)
	st.Equal(436, st.sol.Part1())
	st.Equal(175594, st.sol.Part2())
}

func (st *Day15TestSuite) Test2() {
	input := strings.Split(strings.Trim(testInput2, " \n"), "\n")
	st.sol.Init(input)
	st.Equal(1, st.sol.Part1())
	st.Equal(2578, st.sol.Part2())
}

func (st *Day15TestSuite) Test3() {
	input := strings.Split(strings.Trim(testInput3, " \n"), "\n")
	st.sol.Init(input)
	st.Equal(10, st.sol.Part1())
	st.Equal(3544142, st.sol.Part2())
}

func TestDay15(t *testing.T) {
	st := new(Day15TestSuite)
	suite.Run(t, st)
}
