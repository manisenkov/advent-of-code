package day02

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type Day02TestSuite struct {
	suite.Suite
	sol *Solution
}

func (s *Day02TestSuite) SetupTest() {
	s.sol = new(Solution)
}

func (s *Day02TestSuite) Test1() {
	s.sol.Init([]string{
		"1-3 a: abcde",
		"1-3 b: cdefg",
		"2-9 c: ccccccccc",
	})
	s.Equal(2, s.sol.Part1())
	s.Equal(1, s.sol.Part2())
}

func TestDay02(t *testing.T) {
	s := new(Day02TestSuite)
	suite.Run(t, s)
}
