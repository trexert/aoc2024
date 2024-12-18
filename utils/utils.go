package utils

func Map[T, V any](elts []T, fn func(T) V) []V {
	result := make([]V, len(elts))
	for i, elt := range elts {
		result[i] = fn(elt)
	}
	return result
}

func Gcd(a int, b int) int {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	if a == 0 || b == 0 {
		return 1
	}

	if b > a {
		a, b = b, a
	}

	for b > 0 {
		a, b = b, a%b
	}

	return a
}

type Point struct {
	Row int
	Col int
}

func ArrayContains[T comparable](as []T, x T) bool {
	for _, a := range as {
		if a == x {
			return true
		}
	}
	return false
}
