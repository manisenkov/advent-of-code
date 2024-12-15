package main

import (
	"math/big"
	"strings"

	"github.com/manisenkov/advent-of-code/pkg/collections"
	"github.com/manisenkov/advent-of-code/pkg/numbers"
	"github.com/manisenkov/advent-of-code/pkg/problem"
	"github.com/manisenkov/advent-of-code/pkg/rmat"
	"github.com/manisenkov/advent-of-code/pkg/vec"
)

const (
	AREA_MIN = 200000000000000
	AREA_MAX = 400000000000000
)

type Halestone struct {
	pos vec.Vec[int64]
	vel vec.Vec[int64]
}

// Solution contains a solution for day 24
type Solution struct {
	halestones []Halestone
	area       [2]int64
}

// Init initializes the solution with the input data
func (sol *Solution) Init(input []string) {
	halestones := make([]Halestone, len(input))
	for i, s := range input {
		parts := strings.Split(s, " @ ")
		posParts := strings.Split(parts[0], ", ")
		velParts := strings.Split(parts[1], ", ")
		halestones[i] = Halestone{
			pos: vec.New(collections.MapTo(posParts, numbers.MustAtoi[int64])),
			vel: vec.New(collections.MapTo(velParts, numbers.MustAtoi[int64])),
		}
	}
	sol.halestones = halestones
	sol.area = [2]int64{AREA_MIN, AREA_MAX}
}

// Part1 to solve a "silver" part (for a first star)
func (sol *Solution) Part1() any {
	factors := make([][2]*big.Rat, len(sol.halestones))
	for i, hs := range sol.halestones {
		pos := [2]vec.Vec[int64]{hs.pos, hs.pos.Add(hs.vel)}
		factors[i] = [2]*big.Rat{
			big.NewRat(
				pos[1].Y()-pos[0].Y(),
				pos[1].X()-pos[0].X()),
			big.NewRat(
				pos[0].Y()*pos[1].X()-pos[0].X()*pos[1].Y(),
				pos[1].X()-pos[0].X()),
		}
	}
	reduced := collections.MapTo(sol.halestones, func(hs Halestone) Halestone {
		return Halestone{
			pos: hs.pos.Reduce(2),
			vel: hs.vel.Reduce(2),
		}
	})
	res := 0
	for i, hs0 := range reduced[:len(reduced)-1] {
		for j := i + 1; j < len(reduced); j++ {
			hs1 := reduced[j]
			if numbers.RatEqual(numbers.RatSub(factors[i][0], factors[j][0]), numbers.RatZero()) {
				continue
			}
			crossX := numbers.RatQuo(numbers.RatSub(factors[j][1], factors[i][1]), numbers.RatSub(factors[i][0], factors[j][0]))
			crossY := numbers.RatAdd(numbers.RatMul(factors[i][0], crossX), factors[i][1])
			fCrossX, _ := crossX.Float64()
			fCrossY, _ := crossY.Float64()
			cross := vec.New([]int64{int64(fCrossX), int64(fCrossY)})

			crossDir0 := cross.Sub(hs0.pos)
			crossDir1 := cross.Sub(hs1.pos)
			if crossDir0.IsSimilar(hs0.vel) &&
				crossDir1.IsSimilar(hs1.vel) &&
				cross.X() > sol.area[0] &&
				cross.X() < sol.area[1] &&
				cross.Y() > sol.area[0] &&
				cross.Y() < sol.area[1] {
				res += 1
			}
		}
	}
	return res
}

// Part2 to solve a "gold" part (for a second star)
func (sol *Solution) Part2() any {
	hs0 := sol.halestones[0]
	hs1 := sol.halestones[1]
	hs2 := sol.halestones[2]
	a := rmat.FromIntTable([][]int64{
		{hs1.vel.Y() - hs0.vel.Y(), hs0.vel.X() - hs1.vel.X(), 0, hs0.pos.Y() - hs1.pos.Y(), hs1.pos.X() - hs0.pos.X(), 0},
		{hs2.vel.Y() - hs0.vel.Y(), hs0.vel.X() - hs2.vel.X(), 0, hs0.pos.Y() - hs2.pos.Y(), hs2.pos.X() - hs0.pos.X(), 0},
		{hs1.vel.Z() - hs0.vel.Z(), 0, hs0.vel.X() - hs1.vel.X(), hs0.pos.Z() - hs1.pos.Z(), 0, hs1.pos.X() - hs0.pos.X()},
		{hs2.vel.Z() - hs0.vel.Z(), 0, hs0.vel.X() - hs2.vel.X(), hs0.pos.Z() - hs2.pos.Z(), 0, hs2.pos.X() - hs0.pos.X()},
		{0, hs1.vel.Z() - hs0.vel.Z(), hs0.vel.Y() - hs1.vel.Y(), 0, hs0.pos.Z() - hs1.pos.Z(), hs1.pos.Y() - hs0.pos.Y()},
		{0, hs2.vel.Z() - hs0.vel.Z(), hs0.vel.Y() - hs2.vel.Y(), 0, hs0.pos.Z() - hs2.pos.Z(), hs2.pos.Y() - hs0.pos.Y()},
	})
	b := rmat.ColFromIntSlice([]int64{
		(hs0.pos.Y()*hs0.vel.X() - hs1.pos.Y()*hs1.vel.X()) - (hs0.pos.X()*hs0.vel.Y() - hs1.pos.X()*hs1.vel.Y()),
		(hs0.pos.Y()*hs0.vel.X() - hs2.pos.Y()*hs2.vel.X()) - (hs0.pos.X()*hs0.vel.Y() - hs2.pos.X()*hs2.vel.Y()),
		(hs0.pos.Z()*hs0.vel.X() - hs1.pos.Z()*hs1.vel.X()) - (hs0.pos.X()*hs0.vel.Z() - hs1.pos.X()*hs1.vel.Z()),
		(hs0.pos.Z()*hs0.vel.X() - hs2.pos.Z()*hs2.vel.X()) - (hs0.pos.X()*hs0.vel.Z() - hs2.pos.X()*hs2.vel.Z()),
		(hs0.pos.Z()*hs0.vel.Y() - hs1.pos.Z()*hs1.vel.Y()) - (hs0.pos.Y()*hs0.vel.Z() - hs1.pos.Y()*hs1.vel.Z()),
		(hs0.pos.Z()*hs0.vel.Y() - hs2.pos.Z()*hs2.vel.Y()) - (hs0.pos.Y()*hs0.vel.Z() - hs2.pos.Y()*hs2.vel.Z()),
	})
	c := a.Solve(b)
	return numbers.RatToInt[int64](numbers.RatAdd(numbers.RatAdd(c[0], c[1]), c[2]))
}

func main() {
	problem.Solve(new(Solution))
}
