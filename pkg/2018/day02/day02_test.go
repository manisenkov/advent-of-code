package day02

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

const testInput1 = `
abcdef
bababc
abbcde
abcccd
aabcdd
abcdee
ababab
`

const testInput2 = `
abcde
fghij
klmno
pqrst
fguij
axcye
wvxyz
`

type Day02TestSuite struct {
	suite.Suite
	sol *Solution
}

func (st *Day02TestSuite) SetupTest() {
	st.sol = new(Solution)
}

func (st *Day02TestSuite) Test1() {
	input := strings.Split(strings.Trim(testInput1, " \n"), "\n")
	st.sol.Init(input)
	st.Equal(12, st.sol.Part1())
	st.Equal("abcde", st.sol.Part2())
}

func (st *Day02TestSuite) Test2() {
	input := strings.Split(strings.Trim(testInput2, " \n"), "\n")
	st.sol.Init(input)
	st.Equal(0, st.sol.Part1())
	st.Equal("fgij", st.sol.Part2())
}

func TestDay02(t *testing.T) {
	st := new(Day02TestSuite)
	suite.Run(t, st)
}
