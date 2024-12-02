package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	N := rand.Intn(16) + 10

	var array []int
	for i := 50; i <= 150; i++ {
		j := i
		for ; j > 0; j -= N {
		}
		if j == 0 {
			array = append(array, i)
		}
	}
	fmt.Println(N)
	fmt.Println(array)
}
