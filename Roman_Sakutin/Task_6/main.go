package main

import "fmt"

func main() {

	var gold, quantity int

	var (
		gems     = "кристаллы"
		gemsCost = 3
	)

	fmt.Println("Я: Ох, наконец-то я дошел до магазина! Нести мешок золота было очень сложно. Сколько у меня там его?")
	fmt.Scan(&gold)
	fmt.Printf("Я: И зачем мне дали %d золота? Ах да, я должен купить %s.\n", gold, gems)
	fmt.Println("Торговец: Здравствуйте, чем могу помочь?")
	fmt.Printf("Я: Здравствуйте, %s в наличии? Сколько они стоят?\n", gems)
	fmt.Printf("Торговец: Да, есть. Один стоит %d золота.\nТорговец: Сколько вам нужно?\n", gemsCost)
	fmt.Scan(&quantity)

	finalCost := quantity * gemsCost

	for gold-finalCost < 0 {

		fmt.Println("Торговец: У вас недостаточно золота. Возьмите меньше.")
		fmt.Scan(&quantity)
		finalCost = quantity * gemsCost
	}

	gold -= finalCost

	fmt.Printf("Торговец: C вас %d золота.\n", finalCost)
	fmt.Println("Я: Вот держите.")
	fmt.Printf("(Получено %d кристаллов, осталось %d золота)", quantity, gold)
}
