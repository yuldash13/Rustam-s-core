package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	randomNum := rand.Intn(101) // лучше поменяй 101 на 11, чтобы было наглядней

	var (
		result int
		array  []int
	)
	for i := 0; i <= randomNum; i++ {
		if i%3 == 0 || i%5 == 0 {
			array = append(array, i) // это если захочешь проверить на калькуляторе
			result += i
		}
	}

	fmt.Println(randomNum)
	fmt.Println(array)
	fmt.Println(result)
}
