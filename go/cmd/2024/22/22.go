package main

import (
	"maps"
	"sync"

	"github.com/manisenkov/advent-of-code/pkg/collections"
	"github.com/manisenkov/advent-of-code/pkg/numbers"
	"github.com/manisenkov/advent-of-code/pkg/problem"
)

const prun int64 = (1 << 24) - 1 // 16777216 - 1

func gen(r int64) int64 {
	r = ((r << 6) ^ r) & prun
	r = ((r >> 5) ^ r) & prun
	r = ((r << 11) ^ r) & prun
	return r
}

// Solution contains a solution for day 22
type Solution struct {
	initialSecrets []int64
	table          [][2001]int64
}

// Init initializes the solution with the input data
func (sol *Solution) Init(input []string) {
	sol.initialSecrets = collections.MapTo(input, numbers.MustAtoi[int64])
}

// Part1 to solve a "silver" part (for a first star)
func (sol *Solution) Part1() any {
	res := int64(0)
	sol.table = make([][2001]int64, len(sol.initialSecrets))
	for i, secret := range sol.initialSecrets {
		r := secret
		sol.table[i][0] = r
		for j := 1; j < 2001; j++ {
			r = gen(r)
			sol.table[i][j] = r % 10
		}
		res += r
	}
	return res
}

// Part2 to solve a "gold" part (for a second star)
func (sol *Solution) Part2() any {
	diffTable := make([]map[[4]int64]int64, len(sol.table))
	for i, col := range sol.table {
		diffTable[i] = make(map[[4]int64]int64)
		for j := 4; j < 2001; j++ {
			c := [5]int64{col[j-4], col[j-3], col[j-2], col[j-1], col[j]}
			d := [4]int64{c[1] - c[0], c[2] - c[1], c[3] - c[2], c[4] - c[3]}
			if _, ok := diffTable[i][d]; !ok {
				diffTable[i][d] = c[4]
			}
		}
	}
	uniqueDiffs := [][4]int64{}
	for _, diffs := range diffTable {
		uniqueDiffs = append(uniqueDiffs, collections.IterToSlice(maps.Keys(diffs))...)
	}
	uniqueDiffs = collections.Unique(uniqueDiffs)
	ch := make(chan int64)
	wg := new(sync.WaitGroup)
	wg.Add(len(uniqueDiffs))
	for _, diff := range uniqueDiffs {
		go func(diff [4]int64) {
			s := int64(0)
			for _, dt := range diffTable {
				s += dt[diff]
			}
			ch <- s
			wg.Done()
		}(diff)
	}
	res := int64(0)
	go func() {
		for s := range ch {
			if s > res {
				res = s
			}
		}
	}()
	wg.Wait()
	return res
}

func main() {
	problem.Solve(new(Solution))
}
