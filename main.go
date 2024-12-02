package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Input day [2]")
	text, _ := reader.ReadString('\n')
	switch strings.TrimSpace(text) {
	case "2":
		day2()
	default:
		fmt.Println("No function available for day ", text)
	}
}
