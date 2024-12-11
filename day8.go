package main

import (
	"aoc2024/utils"
	"bufio"
	"log"
	"os"
)

type Point struct {
	row int
	col int
}

func day8() {
	f, err := os.Open("day8.input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var width, height int
	antennas := map[byte][]Point{}
	// antennaLocs := map[Point]bool{}
	for i := 0; scanner.Scan(); i++ {
		if i == 0 {
			width = len(scanner.Bytes())
		}
		height = i + 1
		for j, val := range scanner.Bytes() {
			if val != '.' {
				antennas[val] = append(antennas[val], Point{i, j})
				// antennaLocs[Point{i, j}] = true
			}
		}
	}

	println("Part1:", len(aNodesP1(antennas, height, width)))
	println("Part2:", len(aNodesP2(antennas, height, width)))
}

func aNodesP1(antennas map[byte][]Point, height int, width int) map[Point]bool {
	result := map[Point]bool{}
	for _, locations := range antennas {
		for i, p := range locations {
			for _, q := range locations[i+1:] {
				rowDiff := p.row - q.row
				colDiff := p.col - q.col

				newANodes := []Point{
					{p.row + rowDiff, p.col + colDiff},
					{q.row - rowDiff, q.col - colDiff},
				}

				for _, aNode := range newANodes {
					if aNode.row >= 0 && aNode.row < height &&
						aNode.col >= 0 && aNode.col < width {
						result[aNode] = true
					}
				}
			}
		}
	}
	return result
}

func aNodesP2(antennas map[byte][]Point, height int, width int) map[Point]bool {
	result := map[Point]bool{}
	for _, locations := range antennas {
		for i, p := range locations {
			for _, q := range locations[i+1:] {
				rowDiff := p.row - q.row
				colDiff := p.col - q.col
				gcd := utils.Gcd(rowDiff, colDiff)
				rowDiff, colDiff = rowDiff/gcd, colDiff/gcd

				for row, col := p.row, p.col; row >= 0 && row < height && col >= 0 && col < width; row, col = row-rowDiff, col-colDiff {
					result[Point{row, col}] = true
				}
				for row, col := p.row+rowDiff, p.col+colDiff; row >= 0 && row < height && col >= 0 && col < width; row, col = row+rowDiff, col+colDiff {
					result[Point{row, col}] = true
				}
			}
		}
	}
	return result
}
