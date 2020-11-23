package registry

import (
	"github.com/manisenkov/advent-of-code/pkg/2018/solutions/day01"
	"github.com/manisenkov/advent-of-code/pkg/2018/solutions/day02"
	"github.com/manisenkov/advent-of-code/pkg/2018/solutions/day03"
)

func init() {
	Register(2018, 1, new(day01.Solution))
	Register(2018, 2, new(day02.Solution))
	Register(2018, 3, new(day03.Solution))
}
