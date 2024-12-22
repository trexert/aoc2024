package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func day22() {
	f, err := os.Open("day22.input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	input := []uint{}
	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		input = append(input, uint(i))
	}
	// input := []uint{1, 2, 3, 2024}

	part1Sum := uint(0)
	diffValues := map[string]uint{}
	for _, num := range input {
		diffs := []string{}
		newDiffValues := map[string]uint{}
		for i := 0; i < 2000; i++ {
			nextNum := evolveNum(num)
			value := nextNum % 10
			diffs = append(diffs, fmt.Sprint(int(value)-(int(num)%10)))
			if len(diffs) > 4 {
				diffs = diffs[1:]
			}
			if len(diffs) == 4 {
				diffStr := strings.Join(diffs, ",")
				_, exists := newDiffValues[diffStr]
				if !exists {
					newDiffValues[diffStr] = value
				}
			}
			num = nextNum
		}
		part1Sum += num
		for diffStr, value := range newDiffValues {
			prevVal, exists := diffValues[diffStr]
			if exists {
				diffValues[diffStr] = prevVal + value
			} else {
				diffValues[diffStr] = value
			}
		}
	}

	println("Part1:", part1Sum)

	maxBananas := uint(0)
	for _, bananas := range diffValues {
		if bananas > maxBananas {
			maxBananas = bananas
		}
	}

	println("Part2:", maxBananas)
}

const MODULUS = 16777216

func evolveNum(num uint) uint {
	a := num << 6
	num = (num ^ a) % MODULUS
	a = num >> 5
	num = (num ^ a) % MODULUS
	a = num << 11
	num = (num ^ a) % MODULUS
	return num
}
