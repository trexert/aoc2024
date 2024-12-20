package utils

import (
	"aoc2024/set"
	"container/heap"
)

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

type DijkstraNode[T comparable, D int | float32 | float64] struct {
	Location T
	Distance D
}

type DijkstraState[T comparable, D int | float32 | float64] struct {
	location T
	path []T
	distance D
}

type DijkstraQueue[T comparable, D int | float32 | float64] []DijkstraState[T, D]

func (q DijkstraQueue[T, D]) Len() int { return len(q) }
func (q DijkstraQueue[T, D]) Less(i, j int) bool { return q[i].distance < q[j].distance }
func (q DijkstraQueue[T, D]) Swap(i, j int) { q[i], q[j] = q[j], q[i] }

func (q *DijkstraQueue[T, D]) Push(state any) {
	*q = append(*q, state.(DijkstraState[T, D]))
}

func (q *DijkstraQueue[T, D]) Pop() any {
	old := *q
	state := old[len(old) - 1]
	*q = old[:len(old) - 1]
	return state
}

func Dijkstras[T comparable, D int | float32 | float64](start T, isEnd func(T) bool, neighbours func(T) []DijkstraNode[T, D]) ([]T, D) {
	visited := set.New[T]()
	queue := make(DijkstraQueue[T, D], 0)
	heap.Push(&queue, DijkstraState[T, D]{ location: start, distance: 0 })

	path := []T{}
	distance := D(-1)
	for len(queue) > 0 {
		state := heap.Pop(&queue).(DijkstraState[T, D])

		if isEnd(state.location) {
			path = state.path
			distance = state.distance
			break
		}

		visited.Add(state.location)
		for _, neighbour := range neighbours(state.location) {
			if !visited.Has(neighbour.location) {
				newPath := make([]T, len(state.path) + 1)
				copy(newPath, state.path)
				newState := DijkstraState[T, D] {
					location: neighbour.location,
					distance: state.distance + neighbour.distance,
					path: append(newPath, neighbour.location),
				}
				heap.Push(&queue, newState)
			}
		}
	}

	return path, distance
}
