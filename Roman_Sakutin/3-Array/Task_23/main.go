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
		quantity              = 1
		maxQuantity, finalNum int
	)
	for i := 0; i < len(array); i++ {
		if i == len(array)-1 {
			if maxQuantity == 0 {
				fmt.Println("В этом массиве нет повторяющихся чисел")
				break
			}
			fmt.Printf("Часто повторяющееся число: %d\n", finalNum)
			fmt.Printf("Сколько раз повторяется: %d\n", maxQuantity)
			break
		}
		if array[i] == array[i+1] {
			quantity++
			if quantity > maxQuantity {
				maxQuantity = quantity
				finalNum = array[i]
			}
		} else {
			quantity = 1
		}
	}
}
