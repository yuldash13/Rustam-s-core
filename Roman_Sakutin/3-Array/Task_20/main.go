package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	n = 100
)

func main() {
	var (
		array [10][10]int
		maxi  = 0
	)

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			randomNum := rand.Intn(n) + 1
			array[i][j] = randomNum
		}
	}

	fmt.Println("Массив:")

	for _, row := range array {
		for _, elem := range row {
			fmt.Print(elem, " ")
		}
		fmt.Println()
	}

	for _, row := range array {
		for _, elem := range row {
			if elem > maxi {
				maxi = elem
			}
		}
	}
	fmt.Println()
	fmt.Printf("Максимальное значение элемента массива: %d\n\n", maxi)

	fmt.Println("Массив с заменой максимального значения элемента на 0:")

	for _, row := range array {
		for _, elem := range row {
			if elem == maxi {
				elem = 0
			}
			fmt.Print(elem, " ")
		}
		fmt.Println()
	}
}
