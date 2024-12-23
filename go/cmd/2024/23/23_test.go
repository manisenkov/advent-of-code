package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

const testInput = `
kh-tc
qp-kh
de-cg
ka-co
yn-aq
qp-ub
cg-tb
vc-aq
tb-ka
wh-tc
yn-cg
kh-ub
ta-co
de-co
tc-td
tb-wq
wh-td
ta-ka
td-qp
aq-cg
wq-ub
ub-vc
de-ta
wq-aq
wq-vc
wh-yn
ka-de
kh-ta
co-tc
wh-qp
tb-vc
td-yn`

type Day23TestSuite struct {
	suite.Suite
	sol *Solution
}

func (st *Day23TestSuite) SetupTest() {
	st.sol = new(Solution)
}

func (st *Day23TestSuite) Test1() {
	input := strings.Split(strings.Trim(testInput, " \n"), "\n")
	st.sol.Init(input)
	st.Equal(7, st.sol.Part1())
	// st.Equal("co,de,ka,ta", st.sol.Part2())
}

func TestDay23(t *testing.T) {
	st := new(Day23TestSuite)
	suite.Run(t, st)
}
