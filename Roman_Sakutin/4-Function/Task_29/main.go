package main

import (
	"fmt"
)

func main() {
	var (
		health = "health"
		mana   = "mana"
	)
	initBar(health)
	initBar(mana)
}

func initBar(barName string) {
	var (
		symbol   rune
		quantity int
	)
	fmt.Printf("Введите значок который будет показывать %s:\n", barName)
	fmt.Scanf("%c", &symbol)
	fmt.Scanf("%c", &symbol)
	fmt.Printf("Введите количество %s в процентах (максимум 100):\n", barName)
	fmt.Scan(&quantity)

	bar := showBar(symbol, quantity)
	fmt.Printf("У вас %d %s:\n%s\n", quantity, barName, bar)
}

func showBar(symbol rune, n int) string {
	array := []rune{'_', '_', '_', '_', '_', '_', '_', '_', '_', '_'}
	n /= 10
	for i := 0; i < n; i++ {
		array[i] = symbol
	}
	str := string(array)
	return str
}
