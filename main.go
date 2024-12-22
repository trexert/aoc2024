package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Input day [2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22]")
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
	case "12":
		day12()
	case "14":
		day14()
	case "16":
		day16()
	case "18":
		day18()
	case "20":
		day20()
	case "22":
		day22()
	default:
		fmt.Println("No function available for day ", text)
	}
}
