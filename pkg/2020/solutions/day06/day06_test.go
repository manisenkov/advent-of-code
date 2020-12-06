package day06

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type Day06TestSuite struct {
	suite.Suite
	sol *Solution
}

func (s *Day06TestSuite) SetupTest() {
	s.sol = new(Solution)
}

func (s *Day06TestSuite) Test1() {
	s.sol.Init([]string{
		"abc",
		"",
		"a",
		"b",
		"c",
		"",
		"ab",
		"ac",
		"",
		"a",
		"a",
		"a",
		"a",
		"",
		"b",
	})
	s.Equal(11, s.sol.Part1())
	s.Equal(6, s.sol.Part2())
}

func TestDay06(t *testing.T) {
	s := new(Day06TestSuite)
	suite.Run(t, s)
}
