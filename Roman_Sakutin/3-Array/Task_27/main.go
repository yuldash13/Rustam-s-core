package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		bracket           string
		depth, balance, n int
	)

	fmt.Println("Введите скобочное выражение:")
	fmt.Scan(&bracket)

	array := strings.Split(bracket, "")

	for _, line := range array {
		if line == "(" {
			balance++
			if balance > depth {
				depth = balance
			}
		} else if line == ")" {
			balance--
		}
		if balance == -1 {
			depth = 0
			balance = 0
			fmt.Println("Вы ввели выражение неверно")
			return
		}

		if n == len(array)-1 {
			fmt.Printf("Ваше скобочное выражение: %v\nГлубина выражения: %d", bracket, depth)
			return
		}
		n++
	}

	if balance > 0 {
		depth = 0
		balance = 0
		fmt.Println("Вы ввели выражение неверно")
		return
	}
}
