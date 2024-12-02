package rmat

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func R(a, b int64) *big.Rat {
	return big.NewRat(a, b)
}

func TestNew(t *testing.T) {
	m := New([]*big.Rat{
		R(1, 2), R(3, 4), R(5, 6),
		R(-2, 1), R(-3, 2), R(0, 1),
	}, 2, 3)
	assert.Equal(t, "[[1/2,3/4,5/6],[-2,-3/2,0]]", m.String())
	assert.Equal(t, 0, R(1, 2).Cmp(m.At(0, 0)))
	assert.Equal(t, 0, R(3, 4).Cmp(m.At(0, 1)))
	assert.Equal(t, 0, R(5, 6).Cmp(m.At(0, 2)))
	assert.Equal(t, 0, R(-2, 1).Cmp(m.At(1, 0)))
	assert.Equal(t, 0, R(-3, 2).Cmp(m.At(1, 1)))
	assert.Equal(t, 0, R(-0, 1).Cmp(m.At(1, 2)))
	assert.Equal(t, 2, m.Rows())
	assert.Equal(t, 3, m.Cols())
}

func TestFromInt(t *testing.T) {
	m := FromIntSlice[int]([]int{
		1, 2,
		3, -1,
		-2, 0,
	}, 3, 2)
	assert.Equal(t, "[[1,2],[3,-1],[-2,0]]", m.String())
	assert.Equal(t, 0, R(1, 1).Cmp(m.At(0, 0)))
	assert.Equal(t, 0, R(2, 1).Cmp(m.At(0, 1)))
	assert.Equal(t, 0, R(3, 1).Cmp(m.At(1, 0)))
	assert.Equal(t, 0, R(-1, 1).Cmp(m.At(1, 1)))
	assert.Equal(t, 0, R(-2, 1).Cmp(m.At(2, 0)))
	assert.Equal(t, 0, R(-0, 1).Cmp(m.At(2, 1)))
	assert.Equal(t, 3, m.Rows())
	assert.Equal(t, 2, m.Cols())
}

func TestAt(t *testing.T) {
	m := FromIntSlice[int]([]int{
		1, 2,
		3, -1,
		-2, 0,
	}, 3, 2)
	assert.Equal(t, 0, R(-1, 1).Cmp(m.At(1, 1)))
	assert.PanicsWithError(t, "out of bound", func() { m.At(-1, 0) })
	assert.PanicsWithError(t, "out of bound", func() { m.At(0, -1) })
	assert.PanicsWithError(t, "out of bound", func() { m.At(3, 0) })
	assert.PanicsWithError(t, "out of bound", func() { m.At(0, 2) })
}

func TestAtInt64(t *testing.T) {
	m := New([]*big.Rat{
		R(1, 2), R(16, 4), R(5, 6),
		R(-2, 1), R(-3, 2), R(0, 1),
	}, 2, 3)
	assert.PanicsWithError(t, "not an integer", func() { m.AtInt64(0, 0) })
	assert.Equal(t, int64(4), m.AtInt64(0, 1))
	assert.PanicsWithError(t, "not an integer", func() { m.AtInt64(0, 2) })
	assert.Equal(t, int64(-2), m.AtInt64(1, 0))
	assert.PanicsWithError(t, "not an integer", func() { m.AtInt64(1, 1) })
	assert.Equal(t, int64(0), m.AtInt64(1, 2))
}

