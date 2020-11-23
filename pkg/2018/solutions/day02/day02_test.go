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
	s.sol.Init([]string{"abcdef", "bababc", "abbcde", "abcccd", "aabcdd", "abcdee", "ababab"})
	s.Equal(12, s.sol.Part1())
	s.Equal("abcde", s.sol.Part2())
}

func (s *Day02TestSuite) Test2() {
	s.sol.Init([]string{"abcde", "fghij", "klmno", "pqrst", "fguij", "axcye", "wvxyz"})
	s.Equal(0, s.sol.Part1())
	s.Equal("fgij", s.sol.Part2())
}

func TestDay02(t *testing.T) {
	s := new(Day02TestSuite)
	suite.Run(t, s)
}
