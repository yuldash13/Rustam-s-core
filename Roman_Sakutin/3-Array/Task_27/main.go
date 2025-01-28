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
			depth = 0
			balance = 0
			fmt.Println("Вы ввели выражение неверно")
			return
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

	if balance > 0 {
		depth = 0
		balance = 0
		fmt.Println("Вы ввели выражение неверно")
		return
	}

	fmt.Printf("Ваше скобочное выражение: %v\nГлубина выражения: %d", bracket, depth)
}