func TestDet(t *testing.T) {
	m0 := New([]*big.Rat{}, 0, 0)
	m1 := New([]*big.Rat{
		R(1, 2),
	}, 1, 1)
	m2 := New([]*big.Rat{
		R(1, 2), R(3, 4),
		R(5, 6), R(7, 8),
	}, 2, 2)
	m3 := New([]*big.Rat{
		R(1, 2), R(16, 4), R(5, 6),
		R(-2, 1), R(-3, 2), R(0, 1),
		R(9, 1), R(-6, 5), R(3, 8),
	}, 3, 3)
	m4 := New([]*big.Rat{
		R(1, 2), R(16, 4), R(5, 6),
		R(-2, 1), R(-3, 2), R(0, 1),
	}, 2, 3)
	assert.True(t, m0.Det().Cmp(R(0, 1)) == 0)
	assert.True(t, m1.Det().Cmp(big.NewRat(1, 2)) == 0)
	assert.True(t, m2.Det().Cmp(big.NewRat(-3, 16)) == 0)
	assert.True(t, m3.Det().Cmp(big.NewRat(511, 32)) == 0)
	assert.PanicsWithError(t, "matrix is not square", func() { m4.Det() })
}

func TestEqual(t *testing.T) {
	m1 := New([]*big.Rat{
		R(1, 2), R(16, 4), R(5, 6),
		R(-2, 1), R(-3, 2), R(0, 1),
	}, 2, 3)
	m2 := New([]*big.Rat{
		R(1, 2), R(16, 4), R(5, 6),
		R(-2, 1), R(-3, 2), R(0, 1),
	}, 2, 3)
	m3 := New([]*big.Rat{
		R(1, 2), R(16, 4), R(5, 6),
		R(-2, 1), R(-3, 2), R(0, 1),
		R(4, 5), R(-6, 8), R(1, 10),
	}, 3, 3)
	assert.True(t, m1.Equal(m2))
	assert.True(t, m2.Equal(m1))
	assert.False(t, m1.Equal(m3))
}

func TestInverse(t *testing.T) {
	m1 := New([]*big.Rat{
		R(1, 2), R(16, 4), R(5, 6),
		R(-2, 1), R(-3, 2), R(0, 1),
		R(9, 1), R(-6, 5), R(3, 8),
	}, 3, 3)
	m2 := FromIntTable([][]int{
		{1, 2, 3, 4},
		{-5, -6, -7, -8},
		{9, 10, 11, 12},
		{0, -15, -14, -13},
	})
	m3 := New([]*big.Rat{
		R(1, 2), R(16, 4), R(5, 6),
		R(-2, 1), R(-3, 2), R(0, 1),
	}, 2, 3)
	inv := m1.Inverse()
	assert.True(t, m1.Mul(inv).Equal(Ident(3)))
	assert.PanicsWithError(t, "matrix is not invertible", func() { m2.Inverse() })
	assert.PanicsWithError(t, "matrix is not square", func() { m3.Inverse() })
}

func TestLUDecompose(t *testing.T) {
	m1 := New([]*big.Rat{
		R(1, 2), R(16, 4), R(5, 6),
		R(-2, 1), R(-3, 2), R(0, 1),
		R(9, 1), R(-6, 5), R(3, 8),
	}, 3, 3)
	m2 := FromIntTable([][]int{
		{1, 2, 3, 4},
		{-5, -6, -7, -8},
		{9, 10, 11, 12},
		{0, -15, -14, -13},
	})
	l1, u1, ok1 := m1.LUDecompose()
	_, _, ok2 := m2.LUDecompose()
	assert.True(t, ok1)
	assert.True(t, l1.Mul(u1).Equal(m1))
	assert.False(t, ok2)
}

func TestLUPDecompose(t *testing.T) {
	m1 := New([]*big.Rat{
		R(1, 2), R(16, 4), R(5, 6),
		R(-2, 1), R(-3, 2), R(0, 1),
		R(9, 1), R(-6, 5), R(3, 8),
	}, 3, 3)
	m2 := FromIntTable([][]int{
		{1, 2, 3, 4},
		{-5, -6, -7, -8},
		{9, 10, 11, 12},
		{8, -15, -14, -13},
	})
	l1, u1, p1, _, ok1 := m1.LUPDecompose()
	_, _, _, _, ok2 := m2.LUPDecompose()
	assert.True(t, ok1)
	assert.False(t, ok2)
	assert.True(t, l1.Mul(u1).Equal(p1.Mul(m1)))
}

