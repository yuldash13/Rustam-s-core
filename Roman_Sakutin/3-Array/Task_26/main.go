package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "я жестко пишу диплом"

	array := strings.Split(text, " ")

	for _, line := range array {
		fmt.Print(line)
		fmt.Println()
	}
}
