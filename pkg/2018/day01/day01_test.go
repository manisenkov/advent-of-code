package day01

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type Day01TestSuite struct {
	suite.Suite
	sol *Solution
}

func (st *Day01TestSuite) SetupTest() {
	st.sol = new(Solution)
}

func (st *Day01TestSuite) Test1() {
	st.sol.Init([]string{"+1", "-2", "+3", "+1"})
	st.Equal(3, st.sol.Part1())
	st.Equal(2, st.sol.Part2())
}

func (st *Day01TestSuite) Test2() {
	st.sol.Init([]string{"+1", "+1", "+1"})
	st.Equal(3, st.sol.Part1())
	// No part 2 for this case
}

func (st *Day01TestSuite) Test3() {
	st.sol.Init([]string{"+1", "+1", "-2"})
	st.Equal(0, st.sol.Part1())
	st.Equal(0, st.sol.Part2())
}

func (st *Day01TestSuite) Test4() {
	st.sol.Init([]string{"-1", "-2", "-3"})
	st.Equal(-6, st.sol.Part1())
	// No part 2 for this case
}

func (st *Day01TestSuite) Test5() {
	st.sol.Init([]string{"+3", "+3", "+4", "-2", "-4"})
	st.Equal(4, st.sol.Part1())
	st.Equal(10, st.sol.Part2())
}

func (st *Day01TestSuite) Test6() {
	st.sol.Init([]string{"-6", "+3", "+8", "+5", "-6"})
	st.Equal(4, st.sol.Part1())
	st.Equal(5, st.sol.Part2())
}

func (st *Day01TestSuite) Test7() {
	st.sol.Init([]string{"+7", "+7", "-2", "-7", "-4"})
	st.Equal(1, st.sol.Part1())
	st.Equal(14, st.sol.Part2())
}

func TestDay01(t *testing.T) {
	st := new(Day01TestSuite)
	suite.Run(t, st)
}
