package day04

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type Day04TestSuite struct {
	suite.Suite
	sol *Solution
}

func (s *Day04TestSuite) SetupTest() {
	s.sol = new(Solution)
}

func (s *Day04TestSuite) Test1() {
	s.sol.Init([]string{
		"ecl:gry pid:860033327 eyr:2020 hcl:#fffffd",
		"byr:1937 iyr:2017 cid:147 hgt:183cm",
		"",
		"iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884",
		"hcl:#cfa07d byr:1929",
		"",
		"hcl:#ae17e1 iyr:2013",
		"eyr:2024",
		"ecl:brn pid:760753108 byr:1931",
		"hgt:179cm",
		"",
		"hcl:#cfa07d eyr:2025 pid:166559648",
		"iyr:2011 ecl:brn hgt:59in",
	})
	s.Equal(2, s.sol.Part1())
	s.Equal(0, s.sol.Part2())
}

func TestDay04(t *testing.T) {
	s := new(Day04TestSuite)
	suite.Run(t, s)
}
