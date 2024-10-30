package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var (
		name, symbol string
		array        []rune
	)
	fmt.Println("Введите свое имя:")
	fmt.Scan(&name)
	fmt.Println("Введите символ:")
	fmt.Scan(&symbol)

	nameStr := symbol + " " + name + " " + symbol

	lenght := utf8.RuneCountInString(nameStr)

	symbolRune := rune(symbol[0])

	for i := 0; i < lenght; i++ {
		array = append(array, symbolRune)
	}

	lid := string(array)

	fmt.Printf("%s\n%s\n%s", lid, nameStr, lid)
}
