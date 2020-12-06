package day01

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

const testInput = `
1721
979
366
299
675
1456
`

type Day01TestSuite struct {
	suite.Suite
	sol *Solution
}

func (st *Day01TestSuite) SetupTest() {
	st.sol = new(Solution)
}

func (st *Day01TestSuite) Test1() {
	input := strings.Split(strings.Trim(testInput, " \n"), "\n")
	st.sol.Init(input)
	st.Equal(514579, st.sol.Part1())
	st.Equal(241861950, st.sol.Part2())
}

func TestDay01(t *testing.T) {
	st := new(Day01TestSuite)
	suite.Run(t, st)
}
