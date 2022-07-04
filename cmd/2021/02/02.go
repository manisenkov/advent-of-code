package main

import (
	"strings"

	"github.com/manisenkov/advent-of-code/pkg/common"
)

type command struct {
	direction string
	arg       int
}

// Solution contains solution for day 2
type Solution struct {
	commands []command
}

// Init initializes solution with input data
func (sol *Solution) Init(input []string) {
	commands := make([]command, len(input))
	for i, s := range input {
		parts := strings.Split(s, " ")
		cmd := command{
			direction: parts[0],
			arg:       common.MustAtoi(parts[1]),
		}
		commands[i] = cmd
	}
	sol.commands = commands
}

// Part1 .
func (sol *Solution) Part1() any {
	h := 0
	d := 0
	for _, c := range sol.commands {
		switch c.direction {
		case "forward":
			h += c.arg
		case "down":
			d += c.arg
		case "up":
			d -= c.arg
		}
	}
	return h * d
}

// Part2 .
func (sol *Solution) Part2() any {
	h := 0
	d := 0
	aim := 0
	for _, c := range sol.commands {
		switch c.direction {
		case "forward":
			h += c.arg
			d += c.arg * aim
		case "down":
			aim += c.arg
		case "up":
			aim -= c.arg
		}
	}
	return h * d
}

func main() {
	common.Run(new(Solution))
}
