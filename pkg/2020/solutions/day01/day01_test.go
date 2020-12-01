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
	s.sol.Init([]string{
		"1721",
		"979",
		"366",
		"299",
		"675",
		"1456",
	})
	s.Equal(514579, s.sol.Part1())
	s.Equal(241861950, s.sol.Part2())
}

func TestDay01(t *testing.T) {
	s := new(Day01TestSuite)
	suite.Run(t, s)
}
