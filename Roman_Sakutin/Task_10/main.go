package main

import (
	"fmt"
)

func main() {
	var array []int

	for i := 5; i < 103; {
		array = append(array, i)
		i += 7
	}

	fmt.Println(array)
}