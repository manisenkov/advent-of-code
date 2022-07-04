package main

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
func (sol *Solution) Init(input []string) {
	sol.passports = make([]map[string]string, 0)
	currentPassport := make(map[string]string)
	newPassport := true
	for _, inp := range input {
		if inp == "" {
			sol.passports = append(sol.passports, currentPassport)
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
		sol.passports = append(sol.passports, currentPassport)
	}
}

// Part1 .
func (sol *Solution) Part1() any {
	res := 0
	for _, p := range sol.passports {
		_, cidPresent := p["cid"]
		if (len(p) == 7 && !cidPresent) || len(p) > 7 {
			res++
		}
	}
	return res
}

// Part2 .
func (sol *Solution) Part2() any {
	res := 0
	for _, p := range sol.passports {
		if validate(p) {
			res++
		}
	}
	return res
}

func validate(p map[string]string) bool {
	// Birth year
	byr, _ := strconv.Atoi(p["byr"])
	if byr < 1920 || byr > 2002 {
		return false
	}

	// Issue year
	iyr, _ := strconv.Atoi(p["iyr"])
	if iyr < 2010 || iyr > 2020 {
		return false
	}

	// Expiration year
	eyr, _ := strconv.Atoi(p["eyr"])
	if eyr < 2020 || eyr > 2030 {
		return false
	}

	// Height
	m := hgtRegex.FindAllStringSubmatch(p["hgt"], -1)
	if len(m) == 0 {
		return false
	}
	hgtType := m[0][2]
	nHgt, _ := strconv.Atoi(m[0][1])
	if (hgtType == "cm" && (nHgt < 150 || nHgt > 193)) ||
		(hgtType == "in" && (nHgt < 59 || nHgt > 76)) {
		return false
	}

	// Hair color
	if !hclRegex.MatchString(p["hcl"]) {
		return false
	}

	// Eye color
	if !eclRegex.MatchString(p["ecl"]) {
		return false
	}

	// Passport ID
	if !pidRegex.MatchString(p["pid"]) {
		return false
	}

	return true
}

func main() {
	common.Run(new(Solution))
}
