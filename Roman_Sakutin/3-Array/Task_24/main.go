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

	var (
		curNum, num, i int
		finalArray     [10]int
	)

	for {
		curNum = array[i]
		num = array[i+1]
		if curNum > num {
			array[i], array[i+1] = array[i+1], array[i]
			i++
		} else {
			i++
		}
		if i == len(array)-1 {
			if array == finalArray {
				break
			}
			finalArray = array
			fmt.Println(array)
			i = 0
		}
	}
}
