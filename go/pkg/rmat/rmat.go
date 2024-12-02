package rmat

import (
	"errors"
	"math/big"

	"github.com/manisenkov/advent-of-code/pkg/collections"
	"github.com/manisenkov/advent-of-code/pkg/numbers"
)

var (
	ErrInvalidDimensions = errors.New("invalid dimensions")
	ErrNotEnoughValues   = errors.New("not enough values to populate the matrix")
	ErrNotInvertible     = errors.New("matrix is not invertible")
	ErrNotSquareMatrix   = errors.New("matrix is not square")
	ErrOutOfBound        = errors.New("out of bound")
)

// Immutable matrix of rational numbers (`*big.Rat`)
type Matrix struct {
	data []*big.Rat
	rows int
	cols int
}

// FromIntSlice takes a 2-dimensional table of any integer values and create a new rational matrix out of them.
// It panics if dimensions are not regular
func FromIntTable[N numbers.AnyInt](values [][]N) Matrix {
	rows := len(values)
	if rows == 0 {
		panic(ErrInvalidDimensions)
	}
	cols := len(values[0])
	data := make([]*big.Rat, rows*cols)
	for r, row := range values {
		if len(row) != cols {
			panic(ErrInvalidDimensions)
		}
		for c, elem := range row {
			data[r*cols+c] = new(big.Rat).SetInt64(int64(elem))
		}
	}
	return Matrix{data, rows, cols}
}

// FromIntSlice takes a slice of any integer values and create a new rational matrix out of them
func FromIntSlice[N numbers.AnyInt](input []N, rows, cols int) Matrix {
	data := collections.Map(input, numbers.RatFromInt[N])
	return Matrix{data, rows, cols}
}

// Ident return a square identity matrix of the given size
func Ident(n int) Matrix {
	res := Null(n, n)
	for i := 0; i < n; i++ {
		res.set(i, i, numbers.RatOne())
	}
	return res
}

// New takes a slice of rational numbers and dimensions to create a new rational matrix.
// It panics if there are invalid dimensions or if the given slice of values does not have as many values
// as specified by the dimension
func New(input []*big.Rat, rows, cols int) Matrix {
	if rows < 0 || cols < 0 {
		panic(ErrInvalidDimensions)
	}
	if rows == 0 || cols == 0 {
		rows = 0
		cols = 0
	}
	size := rows * cols
	if len(input) != size {
		panic(ErrNotEnoughValues)
	}
	data := make([]*big.Rat, size)
	for i := 0; i < size; i++ {
		data[i] = new(big.Rat).Set(input[i])
	}
	return Matrix{
		data,
		rows,
		cols,
	}
}

// Null returns a zero matrix with the given dimensions
func Null(rows, cols int) Matrix {
	if rows < 0 || cols < 0 {
		panic(ErrInvalidDimensions)
	}
	data := make([]*big.Rat, rows*cols)
	return Matrix{data, rows, cols}
}

// Col create a column vector from the given slice of values
func Col(data []*big.Rat) Matrix {
	return New(data, len(data), 1)
}

// ColFromIntSlice create a column vector from the given slice of integers
func ColFromIntSlice[N numbers.AnyInt](input []N) Matrix {
	return FromIntSlice(input, len(input), 1)
}

// Add adds the two given matrices and returns the result.
// It panics if the given matrices have different dimensions
func (a Matrix) Add(b Matrix) Matrix {
	if a.rows != b.rows || a.cols != b.cols {
		panic(ErrInvalidDimensions)
	}
	res := Null(a.rows, a.cols)
	for r := 0; r < a.rows; r++ {
		for c := 0; c < a.rows; c++ {
			res.set(r, c, numbers.RatAdd(a.at(r, c), b.at(r, c)))
		}
	}
	return res
}

// At returns a value from the matrix for the given row and column
func (a Matrix) At(r, c int) *big.Rat {
	if r < 0 || r >= a.rows || c < 0 || c >= a.cols {
		panic(ErrOutOfBound)
	}
	return a.at(r, c)
}

// AtInt64 returns an integer value from the matrix for the given index.
// The function panics if the returned element is not an integer
func (a Matrix) AtInt64(row, col int) int64 {
	return numbers.RatToInt[int64](a.At(row, col))
}

// Det returns the determinant of the given matrix
func (a Matrix) Det() *big.Rat {
	if a.rows == 0 {
		return numbers.RatZero()
	}
	_, dU, _, numSwaps, ok := a.LUPDecompose()
	if !ok {
		// Matrix is not invertible so the determinant is zero
		return numbers.RatZero()
	}
	res := numbers.RatOne()
	for i := 0; i < a.rows; i++ {
		res = numbers.RatMul(res, dU.at(i, i))
	}
	if numSwaps%2 == 0 {
		return res
	} else {
		return numbers.RatNeg(res)
	}
}

