package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

const testInput = "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"

type Day02TestSuite struct {
	suite.Suite
	sol *Solution
}

func (st *Day02TestSuite) SetupTest() {
	st.sol = new(Solution)
}

func (st *Day02TestSuite) Test1() {
	input := strings.Split(strings.Trim(testInput, " \n"), "\n")
	st.sol.Init(input)
	st.Equal(1227775554, st.sol.Part1())
	st.Equal(4174379265, st.sol.Part2())
}

func TestDay02(t *testing.T) {
	st := new(Day02TestSuite)
	suite.Run(t, st)
}
