package main

import "fmt"

func main() {

	var multiOne, sumTwo int

	array := [3][4]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	for _, row := range array {
		fmt.Println(row)
	}

	multiOne = 1
	for j := 0; j < len(array); j++ {
		multiOne *= array[j][0]
	}

	for i := 0; i < len(array[1]); i++ {
		sumTwo += array[1][i]
	}

	fmt.Println(multiOne)
	fmt.Println(sumTwo)
}
