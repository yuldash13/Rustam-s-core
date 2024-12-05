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
		finalNum, maxQuantity, curNum int
		i                             = 1
		quantity                      = 1
		num                           = array[0]
	)
	for {
		curNum = array[i]
		if num == curNum {
			quantity++
			i++
			if quantity > maxQuantity {
				finalNum = curNum
				maxQuantity = quantity
			}
		} else {
			quantity = 1
			num = array[i]
			i++
		}
		if i == len(array) {
			if maxQuantity == 0 {
				fmt.Println("В этом массиве нет повторяющихся чисел")
				break
			}
			fmt.Printf("Часто повторяющееся число: %d\n", finalNum)
			fmt.Printf("Сколько раз повторяется: %d\n", maxQuantity)
			break
		}
	}
}
