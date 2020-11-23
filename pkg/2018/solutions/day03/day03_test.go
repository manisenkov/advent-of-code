package day03

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type Day03TestSuite struct {
	suite.Suite
	sol *Solution
}

func (s *Day03TestSuite) SetupTest() {
	s.sol = new(Solution)
}

func (s *Day03TestSuite) Test1() {
	s.sol.Init([]string{
		"#1 @ 1,3: 4x4",
		"#2 @ 3,1: 4x4",
		"#3 @ 5,5: 2x2",
	})
	s.Equal(4, s.sol.Part1())
	s.Equal("3", s.sol.Part2())
}

func TestDay03(t *testing.T) {
	s := new(Day03TestSuite)
	suite.Run(t, s)
}
