package vec

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewVec(t *testing.T) {
	v1 := New([]int{-1, 0, 1, 2})
	v2 := New([]int64{-1, 2, -3, 5, -8})

	assert.Equal(t, 4, v1.Size())
	assert.Equal(t, "[-1,0,1,2]", v1.String())
	assert.Equal(t, -1, v1.At(0))
	assert.Equal(t, 0, v1.At(1))
	assert.Equal(t, 1, v1.At(2))
	assert.Equal(t, 2, v1.At(3))
	assert.Equal(t, 4, v1.Size())
	assert.PanicsWithError(t, "out of bound", func() { v1.At(4) })

	assert.Equal(t, 5, v2.Size())
	assert.Equal(t, "[-1,2,-3,5,-8]", v2.String())
	assert.Equal(t, int64(-1), v2.At(0))
	assert.Equal(t, int64(2), v2.At(1))
	assert.Equal(t, int64(-3), v2.At(2))
	assert.Equal(t, int64(5), v2.At(3))
	assert.Equal(t, int64(-8), v2.At(4))
	assert.Equal(t, 5, v2.Size())
	assert.PanicsWithError(t, "out of bound", func() { v2.At(5) })
}

func TestAbs(t *testing.T) {
	v1 := New([]int{-1, -2, 3, 4})
	v2 := New([]int{0, 0})
	v3 := New([]int{1, 1, -1})
	assert.Equal(t, 5.477225575051661, v1.Abs())
	assert.Equal(t, 0.0, v2.Abs())
	assert.Equal(t, math.Sqrt(3), v3.Abs())
}

func TestAdd(t *testing.T) {
	v1 := New([]int{-1, -2, 3, 4})
	v2 := New([]int{5, 6, 7, 8})
	v3 := New([]int{1, 1, -1})
	expected := New([]int{4, 4, 10, 12})
	assert.True(t, expected.Equal(v1.Add(v2)))
	assert.PanicsWithError(t, "invalid dimensions", func() { v1.Add(v3) })
}

func TestCross(t *testing.T) {
	v1 := New([]int{-3, 0, 2})
	v2 := New([]int{5, 1, 2})
	v3 := New([]int{-6, 0, 4})
	expected1 := New([]int{-2, 16, -3})
	expected2 := New([]int{0, 0, 0})
	assert.True(t, expected1.Equal(Cross(v1, v2)))
	assert.True(t, expected2.Equal(Cross(v1, v3)))
}

func TestDot(t *testing.T) {
	v1 := New([]int{-1, -2, 3, 4})
	v2 := New([]int{5, 6, 7, 8})
	v3 := New([]int{1, 2, 3, 4, 5})
	assert.Equal(t, 36, v1.Dot(v2))
	assert.PanicsWithError(t, "invalid dimensions", func() { v1.Dot(v3) })
}

func TestIsSimilar(t *testing.T) {
	v1 := New([]int{1, 1, -1})
	v2 := New([]int{2, 3, -6})
	v3 := New([]int{2, -1, 2})
	v4 := New([]int{0, 1, 2})
	v5 := New([]int{-1, -2, -3})
	v6 := New([]int{0, 4, 6})
	assert.True(t, v1.IsSimilar(v2))
	assert.False(t, v1.IsSimilar(v3))
	assert.False(t, v1.IsSimilar(v4))
	assert.False(t, v1.IsSimilar(v5))
	assert.True(t, v4.IsSimilar(v6))
}

func TestSub(t *testing.T) {
	v1 := New([]int{-1, -2, 3, 4})
	v2 := New([]int{5, 6, 7, 8})
	v3 := New([]int{1, 1, -1})
	expected := New([]int{-6, -8, -4, -4})
	assert.True(t, expected.Equal(v1.Sub(v2)))
	assert.PanicsWithError(t, "invalid dimensions", func() { v1.Add(v3) })
}

func TestScale(t *testing.T) {
	v := New([]int{-1, -2, 3, 4})
	expected := New([]int{3, 6, -9, -12})
	assert.True(t, expected.Equal(v.Scale(-3)))
}
