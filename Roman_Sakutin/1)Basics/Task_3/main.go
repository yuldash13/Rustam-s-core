package main

import "fmt"

func main() {

	var age, n int
	var a, name, w, r, h string
	var (
		b = "год"
		c = "года"
		d = "лет"
		x = "фрукт"
		y = "фрукта"
		z = "фруктов"
		g = "его"
		j = "их"
		o = "название"
		p = "названия"
	)

	fmt.Println("Как вас зовут?")
	fmt.Scan(&name)

	fmt.Println("Сколько вам лет?")
	fmt.Scan(&age)

	if age%10 == 1 {
		a = b
	} else if age%10 > 1 && age%10 < 5 {
		a = c
	} else if age%10 >= 5 || age%10 == 0 {
		a = d
	}

	fmt.Println("Сколько у вас любимых фруктов?")
	fmt.Scan(&n)

	array := make([]string, n)

	fmt.Println("Перечислите их.")
	for i := 0; i < n; i++ {
		fmt.Scan(&array[i])
	}

	if n == 1 {
		w = x
		r = g
		h = o
	} else if n > 1 && n < 5 {
		w = y
		r = j
		h = p
	} else if n >= 5 || n == 0 {
		w = z
		r = j
		h = p
	}

	fmt.Printf("Вас зовут %s. Вам %d %s. Вы любите %d %s. Вот %s %s: %s.", name, age, a, n, w, r, h, array)

}
