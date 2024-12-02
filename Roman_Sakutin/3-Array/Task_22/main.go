package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	var (
		array  []int
		menu   string
		number int
	)

	for {
		fmt.Println("Меню:\n-array\n-add\n-sum\n-exit")
		fmt.Scan(&menu)

		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()

		switch menu {
		case "array":
			for _, n := range array {
				fmt.Print(n, " ")
			}
			fmt.Println()
		case "add":
			fmt.Println("Введите число:")
			fmt.Scan(&number)
			array = append(array, number)
		case "sum":
			var sum int
			for _, n := range array {
				sum += n
			}
			fmt.Printf("Сумма чисел массива: %d\n", sum)
		case "exit":
			return
		default:
			fmt.Println("Вы ввели что-то не то...")
		}
	}
}
