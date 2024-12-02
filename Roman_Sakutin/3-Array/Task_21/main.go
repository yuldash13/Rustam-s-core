package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	n = 9
)

func main() {
	var (
		array [30]int
	)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(array); i++ {
		randomNum := rand.Intn(n) + 1
		array[i] = randomNum
	}

	fmt.Println(array)

	var (
		a = 0
		b = 1
		c = 2
	)
	for {
		if a == 0 {
			if array[a] > array[b] {
				fmt.Println(array[a])
			}
		}
		if array[b] > array[a] && array[b] > array[c] {
			fmt.Println(array[b])
		}
		if c == len(array)-1 {
			if array[c] > array[b] {
				fmt.Println(array[c])
			}
			break
		}
		a++
		b++
		c++
	}
}
