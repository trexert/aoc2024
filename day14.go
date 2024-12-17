package main

import (
	. "aoc2024/utils"
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"time"
)

type Robot struct {
	pos      Point
	velocity Point
}

const HEIGHT = 103
const WIDTH = 101

func day14() {
	f, err := os.Open("day14.input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	regex := regexp.MustCompile("p=(\\d+),(\\d+) v=(-?\\d+),(-?\\d+)")
	robots := []Robot{}
	for scanner.Scan() {
		matchingGroups := regex.FindStringSubmatch(scanner.Text())
		robot := Robot{}
		robot.pos.Col, _ = strconv.Atoi(matchingGroups[1])
		robot.pos.Row, _ = strconv.Atoi(matchingGroups[2])
		robot.velocity.Col, _ = strconv.Atoi(matchingGroups[3])
		robot.velocity.Col = (robot.velocity.Col + WIDTH) % WIDTH
		robot.velocity.Row, _ = strconv.Atoi(matchingGroups[4])
		robot.velocity.Row = (robot.velocity.Row + HEIGHT) % HEIGHT

		robots = append(robots, robot)
	}

	// Part 1
	topLeft := 0
	topRight := 0
	bottomLeft := 0
	bottomRight := 0
	for _, robot := range robots {
		finalCol := (robot.pos.Col + 100*robot.velocity.Col) % WIDTH
		finalRow := (robot.pos.Row + 100*robot.velocity.Row) % HEIGHT
		if finalCol < WIDTH/2 {
			if finalRow < HEIGHT/2 {
				topLeft++
			} else if finalRow > HEIGHT/2 {
				bottomLeft++
			}
		} else if finalCol > WIDTH/2 {
			if finalRow < HEIGHT/2 {
				topRight++
			} else if finalRow > HEIGHT/2 {
				bottomRight++
			}
		}
	}

	println("Part1:", topLeft*topRight*bottomLeft*bottomRight)

	for i := 1; true; i++ {
		stepRobots(robots)
		if mostRobotsInMiddle(robots) {
			printBathroom(robots)
			println(i)
		}
		if i%1000 == 0 {
			println(i)
		}
		time.Sleep(1 * time.Millisecond)
	}
}

func stepRobots(robots []Robot) {
	for i := range robots {
		robots[i].pos.Col = (robots[i].pos.Col + robots[i].velocity.Col) % WIDTH
		robots[i].pos.Row = (robots[i].pos.Row + robots[i].velocity.Row) % HEIGHT
	}
}

func mostRobotsInMiddle(robots []Robot) bool {
	inMiddleHalf := 0
	outsideMiddleHalf := 0
	for _, robot := range robots {
		if robot.pos.Col > WIDTH/4 && robot.pos.Col < WIDTH*3/4 {
			inMiddleHalf++
		} else {
			outsideMiddleHalf++
		}
	}
	return inMiddleHalf > outsideMiddleHalf*4
}

func printBathroom(robots []Robot) {
	bathroom := make([][]byte, HEIGHT)
	for row := 0; row < HEIGHT; row++ {
		bathroom[row] = make([]byte, WIDTH)
		for col := 0; col < WIDTH; col++ {
			bathroom[row][col] = ' '
		}
	}

	for _, robot := range robots {
		bathroom[robot.pos.Row][robot.pos.Col] = '#'
	}

	for _, line := range bathroom {
		fmt.Printf("%s\n", string(line))
	}
}
