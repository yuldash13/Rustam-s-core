package main

import "fmt"

func main() {

	var numberOne, numberTwo int

	fmt.Println("Введите первое число:")
	fmt.Scan(&numberOne)
	fmt.Println("Введите второе число:")
	fmt.Scan(&numberTwo)

	if numberOne == numberTwo {
		fmt.Println("Оба числа равны.")
	} else if numberOne > numberTwo {
		fmt.Println("Первое число больше второго.")
	} else {
		fmt.Println("Второе число больше первого.")
	}
}
