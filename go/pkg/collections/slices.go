package collections

// Filter applies predicate to each element of the given slice and return slice with the elements
// where predicates returned true
func Filter[T any](input []T, predicate func(T) bool) []T {
	result := make([]T, 0)
	for _, x := range input {
		if predicate(x) {
			result = append(result, x)
		}
	}
	return result
}

// Map applies the given function to every item in a slice and return slice of results
func Map[T, R any](input []T, mapper func(T) R) []R {
	result := make([]R, len(input))
	for i, x := range input {
		result[i] = mapper(x)
	}
	return result
}

// All returns true if all items in a slice satisfy the given predicate
func All[T any](input []T, predicate func(T) bool) bool {
	for _, item := range input {
		if !predicate(item) {
			return false
		}
	}
	return true
}
