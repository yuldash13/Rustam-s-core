package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	var (
		name, symbol string
	)
	fmt.Println("Введите свое имя:")
	fmt.Scan(&name)
	fmt.Println("Введите символ:")
	fmt.Scan(&symbol)

	nameStr := fmt.Sprintf("%v"+" "+name+" "+"%v", symbol, symbol)

	lenght := utf8.RuneCountInString(nameStr)

	lid := strings.Repeat(symbol, lenght)

	fmt.Printf("%s\n%s\n%s", lid, nameStr, lid)
}
