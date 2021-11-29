package common

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"time"
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

	startTime := time.Now()
	sol.Init(input)
	res1 := sol.Part1()
	p1Time := time.Now()
	res2 := sol.Part2()
	p2Time := time.Now()
	fmt.Printf("Part 1: %v\n", res1)
	fmt.Printf("Part 2: %v\n", res2)
	fmt.Printf(" - Part 1 time: %vms\n", math.Round(float64(p1Time.Sub(startTime).Nanoseconds()) / 1000) / 1000)
	fmt.Printf(" - Part 2 time: %vms\n", math.Round(float64(p2Time.Sub(p1Time).Nanoseconds()) / 1000) / 1000)
	fmt.Printf(" - Total time: %vms\n", math.Round(float64(p2Time.Sub(startTime).Nanoseconds()) / 1000) / 1000)
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
