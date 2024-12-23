package main

import (
	"slices"
	"strings"

	"github.com/manisenkov/advent-of-code/pkg/collections"
	"github.com/manisenkov/advent-of-code/pkg/problem"
)

type groupKey string

func (k groupKey) items() []string {
	return strings.Split(string(k), ",")
}

func (k groupKey) append(item string) groupKey {
	items := k.items()
	items = append(items, item)
	slices.Sort(items)
	return groupKey(strings.Join(items, ","))
}

// Solution contains a solution for day 23
type Solution struct {
	connections map[string]collections.Set[string]
}

// Init initializes the solution with the input data
func (sol *Solution) Init(input []string) {
	sol.connections = make(map[string]collections.Set[string])
	for _, line := range input {
		parts := strings.Split(line, "-")
		if c, ok := sol.connections[parts[0]]; ok {
			c[parts[1]] = true
		} else {
			sol.connections[parts[0]] = collections.Set[string]{parts[1]: true}
		}
		if c, ok := sol.connections[parts[1]]; ok {
			c[parts[0]] = true
		} else {
			sol.connections[parts[1]] = collections.Set[string]{parts[0]: true}
		}
	}
}

func (sol *Solution) findInterconnectedGroups(curGroups collections.Set[groupKey]) collections.Set[groupKey] {
	res := make(collections.Set[groupKey])
	for grp := range curGroups {
		items := grp.items()
		connections := make(collections.Set[string])
		for _, it := range items {
			for c := range sol.connections[it] {
				connections[c] = true
			}
		}
		for c := range connections {
			if collections.All(items, func(t string) bool {
				return sol.connections[c][t]
			}) {
				newGroup := grp.append(c)
				res[newGroup] = true
			}
		}
	}
	return res
}

// Part1 to solve a "silver" part (for a first star)
func (sol *Solution) Part1() any {
	groups := make(collections.Set[groupKey])
	for c := range sol.connections {
		groups[groupKey(c)] = true
	}
	for i := 0; i < 2; i++ {
		groups = sol.findInterconnectedGroups(groups)
	}
	res := 0
	for grp := range groups {
		items := grp.items()
		if collections.Any(items, func(t string) bool {
			return t[0] == 't'
		}) {
			res++
		}
	}
	return res
}

// Part2 to solve a "gold" part (for a second star)
func (sol *Solution) Part2() any {
	groups := make(collections.Set[groupKey])
	for c := range sol.connections {
		groups[groupKey(c)] = true
	}
	for {
		newGroups := sol.findInterconnectedGroups(groups)
		if len(newGroups) == 0 {
			break
		}
		groups = newGroups
	}
	return collections.FirstOfSet(groups)
}

func main() {
	problem.Solve(new(Solution))
}
