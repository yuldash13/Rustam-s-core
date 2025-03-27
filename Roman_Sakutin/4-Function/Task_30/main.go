package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	array := []int{1, 2, 3, 4, 5}
	fmt.Println(array)
	shuffle(array)
}

func shuffle(arr []int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(arr); i++ {
		randomNum := rand.Intn(len(arr))
		arr[i], arr[randomNum] = arr[randomNum], arr[i]
	}
	fmt.Println(arr)
}
