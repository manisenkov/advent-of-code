package registry

import (
	day2018_01 "github.com/manisenkov/advent-of-code/pkg/2018/solutions/day01"
	day2018_02 "github.com/manisenkov/advent-of-code/pkg/2018/solutions/day02"
	day2018_03 "github.com/manisenkov/advent-of-code/pkg/2018/solutions/day03"

	day2020_01 "github.com/manisenkov/advent-of-code/pkg/2020/solutions/day01"
	day2020_02 "github.com/manisenkov/advent-of-code/pkg/2020/solutions/day02"
)

func init() {
	Register(2018, 1, new(day2018_01.Solution))
	Register(2018, 2, new(day2018_02.Solution))
	Register(2018, 3, new(day2018_03.Solution))

	Register(2020, 1, new(day2020_01.Solution))
	Register(2020, 2, new(day2020_02.Solution))
}
