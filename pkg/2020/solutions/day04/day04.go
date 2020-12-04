package day04

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/manisenkov/advent-of-code/pkg/common"
)

var (
	hgtRegex = regexp.MustCompile(`^(\d+)(cm|in)$`)
	hclRegex = regexp.MustCompile(`^#[0-9a-f]{6}$`)
	eclRegex = regexp.MustCompile(`^(amb|blu|brn|gry|grn|hzl|oth)$`)
	pidRegex = regexp.MustCompile(`^\d{9}$`)
)

// Solution contains solution for day 4
type Solution struct {
	passports []map[string]string
}

// Init initializes solution with input data
func (s *Solution) Init(input []string) {
	passports := make([]map[string]string, 0)
	currentPassport := make(map[string]string)
	newPassport := true
	for _, inp := range input {
		if inp == "" {
			passports = append(passports, currentPassport)
			currentPassport = make(map[string]string)
			newPassport = true
			continue
		}
		pairs := strings.Split(inp, " ")
		for _, pair := range pairs {
			keyValue := strings.Split(pair, ":")
			currentPassport[keyValue[0]] = keyValue[1]
		}
		newPassport = false
	}
	if !newPassport {
		passports = append(passports, currentPassport)
	}
	s.passports = passports
}

// Part1 .
func (s *Solution) Part1() common.Any {
	res := 0
	for _, p := range s.passports {
		_, cidPresent := p["cid"]
		if (len(p) == 7 && !cidPresent) || len(p) > 7 {
			res++
		}
	}
	return res
}

// Part2 .
func (s *Solution) Part2() common.Any {
	res := 0
	for _, p := range s.passports {
		if validate(p) {
			res++
		}
	}
	return res
}

func validate(p map[string]string) bool {
	// Birth year
	byr, ok := p["byr"]
	if !ok {
		return false
	}
	nByr, err := strconv.Atoi(byr)
	if err != nil || nByr < 1920 || nByr > 2002 {
		return false
	}

	// Issue year
	iyr, ok := p["iyr"]
	if !ok {
		return false
	}
	nIyr, err := strconv.Atoi(iyr)
	if nIyr < 2010 || nIyr > 2020 {
		return false
	}

	// Expiration year
	eyr, ok := p["eyr"]
	if !ok {
		return false
	}
	nEyr, err := strconv.Atoi(eyr)
	if nEyr < 2020 || nEyr > 2030 {
		return false
	}

	// Height
	hgt, ok := p["hgt"]
	if !ok {
		return false
	}
	m := hgtRegex.FindAllStringSubmatch(hgt, -1)
	if len(m) == 0 {
		return false
	}
	nHgt, err := strconv.Atoi(m[0][1])
	hgtType := m[0][2]
	if (hgtType == "cm" && (nHgt < 150 || nHgt > 193)) ||
			(hgtType == "in" && (nHgt < 59 || nHgt > 76)) {
		return false
	}

	// Hair color
	hcl, ok := p["hcl"]
	if !ok {
		return false
	}
	if !hclRegex.MatchString(hcl) {
		return false
	}

	// Eye color
	ecl, ok := p["ecl"]
	if !ok {
		return false
	}
	if !eclRegex.MatchString(ecl) {
		return false
	}

	// Passport ID
	pid, ok := p["pid"]
	if !ok {
		return false
	}
	if !pidRegex.MatchString(pid) {
		return false
	}

	return true
}
