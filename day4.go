package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Wordsearch struct {
	wordsearch [][]byte
}

func day4() {
	f, err := os.Open("day4.input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	wordsearch := make([][]byte, 140)

	for i := 0; scanner.Scan(); i++ {
		wordsearch[i] = []byte(scanner.Text())
	}

	w := Wordsearch { wordsearch }

	xmases := 0
	for row := 0; row < len(wordsearch); row++ {
		for col := 0; col < len(wordsearch[row]); col++ {
			if wordsearch[row][col] == 'X' {
				xmases += w.SearchDirections(row, col)
			}
		}
	}

	x_mases := 0

	for row := 1; row < len(wordsearch) - 1; row++ {
		for col := 1; col < len(wordsearch[row]) - 1; col++ {
			if wordsearch[row][col] == 'A' {
				down := string(append(make([]byte, 0, 2), wordsearch[row-1][col-1], wordsearch[row+1][col+1]))
				up := string(append(make([]byte, 0, 2), wordsearch[row+1][col-1], wordsearch[row-1][col+1]))
				if (down == "MS" || down == "SM") && (up == "MS" || up == "SM") {
					x_mases++
				}
				fmt.Println(up, down, x_mases)
			}
		}
	}

	fmt.Println("Part1: ", xmases)
	fmt.Println("Part2: ", x_mases)
}

func (w Wordsearch) SearchDirections(startRow int, startCol int) int {
	xmases := 0
	for rowDir := -1; rowDir <= 1; rowDir++ {
		for colDir := -1; colDir <= 1; colDir++ {
			if !(rowDir == 0 && colDir == 0) && w.CheckXmas(startRow, startCol, rowDir, colDir) {
				xmases += 1
			}
		}
	}
	return xmases
}

func (w Wordsearch) CheckXmas(startRow int, startCol int, rowDir int, colDir int) bool {
	for i, char := range []byte("XMAS") {
		row := startRow + rowDir * i
		col := startCol + colDir * i
		if row < 0 || row >= len(w.wordsearch) || col < 0 || col >= len(w.wordsearch[row]) || w.wordsearch[row][col] != char {
			return false
		}
	}
	return true
}
