package main

import (
	"strings"

	"github.com/manisenkov/advent-of-code/pkg/common"
)

// Solution contains solution for day 19
type Solution struct {
	initBeaconPos []map[[3]int]bool
	scanners      [][3]int
}

// Init initializes solution with input data
func (sol *Solution) Init(input []string) {
	sol.initBeaconPos = make([]map[[3]int]bool, 0)
	sol.scanners = [][3]int{}
	var curScanner map[[3]int]bool
	for _, s := range append(input, "") {
		if strings.HasPrefix(s, "---") {
			curScanner = make(map[[3]int]bool)
			continue
		}
		if s == "" {
			sol.initBeaconPos = append(sol.initBeaconPos, curScanner)
			continue
		}
		xs := strings.Split(s, ",")
		curScanner[[3]int{
			common.MustAtoi(xs[0]),
			common.MustAtoi(xs[1]),
			common.MustAtoi(xs[2]),
		}] = true
	}
}

// Part1 .
func (sol *Solution) Part1() any {
	indexes := [][]int{{0, 1, 2}, {0, 2, 1}, {1, 0, 2}, {1, 2, 0}, {2, 1, 0}, {2, 0, 1}}
	signs := [][]int{{-1, -1, -1}, {-1, -1, 1}, {-1, 1, -1}, {-1, 1, 1}, {1, -1, -1}, {1, -1, 1}, {1, 1, -1}, {1, 1, 1}}
	beacons := map[[3]int]bool{}
	for pos := range sol.initBeaconPos[0] {
		beacons[pos] = true
	}
	unmatched := make([]map[[3]int]bool, len(sol.initBeaconPos)-1)
	copy(unmatched, sol.initBeaconPos[1:])
	for len(unmatched) > 0 {
		ibp := unmatched[0]
		unmatched = unmatched[1:]
		matched := false
		for _, idx := range indexes {
			for _, sgn := range signs {
				r := rot(ibp, idx, sgn)
				matchedBeacons, scannerPos := match(beacons, r)
				if matchedBeacons != nil {
					matched = true
					for pos := range matchedBeacons {
						beacons[pos] = true
					}
					sol.scanners = append(sol.scanners, scannerPos)
					break
				}
			}
			if matched {
				break
			}
		}
		if !matched {
			unmatched = append(unmatched, ibp)
		}
	}

	return len(beacons)
}

// Part2 .
func (sol *Solution) Part2() any {
	maxDist := 0
	for i, scanner1 := range sol.scanners {
		for _, scanner2 := range sol.scanners[i:] {
			dist := common.AbsInt(scanner1[0]-scanner2[0]) + common.AbsInt(scanner1[1]-scanner2[1]) + common.AbsInt(scanner1[2]-scanner2[2])
			if dist > maxDist {
				maxDist = dist
			}
		}
	}
	return maxDist
}

func intersect(orig, target map[[3]int]bool) map[[3]int]bool {
	res := map[[3]int]bool{}
	for v := range orig {
		if target[v] {
			res[v] = true
		}
	}
	return res
}

func match(orig, target map[[3]int]bool) (map[[3]int]bool, [3]int) {
	for v1 := range orig {
		for v2 := range target {
			dv := [3]int{v1[0] - v2[0], v1[1] - v2[1], v1[2] - v2[2]}
			movedTarget := trans(target, dv)
			is := intersect(orig, movedTarget)
			if len(is) >= 12 {
				return movedTarget, dv
			}
		}
	}
	return nil, [3]int{0, 0, 0}
}

func rot(target map[[3]int]bool, indexes, signs []int) map[[3]int]bool {
	res := map[[3]int]bool{}
	for pos := range target {
		res[[3]int{pos[indexes[0]] * signs[0], pos[indexes[1]] * signs[1], pos[indexes[2]] * signs[2]}] = true
	}
	return res
}

func trans(target map[[3]int]bool, vector [3]int) map[[3]int]bool {
	res := map[[3]int]bool{}
	for pos := range target {
		res[[3]int{pos[0] + vector[0], pos[1] + vector[1], pos[2] + vector[2]}] = true
	}
	return res
}

func main() {
	common.Run(new(Solution))
}
