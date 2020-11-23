package day01

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type Day01TestSuite struct {
	suite.Suite
	sol *Solution
}

func (s *Day01TestSuite) SetupTest() {
	s.sol = new(Solution)
}

func (s *Day01TestSuite) Test1() {
	s.sol.Init([]string{"+1", "-2", "+3", "+1"})
	s.Equal(3, s.sol.Part1())
	s.Equal(2, s.sol.Part2())
}

func (s *Day01TestSuite) Test2() {
	s.sol.Init([]string{"+1", "+1", "+1"})
	s.Equal(3, s.sol.Part1())
	// No part 2 for this case
}

func (s *Day01TestSuite) Test3() {
	s.sol.Init([]string{"+1", "+1", "-2"})
	s.Equal(0, s.sol.Part1())
	s.Equal(0, s.sol.Part2())
}

func (s *Day01TestSuite) Test4() {
	s.sol.Init([]string{"-1", "-2", "-3"})
	s.Equal(-6, s.sol.Part1())
	// No part 2 for this case
}

func (s *Day01TestSuite) Test5() {
	s.sol.Init([]string{"+3", "+3", "+4", "-2", "-4"})
	s.Equal(4, s.sol.Part1())
	s.Equal(10, s.sol.Part2())
}

func (s *Day01TestSuite) Test6() {
	s.sol.Init([]string{"-6", "+3", "+8", "+5", "-6"})
	s.Equal(4, s.sol.Part1())
	s.Equal(5, s.sol.Part2())
}

func (s *Day01TestSuite) Test7() {
	s.sol.Init([]string{"+7", "+7", "-2", "-7", "-4"})
	s.Equal(1, s.sol.Part1())
	s.Equal(14, s.sol.Part2())
}

func TestDay01(t *testing.T) {
	s := new(Day01TestSuite)
	suite.Run(t, s)
}
