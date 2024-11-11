package main

import "fmt"

const attempt = 3

func main() {
	var password string

	for i := 0; i < attempt; i++ {
		fmt.Println("Введите пароль:\n(Решите ребус: ''Папа рок' ''соль)")
		fmt.Scan(&password)
		if password == "пароль" {
			fmt.Println("Секртеная информация: продай все акции")
			break
		}
		fmt.Println("Неверный пароль")
	}
}
