package main

import (
	"aoc2024/set"
	. "aoc2024/utils"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const MEM_WIDTH, MEM_HEIGHT = 71, 71
const BYTES_TO_DROP_P1 = 1024

func day18() {
	f, err := os.Open("day18.input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	blocks := []Point{}
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		col, _ := strconv.Atoi(line[0])
		row, _ := strconv.Atoi(line[1])
		blocks = append(blocks, Point{Row: row, Col: col})
	}

	steps := FindRoute(blocks, BYTES_TO_DROP_P1)
	println("Part1:", steps)
	droppedBytes := BinaryChop(func(bytesToDrop int) bool { return FindRoute(blocks, bytesToDrop) < 0 }, 1025, len(blocks))
	fmt.Printf("Part2: %d,%d\n", blocks[droppedBytes-1].Col, blocks[droppedBytes-1].Row)
}

func FindRoute(blocks []Point, bytesToDrop int) int {
	steps := 0
	target := Point{MEM_HEIGHT - 1, MEM_WIDTH - 1}
	unusableCells := set.New[Point]()
	unusableCells.AddAll(blocks[:bytesToDrop])
	unusableCells.Add(Point{0, 0})

	reached := []Point{{0, 0}}
	for !ArrayContains(reached, target) && len(reached) > 0 {
		nextReached := []Point{}
		for _, reachedPoint := range reached {
			newPoints := Step(reachedPoint, unusableCells)
			// fmt.Printf("newPoints: %v\n", newPoints)
			unusableCells.AddAll(newPoints)
			nextReached = append(nextReached, newPoints...)
		}
		reached = nextReached
		steps++
		// fmt.Printf("reached: %v\n", reached)
	}

	if len(reached) == 0 {
		return -1
	} else {
		return steps
	}
}

func Step(from Point, inaccessible set.Set[Point]) []Point {
	nextPoints := []Point{}
	for _, diff := range []Point{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} {
		p := Point{from.Row + diff.Row, from.Col + diff.Col}
		// println(p.Col, p.Row)
		if !inaccessible.Has(p) && p.Row >= 0 && p.Row < MEM_HEIGHT && p.Col >= 0 && p.Col < MEM_WIDTH {
			nextPoints = append(nextPoints, p)
		}
	}
	return nextPoints
}
