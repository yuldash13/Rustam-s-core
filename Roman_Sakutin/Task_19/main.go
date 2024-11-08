package main

import "fmt"

func main() {

	var multiOne, sumTwo int

	array := [2][4]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
	}
	for _, row := range array {
		fmt.Println(row)
	}

	multiOne = array[0][0] * array[1][0]

	for i := 0; i < 4; i++ {
		sumTwo += array[1][i]
	}

	fmt.Println(multiOne)
	fmt.Println(sumTwo)
}
