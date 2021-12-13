package main

import (
	"fmt"
	"strings"

	"github.com/manisenkov/advent-of-code/pkg/common"
)

type fold struct {
	direction string
	coord     int
}

type paper [][]bool

func fillPaper(dots map[[2]int]bool) paper {
	maxRow := -0x7FFFFFFF
	maxCol := -0x7FFFFFFF
	for dot := range dots {
		if dot[0] > maxCol {
			maxCol = dot[0]
		}
		if dot[1] > maxRow {
			maxRow = dot[1]
		}
	}
	res := make(paper, maxRow+1)
	for i := range res {
		res[i] = make([]bool, maxCol+1)
	}
	for dot := range dots {
		res[dot[1]][dot[0]] = true
	}
	return res
}

func (p paper) String() string {
	res := ""
	for _, row := range p {
		for _, point := range row {
			if point {
				res += "█"
			} else {
				res += " "
			}
		}
		res += "\n"
	}
	return res
}

// Solution contains solution for day 13
type Solution struct {
	dots  map[[2]int]bool
	folds []fold
}

// Init initializes solution with input data
func (sol *Solution) Init(input []string) {
	i := 0
	sol.dots = map[[2]int]bool{}
	for _, s := range input {
		if s == "" {
			break
		}
		xs := strings.Split(s, ",")
		sol.dots[[2]int{common.MustAtoi(xs[0]), common.MustAtoi(xs[1])}] = true
		i++
	}
	for _, s := range input[i+1:] {
		xs := strings.Split(string([]byte(s)[11:]), "=")
		sol.folds = append(sol.folds, fold{xs[0], common.MustAtoi(xs[1])})
	}
}

// Part1 .
func (sol *Solution) Part1() common.Any {
	res := foldPaper(sol.dots, sol.folds[0])
	return len(res)
}

// Part2 .
func (sol *Solution) Part2() common.Any {
	dots := sol.dots
	for _, f := range sol.folds {
		dots = foldPaper(dots, f)
	}
	p := fillPaper(dots)
	fmt.Println(p)
	return 0
}

func foldPaper(dots map[[2]int]bool, f fold) map[[2]int]bool {
	res := map[[2]int]bool{}
	if f.direction == "x" {
		for d := range dots {
			if d[0] < f.coord {
				res[d] = true
			} else {
				res[[2]int{2*f.coord - d[0], d[1]}] = true
			}
		}
	}
	if f.direction == "y" {
		for d := range dots {
			if d[1] < f.coord {
				res[d] = true
			} else {
				res[[2]int{d[0], 2*f.coord - d[1]}] = true
			}
		}
	}
	return res
}

func main() {
	common.Run(new(Solution))
}