func TestMul(t *testing.T) {
	m1 := New([]*big.Rat{
		R(1, 2), R(16, 4), R(5, 6),
		R(-2, 1), R(-3, 2), R(0, 1),
	}, 2, 3)
	m2 := New([]*big.Rat{
		R(5, 3), R(-2, 5),
		R(1, 3), R(0, 1),
		R(-3, 2), R(-1, 1),
	}, 3, 2)
	m3 := New([]*big.Rat{
		R(5, 3), R(-2, 5),
		R(1, 3), R(0, 1),
		R(1, 3), R(0, 1),
		R(-3, 2), R(-1, 1),
	}, 4, 2)
	expected := New([]*big.Rat{
		R(11, 12), R(-31, 30),
		R(-23, 6), R(4, 5),
	}, 2, 2)
	assert.True(t, m1.Mul(m2).Equal(expected))
	assert.PanicsWithError(t, "invalid dimensions", func() { m1.Mul(m3) })
}

func TestScale(t *testing.T) {
	m1 := New([]*big.Rat{
		R(1, 2), R(16, 4), R(5, 6),
		R(-2, 1), R(-3, 2), R(0, 1),
	}, 2, 3)
	m2 := New([]*big.Rat{
		R(2, 10), R(32, 20), R(10, 30),
		R(-4, 5), R(-6, 10), R(0, 1),
	}, 2, 3)
	m3 := New([]*big.Rat{
		R(-3, 8), R(-48, 16), R(-15, 24),
		R(6, 4), R(9, 8), R(0, 1),
	}, 2, 3)
	m4 := FromIntTable([][]int64{
		{0, 0, 0},
		{0, 0, 0},
	})
	assert.True(t, m1.Scale(R(2, 5)).Equal(m2))
	assert.True(t, m1.Scale(R(-3, 4)).Equal(m3))
	assert.True(t, m1.Scale(R(0, 1)).Equal(m4))
}

func TestScaleInt(t *testing.T) {
	m1 := New([]*big.Rat{
		R(1, 2), R(16, 4), R(5, 6),
		R(-2, 1), R(-3, 2), R(0, 1),
	}, 2, 3)
	m2 := New([]*big.Rat{
		R(2, 2), R(32, 4), R(10, 6),
		R(-4, 1), R(-6, 2), R(0, 1),
	}, 2, 3)
	m3 := New([]*big.Rat{
		R(-3, 2), R(-48, 4), R(-15, 6),
		R(6, 1), R(9, 2), R(0, 1),
	}, 2, 3)
	m4 := FromIntTable([][]int64{
		{0, 0, 0},
		{0, 0, 0},
	})
	assert.True(t, m1.ScaleInt64(2).Equal(m2))
	assert.True(t, m1.ScaleInt64(-3).Equal(m3))
	assert.True(t, m1.ScaleInt64(0).Equal(m4))
}

func TestSolve(t *testing.T) {
	m := New([]*big.Rat{
		R(1, 2), R(16, 4), R(5, 6),
		R(-2, 1), R(-3, 2), R(0, 1),
		R(9, 1), R(-6, 5), R(3, 8),
	}, 3, 3)
	b := Col([]*big.Rat{R(2, 3), R(1, 5), R(-7, 8)})
	x := m.Solve(b)
	assert.True(t, m.Mul(Col(x)).Equal(b))
}

func TestTranspose(t *testing.T) {
	m1 := New([]*big.Rat{
		R(1, 2), R(16, 4), R(5, 6),
		R(-2, 1), R(-3, 2), R(0, 1),
	}, 2, 3)
	m2 := New([]*big.Rat{
		R(1, 2), R(-2, 1),
		R(16, 4), R(-3, 2),
		R(5, 6), R(0, 1),
	}, 3, 2)
	assert.True(t, m1.Transpose().Equal(m2))
}
