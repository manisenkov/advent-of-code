package common

import (
	"bufio"
	"fmt"
	"os"
)

// DaySolution represents a one AoC day solution
type DaySolution interface {
	Init([]string)
	Part1() Any
	Part2() Any
}

// Run day solution, taking input from stdin
func Run(sol DaySolution) {
	input, err := readInput()
	if err != nil {
		fail(fmt.Sprintf("Can't read input: %s", err))
	}

	sol.Init(input)

	res1 := sol.Part1()
	fmt.Printf("Part 1: %v\n", res1)

	res2 := sol.Part2()
	fmt.Printf("Part 2: %v\n", res2)
}

func fail(errorMessage string) {
	os.Stderr.WriteString(fmt.Sprintf("%v\n", errorMessage))
	os.Exit(1)
}

func readInput() ([]string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