// Inverse returns the inverse matrix for the given matrix, so that `m.Mul(m.Inv()).Equal(Ident(m.Rows()))` is true.
// It panics if no inversion matrix exists for the given one
//
// https://en.wikipedia.org/wiki/LU_decomposition#Inverting_a_matrix
func (a Matrix) Inverse() Matrix {
	if a.rows == 0 {
		return a
	}
	dL, dU, ok := a.LUDecompose()
	if !ok {
		panic(ErrNotInvertible)
	}
	dA := dL.Sub(Ident(a.rows)).Add(dU)
	res := Ident(a.rows)
	for c := 0; c < a.cols; c++ {
		for r := 0; r < a.rows; r++ {
			for k := 0; k < r; k++ {
				p := numbers.RatSub(res.at(r, c), numbers.RatMul(dA.at(r, k), res.at(k, c)))
				res.set(r, c, p)
			}
		}

		for r := a.rows - 1; r >= 0; r-- {
			for k := r + 1; k < a.rows; k++ {
				p := numbers.RatSub(res.at(r, c), numbers.RatMul(dA.at(r, k), res.at(k, c)))
				res.set(r, c, p)
			}
			res.set(r, c, numbers.RatQuo(res.at(r, c), dA.at(r, r)))
		}
	}

	return res
}

// LUDecompose calculate the LU-decomposition of the given matrix, so that `L.Mul(U).Equal(m)` is true.
// If the given matrix is not invertible, LU-decomposition can not be found
//
// https://en.wikipedia.org/wiki/LU_decomposition
func (a Matrix) LUDecompose() (Matrix, Matrix, bool) {
	if a.rows != a.cols {
		panic(ErrNotSquareMatrix)
	}
	dL := Ident(a.rows)
	dU := Null(a.rows, a.cols)
	for r := 0; r < a.rows; r++ {
		for c := 0; c < a.cols; c++ {
			if r <= c {
				s := numbers.RatZero()
				for k := 0; k < r; k++ {
					s = numbers.RatAdd(s, numbers.RatMul(dL.at(r, k), dU.at(k, c)))
				}
				dU.set(r, c, numbers.RatSub(a.at(r, c), s))
			} else {
				if numbers.RatIsZero(dU.at(c, c)) {
					return Matrix{}, Matrix{}, false
				}
				s := numbers.RatZero()
				for k := 0; k < c; k++ {
					s = numbers.RatAdd(s, numbers.RatMul(dL.at(r, k), dU.at(k, c)))
				}
				dL.set(r, c, numbers.RatQuo(numbers.RatSub(a.at(r, c), s), dU.at(c, c)))
			}
		}
	}
	return dL, dU, true
}

// LUPDecompose calculate the LUP-decomposition of the given matrix, so that `L.Mul(U).Equal(P.Mul(m))` is true.
// The function return L, U, P matrices and number of swaps as well as the result of decomposition (if it's false
// the decomposition can't be found).
//
// https://en.wikipedia.org/wiki/LU_decomposition
func (a Matrix) LUPDecompose() (Matrix, Matrix, Matrix, int, bool) {
	if a.rows != a.cols {
		panic(ErrNotSquareMatrix)
	}
	sz := a.rows
	dP := Ident(a.rows)
	dLU := Null(a.rows, a.cols).Add(a)
	numSwaps := 0
	for c := 0; c < sz; c++ {
		maxValue := numbers.RatZero()
		maxRow := c

		for r := c; r < sz; r++ {
			elAbs := numbers.RatAbs(dLU.at(r, c))
			if elAbs.Cmp(maxValue) > 0 {
				maxValue = elAbs
				maxRow = r
			}
		}

		if maxValue.Cmp(numbers.RatZero()) == 0 {
			return Matrix{}, Matrix{}, Matrix{}, 0, false
		}

		if maxRow != c {
			row := make([]*big.Rat, sz)
			copy(row, dP.data[maxRow*sz:maxRow*sz+sz])
			copy(dP.data[maxRow*sz:], dP.data[c*sz:c*sz+sz])
			copy(dP.data[c*sz:], row)

			row = make([]*big.Rat, sz)
			copy(row, dLU.data[maxRow*sz:maxRow*sz+sz])
			copy(dLU.data[maxRow*sz:], dLU.data[c*sz:c*sz+sz])
			copy(dLU.data[c*sz:], row)

			numSwaps += 1
		}

		for r := c + 1; r < sz; r++ {
			dLU.set(r, c, numbers.RatQuo(dLU.at(r, c), dLU.at(c, c)))
			for k := c + 1; k < sz; k++ {
				dLU.set(r, k, numbers.RatSub(dLU.at(r, k), numbers.RatMul(dLU.at(r, c), dLU.at(c, k))))
			}
		}
	}
	dL := Ident(dLU.rows)
	dU := Null(dLU.rows, dLU.cols)
	for r := 0; r < sz; r++ {
		for c := 0; c < sz; c++ {
			if r > c {
				dL.set(r, c, dLU.at(r, c))
			} else if r <= c {
				dU.set(r, c, dLU.at(r, c))
			}
		}
	}
	return dL, dU, dP, numSwaps, true
}

