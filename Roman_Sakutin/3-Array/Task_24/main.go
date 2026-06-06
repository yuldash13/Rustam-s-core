package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	n = 10
)

func main() {
	var array [10]int

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(array); i++ {
		randomNum := rand.Intn(n) + 1
		array[i] = randomNum
	}
	fmt.Println(array)
	var m int
	for i := 0; ; i++ {
		if m == 10 {
			fmt.Println(array)
			break
		}
		if i == len(array)-1 {
			m++
			i = 0
		}
		if array[i] > array[i+1] {
			array[i], array[i+1] = array[i+1], array[i]
		}
	}
}
