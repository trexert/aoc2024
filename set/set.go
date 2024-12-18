package set

type Set[T comparable] struct {
	backing map[T]bool
}

func New[T comparable]() Set[T] {
	return Set[T]{map[T]bool{}}
}

func (this Set[T]) Has(value T) bool {
	return this.backing[value]
}

func (this Set[T]) Add(value T) {
	this.backing[value] = true
}

func (this Set[T]) Remove(value T) {
	delete(this.backing, value)
}

func (this Set[T]) Size() int {
	return len(this.backing)
}

func (this Set[T]) Values() []T {
	values := []T{}
	for value := range this.backing {
		values = append(values, value)
	}
	return values
}

func (this Set[T]) AddAll(ts []T) {
	for _, t := range ts {
		this.Add(t)
	}
}

func Intersection[T comparable](a, b Set[T]) Set[T] {
	result := New[T]()
	for entry := range a.backing {
		if b.Has(entry) {
			result.Add(entry)
		}
	}
	return result
}

func Union[T comparable](a, b Set[T]) Set[T] {
	result := New[T]()
	for entry := range a.backing {
		result.Add(entry)
	}
	for entry := range b.backing {
		result.Add(entry)
	}
	return result
}

func Difference[T comparable](left, right Set[T]) Set[T] {
	result := New[T]()
	for entry := range left.backing {
		if !right.Has(entry) {
			result.Add(entry)
		}
	}
	return result
}

func DisjointUnion[T comparable](a, b Set[T]) Set[T] {
	result := New[T]()
	for entry := range a.backing {
		if !b.Has(entry) {
			result.Add(entry)
		}
	}
	for entry := range b.backing {
		if !a.Has(entry) {
			result.Add(entry)
		}
	}
	return result
}
