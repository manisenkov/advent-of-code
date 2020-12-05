package day05

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type Day05TestSuite struct {
	suite.Suite
	sol *Solution
}

func (s *Day05TestSuite) SetupTest() {
	s.sol = new(Solution)
}

func (s *Day05TestSuite) Test1() {
	s.sol.Init([]string{
		"FBFBBFFRLR",
		"BFFFBBFRRR",
		"FFFBBBFRRR",
		"BBFFBBFRLL",
	})
	s.Equal(820, s.sol.Part1())
}

func TestDay05(t *testing.T) {
	s := new(Day05TestSuite)
	suite.Run(t, s)
}
