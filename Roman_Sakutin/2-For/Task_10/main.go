package main

import (
	"fmt"
)

func main() {
	var array []int

	for i := 5; i < 103; i += 7 {
		array = append(array, i)
	}

	fmt.Println(array)
}
