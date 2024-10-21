package main

import "fmt"

func main() {

	var a, b, f int

	fmt.Println("Введите три переменные через пробел:")
	fmt.Scan(&a, &b, &f)

	fmt.Printf("Вы ввели следующие значения:\na = %d\nb = %d\nf = %d\n", a, b, f)
	fmt.Println("Решаем пример: x = (a + b - f/a) + f*a*a - (a + b)")

	x := (a + b - f/a) + f*a*a - (a + b)

	fmt.Printf("Ответ: x = %d", x)
}
