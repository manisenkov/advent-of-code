package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/manisenkov/advent-of-code/pkg/registry"
)

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

func main() {
	if len(os.Args) != 3 {
		fail("Please give a year and a day number")
	}

	year, err := strconv.Atoi(os.Args[1])
	if err != nil && year >= 2010 && year <= 2100 {
		fail("Year should be a 4-digit integer")
	}

	day, err := strconv.Atoi(os.Args[2])
	if err != nil || day < 1 || day > 25 {
		fail("Day number should be between 1 and 25")
	}

	sol, ok := registry.Get(year, day)
	if !ok {
		fail(fmt.Sprintf("There's no solution for day %v", day))
	}

	input, err := readInput()
	if err != nil {
		fail(fmt.Sprintf("Can't read input: %s", err))
	}

	if err := sol.Init(input); err != nil {
		fail(fmt.Sprintf("Can't initialize solution: %s", err))
	}

	res1 := sol.Part1()
	fmt.Printf("Part 1: %v\n", res1)

	res2 := sol.Part2()
	fmt.Printf("Part 2: %v\n", res2)
}
