package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 4, 5}

	var (
		i, position, n, a int
	)

	fmt.Println(array)
	fmt.Println("На сколько позиций влево вы хотите сдвинуть цифры?")
	fmt.Scan(&position)

	for {
		array[i], array[i+1] = array[i+1], array[i]
		i++
		a++
		if a == 4 {
			fmt.Println(array)
			a = 0
			n++
		}
		if i == len(array)-1 {
			i = 0
		}
		if n == position {
			break
		}
	}
}
