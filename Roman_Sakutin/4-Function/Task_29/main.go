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
		symbol   string
		quantity int
	)
	fmt.Printf("Введите значок который будет показывать %s:", barName)
	fmt.Scan(&symbol)
	fmt.Printf("Введите количество %s в процентах (максимум 100):", barName)
	fmt.Scan(&quantity)

	bar := showBar(symbol, quantity)
	fmt.Printf("У вас %d %s:\n%s\n", quantity, barName, bar)
}

func showBar(symbol string, n int) string {
	array := []rune{'_', '_', '_', '_', '_', '_', '_', '_', '_', '_'}
	n /= 10
	for i := 0; i < n; i++ {
		array[i] = rune(symbol[0])
	}
	return string(array)
}
