package vec

import (
	"errors"
	"fmt"
	"math"
	"math/big"

	"github.com/manisenkov/advent-of-code/pkg/numbers"
	"github.com/manisenkov/advent-of-code/pkg/rmat"
)

var (
	ErrInvalidDimensions = errors.New("invalid dimensions")
	ErrOutOfBound        = errors.New("out of bound")
	ErrInvalidType       = errors.New("invalid type")
)

// Vec represents an immutable vector of values with support of vector operations
type Vec[N numbers.Number] struct {
	data []N
}

// New return new vector from the given slice of values
func New[N numbers.Number](values []N) Vec[N] {
	data := make([]N, len(values))
	copy(data, values)
	return Vec[N]{data}
}

// Add adds two given vectors and return the result.
// It panics in case the given vectors have different size.
func (v Vec[N]) Add(w Vec[N]) Vec[N] {
	if len(v.data) != len(w.data) {
		panic(ErrInvalidDimensions)
	}
	data := make([]N, len(v.data))
	for i, x := range v.data {
		data[i] = x + w.data[i]
	}
	return Vec[N]{data}
}

// Abs returns absolute value of the given vector
func (v Vec[N]) Abs() float64 {
	s := float64(0)
	for _, n := range v.data {
		s += math.Pow(float64(n), 2)
	}
	return math.Sqrt(s)
}

// At returns i-th element of the given vector
func (v Vec[N]) At(i int) N {
	if i < 0 || i >= len(v.data) {
		panic(ErrOutOfBound)
	}
	return v.data[i]
}

// IsSimilar returns true if all dimensions of the given vectors have same signs
func (v Vec[N]) IsSimilar(w Vec[N]) bool {
	for i, item := range v.data {
		other := w.data[i]
		if (item > 0 && other < 0) ||
			(item < 0 && other > 0) ||
			(item == 0 && other != 0) ||
			(item != 0 && other == 0) {
			return false
		}
	}
	return true
}

// Dot returns the result of dot product of the two given vectors
func (v Vec[N]) Dot(w Vec[N]) N {
	if len(v.data) != len(w.data) {
		panic(ErrInvalidDimensions)
	}
	res := N(0)
	for i, item := range v.data {
		res += item * w.data[i]
	}
	return res
}

// Equal returns true if both vectors are equal
func (v Vec[N]) Equal(w Vec[N]) bool {
	if len(v.data) != len(w.data) {
		return false
	}
	for i, item := range v.data {
		if item != w.data[i] {
			return false
		}
	}
	return true
}

// Reduce returns a copy of the given vector with reduced dimensions
func (v Vec[N]) Reduce(dim int) Vec[N] {
	return Vec[N]{v.data[:dim]}
}

// Scale multiplies every element of the vector to the given factor and returns the result as a new vector
func (v Vec[N]) Scale(k N) Vec[N] {
	data := make([]N, len(v.data))
	for i, x := range v.data {
		data[i] = k * x
	}
	return Vec[N]{data}
}

// Size returns size of the given vector
func (v Vec[N]) Size() int {
	return len(v.data)
}

// Sub subtracts one given vector from another and return the result.
// It panics in case the given vectors have different size.
func (v Vec[N]) Sub(w Vec[N]) Vec[N] {
	if len(v.data) != len(w.data) {
		panic(ErrInvalidDimensions)
	}
	data := make([]N, len(v.data))
	for i, x := range v.data {
		data[i] = x - w.data[i]
	}
	return Vec[N]{data}
}

// String returns a string representation of the given vector
func (v Vec[N]) String() string {
	res := "["
	for i, item := range v.data {
		res += fmt.Sprint(item)
		if i < len(v.data)-1 {
			res += ","
		}
	}
	return res + "]"
}

// X returns the first component of a vector
func (v Vec[N]) X() N {
	return v.At(0)
}

// Y returns the second component of a vector
func (v Vec[N]) Y() N {
	return v.At(1)
}

// Z returns the third component of a vector
func (v Vec[N]) Z() N {
	return v.At(2)
}

// Cross returns the result of cross product of the two given vectors
func Cross[N numbers.Number](vs ...Vec[N]) Vec[N] {
	if len(vs) == 0 {
		return New([]N{})
	}
	sz := len(vs[0].data)
	if len(vs) != sz-1 {
		panic(ErrInvalidDimensions)
	}
	for _, v := range vs[1:] {
		if len(v.data) != sz {
			panic(ErrInvalidDimensions)
		}
	}
	res := make([]N, sz)
	for i := 0; i < sz; i++ {
		matrixData := []*big.Rat{}
		for j := 0; j < sz; j++ {
			for _, v := range vs {
				if i == j {
					continue
				}
				matrixData = append(matrixData, numbers.NumberToRat(v.data[j]))
			}
		}
		m := rmat.New(matrixData, sz-1, sz-1)
		s := m.Det()
		if i%2 == 1 {
			s = numbers.RatNeg(s)
		}
		res[i] = numbers.RatToNumber[N](s)

	}
	return Vec[N]{res}
}
