package main

import "fmt"

func main() {
	var (
		frase    string
		quantity int
	)

	fmt.Println("Введите фразу мотивирующую к действию:")
	fmt.Scan(&frase)
	fmt.Println("Введите, сколько раз эта фраза должна звучать:")
	fmt.Scan(&quantity)

	for i := 0; i < quantity; i++ {
		fmt.Printf("%s\n", frase)
	}
}
