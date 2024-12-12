package main

import (
	. "aoc2024/utils"
	"bufio"
	"log"
	"os"
)

func day8() {
	f, err := os.Open("day8.input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var width, height int
	antennas := map[byte][]Point{}
	for i := 0; scanner.Scan(); i++ {
		if i == 0 {
			width = len(scanner.Bytes())
		}
		height = i + 1
		for j, val := range scanner.Bytes() {
			if val != '.' {
				antennas[val] = append(antennas[val], Point{i, j})
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
				rowDiff := p.Row - q.Row
				colDiff := p.Col - q.Col

				newANodes := []Point{
					{p.Row + rowDiff, p.Col + colDiff},
					{q.Row - rowDiff, q.Col - colDiff},
				}

				for _, aNode := range newANodes {
					if aNode.Row >= 0 && aNode.Row < height &&
						aNode.Col >= 0 && aNode.Col < width {
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
				rowDiff := p.Row - q.Row
				colDiff := p.Col - q.Col
				gcd := Gcd(rowDiff, colDiff)
				rowDiff, colDiff = rowDiff/gcd, colDiff/gcd

				for row, col := p.Row, p.Col; row >= 0 && row < height && col >= 0 && col < width; row, col = row-rowDiff, col-colDiff {
					result[Point{row, col}] = true
				}
				for row, col := p.Row+rowDiff, p.Col+colDiff; row >= 0 && row < height && col >= 0 && col < width; row, col = row+rowDiff, col+colDiff {
					result[Point{row, col}] = true
				}
			}
		}
	}
	return result
}
