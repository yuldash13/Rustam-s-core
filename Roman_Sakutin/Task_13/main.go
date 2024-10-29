package main

import (
	"fmt"
)

func main() {
	var (
		quantity       float64
		money          string
		menu           int
		rubblesBalance = 1000.0
		dollarsBalance = 100.0
		euroBalance    = 5.0
	)

	fmt.Println("Здравствуйте, чем могу помочь?")

mainLoop:
	for {
		fmt.Println("Меню:\n1)Баланс\n2)Внести деньги\n3)Снять деньги\n4)Перевод валюты\n5)Выход")
		fmt.Println()
		fmt.Scan(&menu)
		switch menu {
		case 1:
			fmt.Printf("У вас на счету:\n%v рублей\n%v долларов\n%v евро\n", rubblesBalance, dollarsBalance, euroBalance)
			fmt.Println()
		case 2:
			for {
				fmt.Println("Какую валюту вы хотите внести?\n1)Рубли\n2)Доллары\n3)Евро\n4)Назад")
				fmt.Println()
				fmt.Scan(&money)
				if money == "1" {
					fmt.Println("Сколько вы хотите ввести?")
					fmt.Scan(&quantity)
					if quantity < 0 {
						fmt.Println("Вы ввели неверное значение.")
						fmt.Println()
					} else {
						rubblesBalance += quantity
						fmt.Printf("На вашем счету: %v рублей.\n", rubblesBalance)
						fmt.Println()
					}
				} else if money == "2" {
					fmt.Println("Сколько вы хотите ввести?")
					fmt.Scan(&quantity)
					if quantity < 0 {
						fmt.Println("Вы ввели неверное значение.")
						fmt.Println()
					} else {
						dollarsBalance += quantity
						fmt.Printf("На вашем счету: %v долларов.\n", dollarsBalance)
						fmt.Println()
					}
				} else if money == "3" {
					fmt.Println("Сколько вы хотите ввести?")
					fmt.Scan(&quantity)
					if quantity < 0 {
						fmt.Println("Вы ввели неверное значение.")
						fmt.Println()
					} else {
						euroBalance += quantity
						fmt.Printf("На вашем счету: %v евро.\n", euroBalance)
						fmt.Println()
					}
				} else if money == "4" {
					break
				} else {
					fmt.Println("Вы ввели что-то неверно, попробуйте еще раз.")
					fmt.Println()
				}
			}
		case 3:
			for {
				fmt.Println("Какую валюту вы хотите снять?\n1)Рубли\n2)Доллары\n3)Евро\n4)Назад")
				fmt.Println()
				fmt.Scan(&money)
				if money == "1" {
					fmt.Println("Сколько вы хотите снять?")
					fmt.Scan(&quantity)
					if rubblesBalance < quantity {
						fmt.Println("Баланс не может быть отрицательным.")
						fmt.Println()
					} else {
						rubblesBalance -= quantity
						fmt.Printf("На вашем счету: %v рублей.\n", rubblesBalance)
						fmt.Println()
					}
				} else if money == "2" {
					fmt.Println("Сколько вы хотите снять?")
					fmt.Scan(&quantity)
					if dollarsBalance < quantity {
						fmt.Println("Баланс не может быть отрицательным.")
						fmt.Println()
					} else {
						dollarsBalance -= quantity
						fmt.Printf("На вашем счету: %v долларов.\n", dollarsBalance)
						fmt.Println()
					}
				} else if money == "3" {
					fmt.Println("Сколько вы хотите снять?")
					fmt.Scan(&quantity)
					if euroBalance < quantity {
						fmt.Println("Баланс не может быть отрицательным.")
						fmt.Println()
					} else {
						euroBalance -= quantity
						fmt.Printf("На вашем счету: %v евро.\n", euroBalance)
						fmt.Println()
					}
				} else if money == "4" {
					break
				} else {
					fmt.Println("Вы ввели что-то неверно, попробуйте еще раз.")
					fmt.Println()
				}
			}
		case 4:
		convert:
			for {
				fmt.Println("Вы можете совершить следующие переводы:\n1)Перевод рублей в доллары\n2)Перевод рублей в евро\n3)Перевод долларов в рубли\n4)Перевод долларов в евро\n5)Перевод евро в рубли\n6)Перевод евро в доллары\n7)Назад")
				fmt.Println()
				fmt.Scan(&menu)
				switch menu {
				case 1:
					fmt.Printf("У вас на балансе:\n%v рублей. Сколько вы хотите перевести?\n", rubblesBalance)
					fmt.Scan(&quantity)
					if rubblesBalance < quantity || quantity == 0 {
						fmt.Println("Вы делаете что-то не так.")
						fmt.Println()
					} else {
						rubblesBalance -= quantity
						dollarsBalance = quantity/97 + dollarsBalance
						fmt.Printf("У вас на счету:\n%v рублей\n%v долларов\n", rubblesBalance, dollarsBalance)
						fmt.Println()
					}
				case 2:
					fmt.Printf("У вас на балансе:\n%v рублей. Сколько вы хотите перевести?\n", rubblesBalance)
					fmt.Scan(&quantity)
					if rubblesBalance < quantity || quantity == 0 {
						fmt.Println("Вы делаете что-то не так.")
						fmt.Println()
					} else {
						rubblesBalance -= quantity
						euroBalance = quantity/105 + euroBalance
						fmt.Printf("У вас на счету:\n%v рублей\n%v евро\n", rubblesBalance, euroBalance)
						fmt.Println()
					}
				case 3:
					fmt.Printf("У вас на балансе:\n%v долларов. Сколько вы хотите перевести?\n", dollarsBalance)
					fmt.Scan(&quantity)
					if dollarsBalance < quantity || quantity == 0 {
						fmt.Println("Вы делаете что-то не так.")
						fmt.Println()
					} else {
						dollarsBalance -= quantity
						rubblesBalance = quantity*97 + rubblesBalance
						fmt.Printf("У вас на счету:\n%v рублей\n%v долларов\n", rubblesBalance, dollarsBalance)
						fmt.Println()
					}
				case 4:
					fmt.Printf("У вас на балансе:\n%v долларов. Сколько вы хотите перевести?\n", dollarsBalance)
					fmt.Scan(&quantity)
					if dollarsBalance < quantity || quantity == 0 {
						fmt.Println("Вы делаете что-то не так.")
						fmt.Println()
					} else {
						dollarsBalance -= quantity
						euroBalance = quantity*1.1 + euroBalance
						fmt.Printf("У вас на счету:\n%v долларов\n%v евро\n", dollarsBalance, euroBalance)
						fmt.Println()
					}
				case 5:
					fmt.Printf("У вас на балансе:\n%v евро. Сколько вы хотите перевести?\n", euroBalance)
					fmt.Scan(&quantity)
					if euroBalance < quantity || quantity == 0 {
						fmt.Println("Вы делаете что-то не так.")
						fmt.Println()
					} else {
						euroBalance -= quantity
						rubblesBalance = quantity*105 + rubblesBalance
						fmt.Printf("У вас на счету:\n%v рублей\n%v евро\n", rubblesBalance, euroBalance)
						fmt.Println()
					}
				case 6:
					fmt.Printf("У вас на балансе:\n%v евро. Сколько вы хотите перевести?\n", euroBalance)
					fmt.Scan(&quantity)
					if euroBalance < quantity || quantity == 0 {
						fmt.Println("Вы делаете что-то не так.")
						fmt.Println()
					} else {
						euroBalance -= quantity
						dollarsBalance = quantity/1.1 + dollarsBalance
						fmt.Printf("У вас на счету:\n%v долларов\n%v евро\n", dollarsBalance, euroBalance)
						fmt.Println()
					}
				case 7:
					break convert
				default:
					fmt.Println("Что-то пошло не так.")
					fmt.Println()
				}
			}
		case 5:
			break mainLoop
		default:
			fmt.Println("Что-то пошло не так.")
			fmt.Println()
		}
	}
}
