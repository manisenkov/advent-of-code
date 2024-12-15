package collections

import "errors"

var ErrEmptySlice = errors.New("slice is empty")

// Filter applies predicate to each element of the given slice and return slice with the elements
// where predicates returned true
func Filter[T any, S ~[]T](input S, predicate func(T) bool) []T {
	result := make([]T, 0)
	for _, x := range input {
		if predicate(x) {
			result = append(result, x)
		}
	}
	return result
}

// MapTo applies the given function to every item in a slice and return slice of results
func MapTo[T, R any, S ~[]T](input S, mapper func(T) R) []R {
	result := make([]R, len(input))
	for i, x := range input {
		result[i] = mapper(x)
	}
	return result
}

// All returns true if all items in a slice satisfy the given predicate
func All[T any, S ~[]T](input S, predicate func(T) bool) bool {
	for _, item := range input {
		if !predicate(item) {
			return false
		}
	}
	return true
}

// Any returns true if any of the given items satisfies the given predicate
func Any[T any, S ~[]T](input S, predicate func(T) bool) bool {
	for _, item := range input {
		if predicate(item) {
			return true
		}
	}
	return false
}

// Reduce takes a slice of values of type T and reduces it to a single value by applying
// the given reducer function to each value in sequence to combine it with the current result.
// It starts from the first element of the slice. It panics if the given slice is empty.
func Reduce[T any, S ~[]T](input S, reducer func(T, T) T) T {
	if len(input) == 0 {
		panic(ErrEmptySlice)
	}
	return ReduceWithInit(input[1:], input[0], reducer)
}

// ReduceWithInit takes a slice of values of type T and reduces it to a single value by applying
// the given reducer function to each value in sequence to combine it with the current result.
// It starts from the given init value.
func ReduceWithInit[T any, S ~[]T](input S, init T, reducer func(T, T) T) T {
	res := init
	for _, x := range input {
		res = reducer(res, x)
	}
	return res
}
