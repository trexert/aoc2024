package main

import (
	"aoc2024/set"
	. "aoc2024/utils"
	"bufio"
	"log"
	"os"
)

type ReindeerState struct {
	location  Point
	direction Point
}

func day16() {
	f, err := os.Open("day16.input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	racetrack := make([]string, 0)
	var start, end Point
	for row := 0; scanner.Scan(); row++ {
		line := scanner.Text()
		racetrack = append(racetrack, line)
		for col := 0; col < len(line); col++ {
			if line[col] == 'S' {
				start = Point{row, col}
			} else if line[col] == 'E' {
				end = Point{row, col}
			}
		}
	}

	startState := ReindeerState{start, Point{0, 1}}

	paths, score := AllShortestRoutes(
		startState,
		func(state ReindeerState) bool { return state.location == end },
		func(state ReindeerState) []DijkstraNode[ReindeerState, int] { return state.Neighbours(racetrack) },
	)

	println("Part1:", score)

	// println(len(paths))

	topTierSeats := set.New[Point]()
	for _, path := range paths {
		for _, state := range path {
			topTierSeats.Add(state.location)
		}
	}

	println("Part2:", topTierSeats.Size())
}

func (this ReindeerState) Neighbours(racetrack []string) []DijkstraNode[ReindeerState, int] {
	neighbours := []DijkstraNode[ReindeerState, int]{
		{ReindeerState{this.location, Point{this.direction.Col, -this.direction.Row}}, 1000},
		{ReindeerState{this.location, Point{-this.direction.Col, this.direction.Row}}, 1000},
	}
	nextRow := this.location.Row + this.direction.Row
	nextCol := this.location.Col + this.direction.Col
	if nextRow >= 0 && nextRow < len(racetrack) && nextCol >= 0 && nextCol < len(racetrack[nextRow]) && racetrack[nextRow][nextCol] != '#' {
		neighbours = append(neighbours, DijkstraNode[ReindeerState, int]{ReindeerState{Point{nextRow, nextCol}, this.direction}, 1})
	}

	return neighbours
}
