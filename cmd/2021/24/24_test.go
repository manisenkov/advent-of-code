package main

import (
	"os"
	"testing"

	"github.com/manisenkov/advent-of-code/pkg/common"
	"github.com/stretchr/testify/suite"
)

type Day24TestSuite struct {
	suite.Suite
	sol *Solution
}

func (st *Day24TestSuite) SetupTest() {
	st.sol = new(Solution)
}

func (st *Day24TestSuite) Test1() {
	f, err := os.Open("../../../inputs/2021/24.txt")
	if err != nil {
		panic(err)
	}
	input, err := common.ReadInput(f)
	if err != nil {
		panic(err)
	}
	st.sol.Init(input)
	st.Equal("99911993949684", st.sol.Part1())
	st.Equal("62911941716111", st.sol.Part2())
}

func TestDay24(t *testing.T) {
	st := new(Day24TestSuite)
	suite.Run(t, st)
}
