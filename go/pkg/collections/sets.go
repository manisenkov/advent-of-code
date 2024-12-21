package collections

// Set is an unordered collection of unique elements
type Set[T comparable] map[T]bool

// FirstOfSet returns first encountered element of set. It panics if set is empty
func FirstOfSet[T comparable](input Set[T]) T {
	for x := range input {
		return x
	}
	panic("empty set")
}

// SetFromSlice creates new set from the give input slice
func SetFromSlice[T comparable, S ~[]T](input S) Set[T] {
	res := make(Set[T])
	for _, el := range input {
		res[el] = true
	}
	return res
}

// SetToSlice converts the given set to a slice
func SetToSlice[T comparable, S ~[]T](input Set[T]) S {
	res := make(S, 0)
	for el := range input {
		res = append(res, el)
	}
	return res
}

// AppendToSet adds elements of the given input set to the target set
func AppendToSet[T comparable](input Set[T], target Set[T]) Set[T] {
	for el := range input {
		target[el] = true
	}
	return target
}

// AppendSliceToSet adds elements of the given input slice to the target set
func AppendSliceToSet[T comparable, S ~[]T](input S, target Set[T]) Set[T] {
	for _, el := range input {
		target[el] = true
	}
	return target
}

// MapSetTo applies the given function to every item in a set and return set of results
func MapSetTo[T, R comparable](input Set[T], mapper func(T) R) Set[R] {
	result := make(Set[R])
	for x := range input {
		result[mapper(x)] = true
	}
	return result
}

// MergeSets add values of the given source set into a target set and return target set as a result
func MergeSets[T comparable](source, target Set[T]) Set[T] {
	for el := range source {
		target[el] = true
	}
	return target
}
