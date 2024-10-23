package main

import "fmt"

func main() {

	var letter rune

	fmt.Println("Введите маленькую букву:")
	fmt.Scanf("%c", &letter)

	a := 'a'
	A := 'A'

	LETTER := A + (letter - a)
	fmt.Printf("Большая буква выглядит так: %s", string(LETTER))
}
