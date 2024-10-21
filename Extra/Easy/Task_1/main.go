package main

import (
	"fmt"
)

func main() {

	var (
		number int
		nums   []int
	)

	fmt.Println("Введите число:")
	fmt.Scan(&number)

	//firstOne := number / 10000
	//first := firstOne % 10
	//
	//secondTwo := number / 1000
	//second := secondTwo % 10
	//
	//thirdThree := number / 100
	//third := thirdThree % 10
	//
	//fourthFour := number / 10
	//fourth := fourthFour % 10
	//
	//fifth := number % 10
	//
	//fmt.Printf("Первая цифра: %d \nВторая цифра: %d \nТретья цифра %d \nЧетвертая цифра %d \nПятая цифра %d \n", first, second, third, fourth, fifth)

	//numberString := strconv.Itoa(number)
	//
	//for i := 0; i < len(numberString); i++ {
	//	fmt.Println(string(numberString[i]))
	//}

	//numberString := strconv.Itoa(number)
	//
	//for i := len(numberString) - 1; i >= 0; i-- {
	//	fmt.Println(string(numberString[i]))
	//}

	for number > 0 {
		nums = append(nums, number%10)
		number /= 10
	}

	for i := len(nums) - 1; i >= 0; i-- {
		fmt.Printf("%d цифра равна %d\n", (i-len(nums))*-1, nums[i])
	}
}
