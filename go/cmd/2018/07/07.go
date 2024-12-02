package main

import (
	"sort"

	"github.com/manisenkov/advent-of-code/pkg/problem"
)

// Solution contains solution for day 7
type Solution struct {
	edges        map[byte][]byte
	requirements map[byte][]byte
	vertices     map[byte]bool
}

// Init initializes solution with input data
func (sol *Solution) Init(input []string) {
	edges := map[byte][]byte{}
	requirements := map[byte][]byte{}
	vertices := map[byte]bool{}
	for _, s := range input {
		from := []byte(s)[5]
		to := []byte(s)[36]
		edges[from] = append(edges[from], to)
		requirements[to] = append(requirements[to], from)
		vertices[from] = true
		vertices[to] = true
	}
	sol.edges = edges
	sol.requirements = requirements
	sol.vertices = vertices
}

// Part1 .
func (sol *Solution) Part1() any {
	// Find start vertex
	verticesLeft := map[byte]bool{}
	for v := range sol.vertices {
		verticesLeft[v] = true
	}
	for _, dest := range sol.edges {
		for _, to := range dest {
			if ok := verticesLeft[to]; ok {
				delete(verticesLeft, to)
			}
		}
	}
	queue := []byte{}
	for v := range verticesLeft {
		queue = append(queue, v)
	}
	sort.Slice(queue, func(i, j int) bool { return queue[i] < queue[j] })

	// Process
	res := ""
	processed := map[byte]bool{}
	for len(queue) > 0 {
		from := queue[0]
		queue = queue[1:]
		if processed[from] {
			continue
		}

		// Check if requirements satisfied
		reqsSatisfied := true
		for _, req := range sol.requirements[from] {
			if !processed[req] {
				reqsSatisfied = false
				break
			}
		}

		if reqsSatisfied {
			processed[from] = true
			res += string([]byte{from})
			queue = append(queue, sol.edges[from]...)
			sort.Slice(queue, func(i, j int) bool { return queue[i] < queue[j] })
		} else {
			queue = append(queue, from)
		}
	}
	return res
}

// Part2 .
func (sol *Solution) Part2() any {
	return 0
}

func main() {
	problem.Solve(new(Solution))
}
