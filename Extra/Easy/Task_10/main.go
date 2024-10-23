package main

import "fmt"

func main() {

	var numberOne, numberTwo, numberThree int

	fmt.Println("Введите первое число:")
	fmt.Scan(&numberOne)
	fmt.Println("Введите второе число:")
	fmt.Scan(&numberTwo)
	fmt.Println("Введите третье число:")
	fmt.Scan(&numberThree)

	if numberOne == numberTwo {
		numberOne += 5
		numberTwo += 5
	} else if numberTwo == numberThree {
		numberTwo += 5
		numberThree += 5
	} else if numberOne == numberThree {
		numberOne += 5
		numberThree += 5
	}
	fmt.Printf("1 число: %d\n2 число: %d\n3 число: %d\n", numberOne, numberTwo, numberThree)
}
