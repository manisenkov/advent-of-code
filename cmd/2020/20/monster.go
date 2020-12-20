package main

var monsterPattern = [][]rune{
	[]rune("                  # "),
	[]rune("#    ##    ##    ###"),
	[]rune(" #  #  #  #  #  #   "),
}

type monsterType struct {
	width  int
	height int
	cells  map[[2]int]bool
}

func (m *monsterType) match(grid [][]rune, x, y int) bool {
	for coord := range m.cells {
		dy := coord[0]
		dx := coord[1]
		if grid[y+dy][x+dx] != '#' {
			return false
		}
	}
	return true
}

func (m *monsterType) numMonsters(grid [][]rune) int {
	res := 0
	for y := 0; y < len(grid)-m.height; y++ {
		for x := 0; x < len(grid[0])-m.width; x++ {
			if m.match(grid, x, y) {
				res++
			}
		}
	}
	return res
}

func (m *monsterType) size() int {
	return len(m.cells)
}

func parseMonster() monsterType {
	m := monsterType{
		width:  len(monsterPattern[0]),
		height: len(monsterPattern),
		cells:  make(map[[2]int]bool),
	}
	for i, row := range monsterPattern {
		for j, c := range row {
			if c == '#' {
				m.cells[[2]int{i, j}] = true
			}
		}
	}
	return m
}
