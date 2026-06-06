package main

import (
	"fmt"
	"strconv"
)

func main() {
	n := convert()
	fmt.Println(n)
}

//func convert() int {
//	var number string
//	for {
//		fmt.Println("Введите число:")
//		fmt.Scan(&number)
//		s, err := strconv.Atoi(number)
//		if err != nil {
//			fmt.Println("Вы ввели что-то неверно")
//		} else {
//			return s
//		}
//	}

func convert() int {
	var number string
	fmt.Println("Введите число:")
	fmt.Scan(&number)
	s, err := strconv.Atoi(number)
	if err != nil {
		fmt.Println("Вы ввели что-то неверно")
		return convert()
	}
	return s
}
