package main

import (
	"fmt"

	"github.com/manisenkov/advent-of-code/pkg/common"
)

type coord struct {
	x, y, z, w int
}

func (c coord) String() string {
	return fmt.Sprintf("(%v, %v, %v, %v)", c.x, c.y, c.z, c.w)
}

type space struct {
	active map[coord]bool
	xRange [2]int
	yRange [2]int
	zRange [2]int
	wRange [2]int
}

// Solution contains solution for day 17
type Solution struct {
	initSpace space
}

// Init initializes solution with input data
func (sol *Solution) Init(input []string) {
	active := make(map[coord]bool)
	h := len(input)
	w := len(input[0])
	for i, inp := range input {
		for j, c := range inp {
			if c == '#' {
				active[coord{j - w/2, i - h/2, 0, 0}] = true
			}
		}
	}
	sol.initSpace = space{
		active: active,
		xRange: [2]int{-w / 2, w / 2},
		yRange: [2]int{-h / 2, w / 2},
		zRange: [2]int{0, 0},
		wRange: [2]int{0, 0},
	}
}

// Part1 .
func (sol *Solution) Part1() any {
	curSpace := sol.initSpace
	for i := 0; i < 6; i++ {
		curSpace = simulate3d(curSpace)
	}
	return len(curSpace.active)
}

// Part2 .
func (sol *Solution) Part2() any {
	curSpace := sol.initSpace
	for i := 0; i < 6; i++ {
		curSpace = simulate4d(curSpace)
	}
	return len(curSpace.active)
}

func simulate3d(currentSpace space) space {
	res := space{
		active: make(map[coord]bool),
	}
	xRange := [2]int{0, 0}
	yRange := [2]int{0, 0}
	zRange := [2]int{0, 0}
	for x := currentSpace.xRange[0] - 1; x <= currentSpace.xRange[1]+1; x++ {
		for y := currentSpace.yRange[0] - 1; y <= currentSpace.yRange[1]+1; y++ {
			for z := currentSpace.zRange[0] - 1; z <= currentSpace.zRange[1]+1; z++ {
				c := coord{x, y, z, 0}
				n := calcNeighbors(currentSpace, c)
				if (currentSpace.active[c] && (n == 2 || n == 3)) || (!currentSpace.active[c] && n == 3) {
					res.active[c] = true
					if x < xRange[0] {
						xRange[0] = x
					}
					if x > xRange[1] {
						xRange[1] = x
					}
					if y < yRange[0] {
						yRange[0] = y
					}
					if y > yRange[1] {
						yRange[1] = y
					}
					if z < zRange[0] {
						zRange[0] = z
					}
					if z > zRange[1] {
						zRange[1] = z
					}
				}
			}
		}
	}
	res.xRange = xRange
	res.yRange = yRange
	res.zRange = zRange
	res.wRange = [2]int{0, 0}
	return res
}

func simulate4d(currentSpace space) space {
	res := space{
		active: make(map[coord]bool),
	}
	xRange := [2]int{0, 0}
	yRange := [2]int{0, 0}
	zRange := [2]int{0, 0}
	wRange := [2]int{0, 0}
	for x := currentSpace.xRange[0] - 1; x <= currentSpace.xRange[1]+1; x++ {
		for y := currentSpace.yRange[0] - 1; y <= currentSpace.yRange[1]+1; y++ {
			for z := currentSpace.zRange[0] - 1; z <= currentSpace.zRange[1]+1; z++ {
				for w := currentSpace.wRange[0] - 1; w <= currentSpace.wRange[1]+1; w++ {
					c := coord{x, y, z, w}
					n := calcNeighbors(currentSpace, c)
					if (currentSpace.active[c] && (n == 2 || n == 3)) || (!currentSpace.active[c] && n == 3) {
						res.active[c] = true
						if x < xRange[0] {
							xRange[0] = x
						}
						if x > xRange[1] {
							xRange[1] = x
						}
						if y < yRange[0] {
							yRange[0] = y
						}
						if y > yRange[1] {
							yRange[1] = y
						}
						if z < zRange[0] {
							zRange[0] = z
						}
						if z > zRange[1] {
							zRange[1] = z
						}
						if w < wRange[0] {
							wRange[0] = w
						}
						if w > wRange[1] {
							wRange[1] = w
						}
					}
				}
			}
		}
	}
	res.xRange = xRange
	res.yRange = yRange
	res.zRange = zRange
	res.wRange = wRange
	return res
}

func calcNeighbors(s space, c coord) int {
	res := 0
	for x := c.x - 1; x <= c.x+1; x++ {
		for y := c.y - 1; y <= c.y+1; y++ {
			for z := c.z - 1; z <= c.z+1; z++ {
				for w := c.w - 1; w <= c.w+1; w++ {
					if x == c.x && y == c.y && z == c.z && w == c.w {
						continue
					}
					if s.active[coord{x, y, z, w}] {
						res++
					}
				}
			}
		}
	}
	return res
}

func main() {
	common.Run(new(Solution))
}
