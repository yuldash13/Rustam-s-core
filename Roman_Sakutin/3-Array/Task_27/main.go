package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		bracket        string
		depth, balance int
	)

	fmt.Println("Введите скобочное выражение:")
	fmt.Scan(&bracket)

	array := strings.Split(bracket, "")

	for i := 0; i < len(array); i++ {
		if balance == -1 {
			break
		}
		if array[i] == "(" {
			balance++
			if balance > depth {
				depth = balance
			}
		} else if array[i] == ")" {
			balance--
		}
	}

	if balance == 0 {
		fmt.Printf("Ваше скобочное выражение: %v\nГлубина выражения: %d", bracket, depth)
	} else {
		fmt.Println("Вы ввели выражение неверно")
	}
}
