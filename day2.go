package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"aoc2024/utils"
)

func day2() {
	f, err := os.Open("day2.input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	reports := make([][]int, 1000)

	for i := 0; scanner.Scan(); i++ {
		reports[i] = utils.Map(strings.Split(strings.TrimSpace(scanner.Text()), " "), func(s string) int {
			val, err := strconv.Atoi(s)
			if err != nil {
				log.Fatal(err, "\n", scanner.Text())
			}
			return val
		})
	}

	fmt.Println("part1: ", part1(reports))
	fmt.Println("part2: ", part2(reports))
}

func part1(reports [][]int) int {
	safeCount := 0
	for _, levels := range reports {
		if AreLevelsSafe(levels) {
			safeCount++
		}
	}
	return safeCount
}

func part2(reports [][]int) int {
	safeCount := 0
	for _, levels := range reports {
		for removedIndex := 0; removedIndex < len(levels); removedIndex++ {
			newLevels := make([]int, removedIndex, len(levels)-1)
			copy(newLevels, levels[:removedIndex])
			if AreLevelsSafe(append(newLevels, levels[removedIndex+1:]...)) {
				safeCount++
				break
			}
		}
	}
	return safeCount
}

func AreLevelsSafe(levels []int) bool {
	safe := true

	if len(levels) <= 1 {
		// We are safe
	} else if levels[0] < levels[1] {
		for i := 0; i < len(levels)-1; i++ {
			if levels[i+1]-levels[i] < 1 || levels[i+1]-levels[i] > 3 {
				safe = false
				break
			}
		}
	} else {
		for i := 0; i < len(levels)-1; i++ {
			if levels[i]-levels[i+1] < 1 || levels[i]-levels[i+1] > 3 {
				safe = false
				break
			}
		}
	}

	return safe
}
