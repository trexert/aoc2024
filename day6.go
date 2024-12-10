package main

import (
	"bufio"
	"log"
	"os"
	"slices"
)

type GuardState struct {
	row    int
	col    int
	rowDir int
	colDir int
}

type GuardPos struct {
	row int
	col int
}

func day6() {
	f, err := os.Open("day6.input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	guardMap := make([][]byte, 130)
	startRow, startCol := -1, -1
	for i := 0; scanner.Scan(); i++ {
		if len(scanner.Text()) > 0 {
			guardMap[i] = []byte(scanner.Text())
			caratPos := slices.IndexFunc(guardMap[i], func(b byte) bool { return b == '^' })
			if caratPos > 0 {
				startRow = i
				startCol = caratPos
			}
		}
	}
	if startRow < 0 || startCol < 0 {
		log.Fatal("Didn't find carat")
	}

	startingState := GuardState{startRow, startCol, -1, 0}

	println("Part1: ", FindGuardRoute(guardMap, startingState))

	locationsForLoops := 0
	for row := range guardMap {
		for col := range guardMap[row] {
			if guardMap[row][col] != '.' {
				// Skip if there's already something here
				continue
			}

			guardMap[row][col] = '#'

			if FindGuardRoute(guardMap, startingState) < 0 {
				locationsForLoops++
			}

			// Put the old space back
			guardMap[row][col] = '.'
		}
		println(row)
	}

	println("Part2: ", locationsForLoops)
}

func GuardStep(guardMap [][]byte, guardState GuardState) GuardState {
	nextRow := guardState.row + guardState.rowDir
	nextCol := guardState.col + guardState.colDir

	var result GuardState
	if nextRow < 0 || nextRow >= len(guardMap) ||
		nextCol < 0 || nextCol >= len(guardMap[nextRow]) ||
		guardMap[nextRow][nextCol] == '.' || guardMap[nextRow][nextCol] == '^' {
		result = GuardState{nextRow, nextCol, guardState.rowDir, guardState.colDir}
	} else if guardMap[nextRow][nextCol] == '#' {
		result = GuardState{guardState.row, guardState.col, guardState.colDir, -guardState.rowDir}
	} else {
		log.Fatal("Unexpected guardMap value: ", nextRow, nextCol, guardMap[nextRow][nextCol])
	}

	return result
}

func FindGuardRoute(guardMap [][]byte, startingState GuardState) int {
	guardState := startingState
	guardStates := map[GuardState]bool{}
	guardPoses := map[GuardPos]bool{}
	for guardState.row >= 0 && guardState.row < len(guardMap) &&
		guardState.col >= 0 && guardState.col < len(guardMap[guardState.row]) {
		if guardStates[guardState] {
			return -1
		}
		guardStates[guardState] = true
		guardPoses[GuardPos{guardState.row, guardState.col}] = true
		guardState = GuardStep(guardMap, guardState)
	}

	return len(guardPoses)
}
