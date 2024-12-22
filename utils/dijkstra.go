package utils

import (
	"aoc2024/set"
	"container/heap"
)

type DijkstraNode[T comparable, D int | float32 | float64] struct {
	Location T
	Distance D
}

type DijkstraState[T comparable, D int | float32 | float64] struct {
	location T
	path     []T
	distance D
}

type DijkstraQueue[T comparable, D int | float32 | float64] []DijkstraState[T, D]

func (q DijkstraQueue[T, D]) Len() int           { return len(q) }
func (q DijkstraQueue[T, D]) Less(i, j int) bool { return q[i].distance < q[j].distance }
func (q DijkstraQueue[T, D]) Swap(i, j int)      { q[i], q[j] = q[j], q[i] }

func (q *DijkstraQueue[T, D]) Push(state any) {
	*q = append(*q, state.(DijkstraState[T, D]))
}

func (q *DijkstraQueue[T, D]) Pop() any {
	old := *q
	state := old[len(old)-1]
	*q = old[:len(old)-1]
	return state
}

func ShortestRoute[T comparable, D int | float32 | float64](start T, isEnd func(T) bool, neighbours func(T) []DijkstraNode[T, D]) ([]T, D) {
	paths, distance := dijkstras(start, isEnd, neighbours, false)

	return paths[0], distance
}

func AllShortestRoutes[T comparable, D int | float32 | float64](start T, isEnd func(T) bool, neighbours func(T) []DijkstraNode[T, D]) ([][]T, D) {
	return dijkstras(start, isEnd, neighbours, true)
}

func dijkstras[T comparable, D int | float32 | float64](start T, isEnd func(T) bool, neighbours func(T) []DijkstraNode[T, D], findAllRoutes bool) ([][]T, D) {
	visited := set.New[T]()
	queue := make(DijkstraQueue[T, D], 0)
	heap.Push(&queue, DijkstraState[T, D]{location: start, distance: 0})

	paths := [][]T{}
	distance := D(-1)
	foundEnd := false
	for len(queue) > 0 {
		state := heap.Pop(&queue).(DijkstraState[T, D])
		if !findAllRoutes && visited.Has(state.location) {
			continue
		}

		if foundEnd && state.distance > distance {
			break
		}

		if isEnd(state.location) {
			paths = append(paths, state.path)
			distance = state.distance
			foundEnd = true
			if !findAllRoutes {
				break
			}
		}

		visited.Add(state.location)
		for _, neighbour := range neighbours(state.location) {
			if !visited.Has(neighbour.Location) {
				newPath := make([]T, len(state.path)+1)
				copy(newPath, state.path)
				newState := DijkstraState[T, D]{
					location: neighbour.Location,
					distance: state.distance + neighbour.Distance,
					path:     append(newPath, neighbour.Location),
				}
				heap.Push(&queue, newState)
			}
		}
	}

	return paths, distance
}
