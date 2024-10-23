package main

import (
	"fmt"
	"strings"
)

func main() {

	var (
		number int
		array2 []string
	)

	fmt.Println("Введите цифру:")
	fmt.Scan(&number)

	//for i := 0; i < number; i++ {
	//	for j := 0; j <= i; j++ {
	//		fmt.Print("+")
	//	}
	//	fmt.Println()
	//}

	for i := 0; i < number; i++ {
		var array1 []rune
		for j := 0; j <= i; j++ {
			a := '+'
			array1 = append(array1, a)
		}
		array2 = append(array2, string(array1))
	}
	array3 := strings.Join(array2, "\n")
	fmt.Println(array3)
}
