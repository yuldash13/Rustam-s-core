package main

import "fmt"

func main() {
	var attemptOne, attemptTwo, attemptThree string

	fmt.Println("Введите пароль:\n(Решите ребус: ''Папа рок' ''соль)")

	fmt.Scan(&attemptOne)
	if attemptOne == "пароль" {
		fmt.Println("Секртеная информация: продай все акции")
		return
	} else {
		fmt.Println("Неверный пароль")
		fmt.Scan(&attemptTwo)
	}

	if attemptTwo == "пароль" {
		fmt.Println("Секртеная информация: продай все акции")
		return
	} else {
		fmt.Println("Неверный пароль")
		fmt.Scan(&attemptThree)
	}

	if attemptThree == "пароль" {
		fmt.Println("Секртеная информация: продай все акции")
		return
	} else {
		fmt.Println("Неверный пароль\nВы израсходавали все попытки, повторите попытку позже")
		return
	}
}
