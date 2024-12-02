package utils

func Map[T, V any](elts []T, fn func(T) V) []V {
	result := make([]V, len(elts))
	for i, elt := range elts {
		result[i] = fn(elt)
	}
	return result
}