// Rows returns the number of rows of the given matrix
func (a Matrix) Rows() int {
	return a.rows
}

// Cols returns the number of columns of the given matrix
func (a Matrix) Cols() int {
	return a.cols
}

// Mul multiplies the given matrices and returns the result
func (a Matrix) Mul(b Matrix) Matrix {
	if a.cols != b.rows {
		panic(ErrInvalidDimensions)
	}
	res := Null(a.rows, b.cols)
	for r := 0; r < a.rows; r++ {
		for c := 0; c < b.cols; c++ {
			s := numbers.RatZero()
			for i := 0; i < a.cols; i++ {
				s = numbers.RatAdd(s, numbers.RatMul(a.at(r, i), b.at(i, c)))
			}
			res.set(r, c, s)
		}
	}
	return res
}

// Equal compare the given matrix to another one. It returns false if matrices are not equal
func (a Matrix) Equal(b Matrix) bool {
	if a.rows != b.rows || a.cols != b.cols {
		return false
	}
	for r := 0; r < a.rows; r++ {
		for c := 0; c < a.cols; c++ {
			if !numbers.RatEqual(a.at(r, c), b.at(r, c)) {
				return false
			}
		}
	}
	return true
}

// Scale multiplies every element of the matrix to the given rational factor and returns the result as a new matrix
func (a Matrix) Scale(p *big.Rat) Matrix {
	res := Null(a.rows, a.cols)
	for r := 0; r < a.rows; r++ {
		for c := 0; c < a.cols; c++ {
			res.set(r, c, numbers.RatMul(p, a.at(r, c)))
		}
	}
	return res
}

// ScaleInt64 multiplies every element of the matrix to the given integer factor and returns the result as a new matrix
func (a Matrix) ScaleInt64(p int64) Matrix {
	return a.Scale(numbers.RatFromInt(p))
}

// Solve returns the solution for the system of linear equations `A * x = b` for `x`.
// It panics if system is not solvable
func (a Matrix) Solve(b Matrix) []*big.Rat {
	if b.rows != a.rows && b.cols != 1 {
		panic(ErrInvalidDimensions)
	}
	return a.Inverse().Mul(b).data
}

// Sub subtracts a matrix passed as an argument from the given one and returns the result.
// It panics if matrices have different dimensions
func (a Matrix) Sub(b Matrix) Matrix {
	if a.rows != b.rows || a.cols != b.cols {
		panic(ErrInvalidDimensions)
	}
	res := Null(a.rows, a.cols)
	for r := 0; r < a.rows; r++ {
		for c := 0; c < a.rows; c++ {
			res.set(r, c, numbers.RatSub(a.at(r, c), b.at(r, c)))
		}
	}
	return res
}

// String returns the stringified representation of the given matrix
func (a Matrix) String() string {
	res := "["
	for r := 0; r < a.rows; r++ {
		res += "["
		for c := 0; c < a.cols; c++ {
			res += a.at(r, c).RatString()
			if c < a.cols-1 {
				res += ","
			}
		}
		res += "]"
		if r < a.rows-1 {
			res += ","
		}
	}
	return res + "]"
}

// Transpose transposes the given matrix and returns result as a new matrix
func (a Matrix) Transpose() Matrix {
	res := Null(a.cols, a.rows)
	for r := 0; r < a.rows; r++ {
		for c := 0; c < a.cols; c++ {
			res.set(c, r, numbers.RatCopy(a.at(r, c)))
		}
	}
	return res
}

func (a *Matrix) idx(row, col int) int {
	return row*a.cols + col
}

func (a *Matrix) at(row, col int) *big.Rat {
	res := a.data[a.idx(row, col)]
	if res == nil {
		return new(big.Rat)
	}
	return res
}

func (a *Matrix) set(row, col int, value *big.Rat) {
	a.data[a.idx(row, col)] = numbers.RatCopy(value)
}
