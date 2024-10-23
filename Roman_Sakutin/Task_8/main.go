package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		frase    string
		quantity int
	)

	fmt.Println("Введите фразу мотивирующую к действию:")
	fmt.Scan(&frase)
	fmt.Println("Введите, сколько раз эта фраза должна звучать:")
	fmt.Scan(&quantity)

	var array []string

	for i := 0; i < quantity; i++ {
		array = append(array, frase)
	}

	fmt.Println(strings.Join(array, "\n"))
}
