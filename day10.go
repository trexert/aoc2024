package main

import (
	. "aoc2024/utils"
	"bufio"

	// "fmt"
	"log"
	"os"
	"strconv"
)

func day10() {
	f, err := os.Open("day10.input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	topograph := make([][]byte, 0)
	for scanner.Scan() {
		topograph = append(topograph, scanner.Bytes())
	}

	routeScores := map[int]map[Point]int{}
	for i := 0; i < 10; i++ {
		routeScores[i] = map[Point]int{}
	}
	accessible9s := map[int]map[Point]map[Point]bool{}
	for i := 0; i < 10; i++ {
		accessible9s[i] = map[Point]map[Point]bool{}
	}
	for row := range topograph {
		for col, c := range topograph[row] {
			i, err := strconv.Atoi(string(c))
			if i == 9 {
				routeScores[i][Point{row, col}] = 1
				accessible9s[i][Point{row, col}] = map[Point]bool{{row, col}: true}
			} else if err == nil {
				routeScores[i][Point{row, col}] = 0
				accessible9s[i][Point{row, col}] = map[Point]bool{}
			}
		}
	}

	for i := 8; i >= 0; i-- {
		for point := range routeScores[i] {
			neighbours := [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
			row, col := point.Row, point.Col

			for _, diffs := range neighbours {
				neighbourRow, neighbourCol := row+diffs[0], col+diffs[1]
				routeScores[i][point] += routeScores[i+1][Point{neighbourRow, neighbourCol}]
				for accessible9 := range accessible9s[i+1][Point{neighbourRow, neighbourCol}] {
					accessible9s[i][point][accessible9] = true
				}
			}
		}
	}

	totalTrailScoreP1 := 0
	for trailhead := range accessible9s[0] {
		totalTrailScoreP1 += len(accessible9s[0][trailhead])
	}
	totalTrailScoreP2 := 0
	for trailhead := range routeScores[0] {
		totalTrailScoreP2 += routeScores[0][trailhead]
	}

	println("Part1:", totalTrailScoreP1)
	println("Part2:", totalTrailScoreP2)
}
