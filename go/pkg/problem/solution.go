package problem

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

// Solution represents a interface for the daily solutions of Advent of Code problems
type Solution interface {
	Init([]string)
	Part1() any
	Part2() any
}

// Solve executes the given solution, taking the input from the standard input.
// It outputs the initialization and the execution times for the both parts of the solution
func Solve(sol Solution) {
	input, err := ReadInput(os.Stdin)
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("Can't read input: %s", err))
		os.Exit(1)
	}

	startTime := time.Now()
	sol.Init(input)
	initTime := time.Now()
	res1 := sol.Part1()
	p1Time := time.Now()
	res2 := sol.Part2()
	p2Time := time.Now()
	fmt.Printf("Part 1: %v\n", res1)
	fmt.Printf("Part 2: %v\n", res2)
	fmt.Printf(" - Init time: %v\n", initTime.Sub(startTime))
	fmt.Printf(" - Part 1 time: %v\n", p1Time.Sub(initTime))
	fmt.Printf(" - Part 2 time: %v\n", p2Time.Sub(p1Time))
	fmt.Printf(" - Total time: %v\n", p2Time.Sub(startTime))
}

// ReadInput reads content of the given file and returns it as a slice of strings.
func ReadInput(f *os.File) ([]string, error) {
	scanner := bufio.NewScanner(f)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
