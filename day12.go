package main

import (
	"aoc2024/set"
	. "aoc2024/utils"
	"bufio"

	"log"
	"os"
)

type Region struct {
	area     set.Set[Point]
	perimeta set.Set[Point]
}

func day12() {
	f, err := os.Open("day12.input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	garden := make([][]byte, 0)
	for i := 0; scanner.Scan(); i++ {
		row := make([]byte, len(scanner.Bytes()))
		copy(row, scanner.Bytes())
		garden = append(garden, row)
		// if i%50 == 0 {
		// 	for _, line := range garden {
		// 		println(string(line))
		// 	}
		// 	println()
		// }
	}

	// fmt.Printf("garden: %v\n", garden)

	regions := map[byte]map[Point]Region{}
	// var minRow, maxRow, minCol, maxCol int

	for row := range garden {
		for col, c := range garden[row] {
			_, regionInitialized := regions[c]
			if !regionInitialized {
				regions[c] = map[Point]Region{}
			}

			newRegion := Region{set.New[Point](), set.New[Point]()}
			newRegion.area.Add(Point{row, col})
			for _, diffs := range [][]int{{0, 1}, {1, 2}, {2, 1}, {1, 0}} {
				newRegion.perimeta.Add(Point{row*2 + diffs[0], col*2 + diffs[1]})
				// if row*2+diffs[0] < minRow {
				// 	minRow = row*2 + diffs[0]
				// }
				// if row*2+diffs[0] > maxRow {
				// 	maxRow = row*2 + diffs[0]
				// }
				// if col*2+diffs[1] < minCol {
				// 	minCol = col*2 + diffs[1]
				// }
				// if col*2+diffs[1] > maxCol {
				// 	maxCol = col*2 + diffs[1]
				// }

			}
			intersectedRegions := []Point{}
			for point, region := range regions[c] {
				if set.Intersection(newRegion.perimeta, region.perimeta).Size() > 0 {
					newRegion.area = set.Union(newRegion.area, region.area)
					newRegion.perimeta = set.DisjointUnion(newRegion.perimeta, region.perimeta)
					intersectedRegions = append(intersectedRegions, point)
				}
			}
			for _, point := range intersectedRegions {
				delete(regions[c], point)
			}

			regions[c][Point{row, col}] = newRegion

			// println(row, col, c)
			// fmt.Printf("fences: %v\n", fences)
			// fmt.Printf("plotAreas: %v\n", plotAreas)
		}
	}

	// rng := rand.New(rand.NewSource(0))
	// var graphicalGarden strings.Builder
	// for row := 0; row <= len(garden)*2; row++ {
	// 	for col := 0; col <= len(garden[0])*2; col++ {
	// 		// println(row, col)
	// 		if row%2 == 1 {
	// 			if col%2 == 1 {
	// 				// println("Print a character")
	// 				c := garden[row/2][col/2]
	// 				p := Point{row / 2, col / 2}
	// 				for q, region := range regions[c] {
	// 					if region.area.Has(p) {
	// 						rng.Seed(int64((q.Row + q.Col + 1) * (q.Row - q.Col + 1)))
	// 						break
	// 					}
	// 				}
	// 				colourNum := 31 + rng.Intn(7)
	// 				bgNum := 40 + rng.Intn(7)
	// 				if bgNum-10 == colourNum {
	// 					bgNum += 1
	// 				}
	// 				graphicalGarden.WriteString(fmt.Sprintf("\033[%d;%dm%s\033[0m", colourNum, bgNum, string(c)))
	// 			} else {
	// 				// println("Print a vertical fence")
	// 				p := Point{row, col}
	// 				foundFence := false
	// 				for _, labeledRegions := range regions {
	// 					for _, region := range labeledRegions {
	// 						if region.perimeta.Has(p) {
	// 							foundFence = true
	// 							break
	// 						}
	// 					}
	// 					if foundFence {
	// 						break
	// 					}
	// 				}
	// 				if foundFence {
	// 					graphicalGarden.WriteString("|")
	// 				} else {
	// 					graphicalGarden.WriteString(" ")
	// 				}
	// 			}
	// 		} else {
	// 			if col%2 == 1 {
	// 				// println("Print a horizontal fence")
	// 				p := Point{row, col}
	// 				foundFence := false
	// 				for _, labeledRegions := range regions {
	// 					for _, region := range labeledRegions {
	// 						if region.perimeta.Has(p) {
	// 							foundFence = true
	// 							break
	// 						}
	// 						if foundFence {
	// 							break
	// 						}
	// 					}
	// 				}
	// 				if foundFence {
	// 					graphicalGarden.WriteString("-")
	// 				} else {
	// 					graphicalGarden.WriteString(" ")
	// 				}
	// 			} else {
	// 				// println("Print a corner")
	// 				graphicalGarden.WriteString(" ")
	// 			}
	// 		}
	// 	}
	// 	graphicalGarden.WriteString("\n")
	// }

	// fmt.Println(graphicalGarden.String())
	// println(minRow, maxRow, minCol, maxCol)

	priceP1 := 0
	priceP2 := 0
	for _, labeledRegions := range regions {
		for _, region := range labeledRegions {
			priceP1 += region.area.Size() * region.perimeta.Size()
			priceP2 += region.area.Size() * Sides(region.perimeta)
		}
	}

	println("Part1:", priceP1)
	println("Part2:", priceP2)
}

func Sides(perimeta set.Set[Point]) int {
	sides := map[Point]set.Set[Point]{}
	for _, fence := range perimeta.Values() {
		newSide := set.New[Point]()
		newSide.Add(fence)
		allignedSides := []Point{}
		var diffs []Point
		if fence.Row%2 == 0 {
			// Horizontal fence
			diffs = []Point{{0, -2}, {0, 2}}
		} else {
			// Vertical fence
			diffs = []Point{{-2, 0}, {2, 0}}
		}

		for p, side := range sides {
			for _, diff := range diffs {
				if side.Has(Point{fence.Row + diff.Row, fence.Col + diff.Col}) && !IsCrossedCorners(perimeta, Point{fence.Row + diff.Row/2, fence.Col + diff.Col/2}) {
					allignedSides = append(allignedSides, p)
					newSide = set.Union(newSide, side)
				}
			}
		}
		for _, p := range allignedSides {
			delete(sides, p)
		}

		sides[fence] = newSide
	}

	// for p, side := range sides {
	// 	fmt.Printf("%v - side: %v\n", p, side.Values())
	// }

	return len(sides)
}

func IsCrossedCorners(perimeta set.Set[Point], point Point) bool {
	diffs := []Point{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	isCrossedCorners := true
	for _, diff := range diffs {
		isCrossedCorners = isCrossedCorners && perimeta.Has(Point{point.Row + diff.Row, point.Col + diff.Col})
	}
	return isCrossedCorners
}
