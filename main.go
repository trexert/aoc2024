package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Input day [2, 4, 6, 8, 10]")
	text, _ := reader.ReadString('\n')
	switch strings.TrimSpace(text) {
	case "2":
		day2()
	case "4":
		day4()
	case "6":
		day6()
	case "8":
		day8()
	case "10":
		day10()
	default:
		fmt.Println("No function available for day ", text)
	}
}
