package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	var array []string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Вводите цифры или числа через пробел, когда закончите напишите Стоп:")
	for scanner.Scan() {
		line := scanner.Text()
		if line == "Стоп" {
			showArray(array)
			break
		}
		array = append(array, addElements(line))
	}
}

func addElements(line string) string {
	splitLine := strings.Fields(line)
	slices.Reverse(splitLine)
	return strings.Join(splitLine, " ")
}

func showArray(arr []string) {
	slices.Reverse(arr)
	fmt.Println(strings.Join(arr, "\n"))
}
