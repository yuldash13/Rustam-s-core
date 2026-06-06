package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 4, 5}

	var n int

	fmt.Println(array)
	fmt.Println("На сколько позиций влево вы хотите сдвинуть массив?")
	fmt.Scan(&n)

	if n > 5 {
		fmt.Println("Максимальное значение на которое можно сдвинуть массив - 5.")
		fmt.Println("На сколько позиций влево вы хотите сдвинуть массив?")
		fmt.Scan(&n)
	}

	slice1 := array[n:5]
	slice2 := array[0:n]
	array = append(slice1, slice2...)

	fmt.Println(array)
}
