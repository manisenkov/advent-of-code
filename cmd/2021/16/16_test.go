package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

const testInputPart1_1 = `
8A004A801A8002F478
`

const testInputPart1_2 = `
620080001611562C8802118E34
`

const testInputPart1_3 = `
C0015000016115A2E0802F182340
`

const testInputPart1_4 = `
A0016C880162017C3686B18A3D4780
`

const testInputPart2_1 = `
C200B40A82
`

const testInputPart2_2 = `
04005AC33890
`

const testInputPart2_3 = `
880086C3E88112
`

const testInputPart2_4 = `
CE00C43D881120
`

const testInputPart2_5 = `
D8005AC2A8F0
`

const testInputPart2_6 = `
F600BC2D8F
`

const testInputPart2_7 = `
9C005AC2F8F0
`

const testInputPart2_8 = `
9C0141080250320F1802104A08
`

type Day16TestSuite struct {
	suite.Suite
	sol *Solution
}

func (st *Day16TestSuite) SetupTest() {
	st.sol = new(Solution)
}

func (st *Day16TestSuite) TestPart1_1() {
	input := strings.Split(strings.Trim(testInputPart1_1, " \n"), "\n")
	st.sol.Init(input)
	st.Equal(16, st.sol.Part1())
}

func (st *Day16TestSuite) TestPart1_2() {
	input := strings.Split(strings.Trim(testInputPart1_2, " \n"), "\n")
	st.sol.Init(input)
	st.Equal(12, st.sol.Part1())
}

func (st *Day16TestSuite) TestPart1_3() {
	input := strings.Split(strings.Trim(testInputPart1_3, " \n"), "\n")
	st.sol.Init(input)
	st.Equal(23, st.sol.Part1())
}

func (st *Day16TestSuite) TestPart1_4() {
	input := strings.Split(strings.Trim(testInputPart1_4, " \n"), "\n")
	st.sol.Init(input)
	st.Equal(31, st.sol.Part1())
}

func (st *Day16TestSuite) TestPart2_1() {
	input := strings.Split(strings.Trim(testInputPart2_1, " \n"), "\n")
	st.sol.Init(input)
	st.Equal(int64(3), st.sol.Part2())
}

func (st *Day16TestSuite) TestPart2_2() {
	input := strings.Split(strings.Trim(testInputPart2_2, " \n"), "\n")
	st.sol.Init(input)
	st.Equal(int64(54), st.sol.Part2())
}

func (st *Day16TestSuite) TestPart2_3() {
	input := strings.Split(strings.Trim(testInputPart2_3, " \n"), "\n")
	st.sol.Init(input)
	st.Equal(int64(7), st.sol.Part2())
}

func (st *Day16TestSuite) TestPart2_4() {
	input := strings.Split(strings.Trim(testInputPart2_4, " \n"), "\n")
	st.sol.Init(input)
	st.Equal(int64(9), st.sol.Part2())
}

func (st *Day16TestSuite) TestPart2_5() {
	input := strings.Split(strings.Trim(testInputPart2_5, " \n"), "\n")
	st.sol.Init(input)
	st.Equal(int64(1), st.sol.Part2())
}

func (st *Day16TestSuite) TestPart2_6() {
	input := strings.Split(strings.Trim(testInputPart2_6, " \n"), "\n")
	st.sol.Init(input)
	st.Equal(int64(0), st.sol.Part2())
}

func (st *Day16TestSuite) TestPart2_7() {
	input := strings.Split(strings.Trim(testInputPart2_7, " \n"), "\n")
	st.sol.Init(input)
	st.Equal(int64(0), st.sol.Part2())
}

func (st *Day16TestSuite) TestPart2_8() {
	input := strings.Split(strings.Trim(testInputPart2_8, " \n"), "\n")
	st.sol.Init(input)
	st.Equal(int64(1), st.sol.Part2())
}

func TestDay16(t *testing.T) {
	st := new(Day16TestSuite)
	suite.Run(t, st)
}
