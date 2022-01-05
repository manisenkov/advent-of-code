package common

import (
	"bufio"
	"fmt"
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
	input, err := ReadInput(os.Stdin)
	if err != nil {
		fail(fmt.Sprintf("Can't read input: %s", err))
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
	fmt.Printf(" - Init time: %vμs\n", float64(initTime.Sub(startTime).Nanoseconds())/1000)
	fmt.Printf(" - Part 1 time: %vμs\n", float64(p1Time.Sub(initTime).Nanoseconds())/1000)
	fmt.Printf(" - Part 2 time: %vμs\n", float64(p2Time.Sub(p1Time).Nanoseconds())/1000)
	fmt.Printf(" - Total time: %vμs\n", float64(p2Time.Sub(startTime).Nanoseconds())/1000)
}

func ReadInput(f *os.File) ([]string, error) {
	scanner := bufio.NewScanner(f)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func fail(errorMessage string) {
	os.Stderr.WriteString(fmt.Sprintf("%v\n", errorMessage))
	os.Exit(1)
}
