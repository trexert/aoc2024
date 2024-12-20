package main

import (
	. "aoc2024/utils"
	"bufio"
	"log"
	"os"
)

type Cheat struct {
	start Point
	end   Point
}

func day20() {
	f, err := os.Open("day20.input")
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

	println("Part1:", FindCheatsBetterThan100(racetrack, start, end, 2))
	println("Part2:", FindCheatsBetterThan100(racetrack, start, end, 20))
}

func FindCheatsBetterThan100(racetrack []string, start Point, end Point, maxCheat int) int {
	p := start
	path := map[Point]int{}
	path[p] = 0
	cheatsBetterThan100 := 0
	for p != end {
		for _, diff := range []Point{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} {
			row := p.Row + diff.Row
			col := p.Col + diff.Col
			nextP := Point{row, col}
			_, haveBeenThere := path[nextP]
			if row >= 0 && row < len(racetrack) && col >= 0 && col < len(racetrack[row]) && racetrack[row][col] != '#' && !haveBeenThere {
				path[nextP] = path[p] + 1
				p = nextP
				break
			}
		}
		for cheatRowDiff := -maxCheat; cheatRowDiff <= maxCheat; cheatRowDiff++ {
			for cheatColDiff := Abs(cheatRowDiff) - maxCheat; cheatColDiff <= maxCheat-Abs(cheatRowDiff); cheatColDiff++ {
				cheatDist := Abs(cheatRowDiff) + Abs(cheatColDiff)
				cheatFrom := Point{p.Row + cheatRowDiff, p.Col + cheatColDiff}
				val, exists := path[cheatFrom]
				if exists && path[p]-val-cheatDist >= 100 {
					cheatsBetterThan100++
				}
			}
		}
	}

	return cheatsBetterThan100
}

// type RaceState struct {
// 	location Point
// 	canCheat bool
// }

// func neighbours(state RaceState, racetrack [][]byte) []DijkstraNode[RaceState, int] {
// 	neighbours := make([]DijkstraNode[RaceState, int], 0)
// 	for _, diff := range []Point{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} {
// 		row := state.location.Row + diff.Row
// 		col := state.location.Col + diff.Col
// 		if row >= 0 && row < len(racetrack) && col >= 0 && col < len(racetrack[row]) && racetrack[row][col] != '#' {
// 			neighbours = append(neighbours, DijkstraNode[RaceState, int]{
// 				Location: RaceState{location: Point{row, col}, canCheat: state.canCheat},
// 				Distance: 1
// 			})
// 		}
// 	}
// 	if state.canCheat {
// 		for _, diff := range []Point{{-2, 0}, {-1, 1}, {0, 2}, {1, 1}, {2, 0}, {1, -1}, {0, -2}, {-1, -1}} {
// 			row := state.location.Row + diff.Row
// 			col := state.location.Col + diff.Col
// 			if row >= 0 && row < len(racetrack) && col >= 0 && col < len(racetrack[row]) && racetrack[row][col] != '#' {
// 				neighbours = append(neighbours, DijkstraNode[RaceState, int]{
// 					Location: RaceState{location: Point{row, col}, canCheat: false},
// 					Distance: 2
// 				})
// 			}
// 		}
// 	}
// }
