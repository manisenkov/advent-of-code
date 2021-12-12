package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

const testInput1 = `
start-A
start-b
A-c
A-b
b-d
A-end
b-end
`

const testInput2 = `
dc-end
HN-start
start-kj
dc-start
dc-HN
LN-dc
HN-end
kj-sa
kj-HN
kj-dc
`

const testInput3 = `
fs-end
he-DX
fs-he
start-DX
pj-DX
end-zg
zg-sl
zg-pj
pj-he
RW-he
fs-DX
pj-RW
zg-RW
start-pj
he-WI
zg-he
pj-fs
start-RW
`

type Day12TestSuite struct {
	suite.Suite
	sol *Solution
}

func (st *Day12TestSuite) SetupTest() {
	st.sol = new(Solution)
}

func (st *Day12TestSuite) Test1() {
	input := strings.Split(strings.Trim(testInput1, " \n"), "\n")
	st.sol.Init(input)
	st.Equal(10, st.sol.Part1())
	st.Equal(36, st.sol.Part2())
}

func (st *Day12TestSuite) Test2() {
	input := strings.Split(strings.Trim(testInput2, " \n"), "\n")
	st.sol.Init(input)
	st.Equal(19, st.sol.Part1())
	st.Equal(103, st.sol.Part2())
}

func (st *Day12TestSuite) Test3() {
	input := strings.Split(strings.Trim(testInput3, " \n"), "\n")
	st.sol.Init(input)
	st.Equal(226, st.sol.Part1())
	st.Equal(3509, st.sol.Part2())
}

func TestDay12(t *testing.T) {
	st := new(Day12TestSuite)
	suite.Run(t, st)
}
