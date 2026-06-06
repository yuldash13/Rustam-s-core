package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	n = 9
	m = 30
)

func main() {
	var (
		array = make([]int, 0, m)
	)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < m; i++ {
		randomNum := rand.Intn(n) + 1
		array = append(array, randomNum)
	}

	fmt.Println(array)

	for i := 0; i < len(array); i++ {
		if (i == len(array)-1 || array[i] > array[i+1]) && (i == 0 || array[i] > array[i-1]) {
			fmt.Println(array[i])
		}
	}
}
