package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

const testInput1 = `
aaa: you hhh
you: bbb ccc
bbb: ddd eee
ccc: ddd eee fff
ddd: ggg
eee: out
fff: out
ggg: out
hhh: ccc fff iii
iii: out
`

const testInput2 = `
svr: aaa bbb
aaa: fft
fft: ccc
bbb: tty
tty: ccc
ccc: ddd eee
ddd: hub
hub: fff
eee: dac
dac: fff
fff: ggg hhh
ggg: out
hhh: out
`

type Day11TestSuite struct {
	suite.Suite
	sol *Solution
}

func (st *Day11TestSuite) SetupTest() {
	st.sol = new(Solution)
}

func (st *Day11TestSuite) Test1() {
	input := strings.Split(strings.Trim(testInput1, " \n"), "\n")
	st.sol.Init(input)
	st.Equal(5, st.sol.Part1())
}

func (st *Day11TestSuite) Test2() {
	input := strings.Split(strings.Trim(testInput2, " \n"), "\n")
	st.sol.Init(input)
	st.Equal(2, st.sol.Part2())
}

func TestDay11(t *testing.T) {
	st := new(Day11TestSuite)
	suite.Run(t, st)
}
