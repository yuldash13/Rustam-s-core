package main

import (
	"fmt"
	"strings"
)

const (
	rubUsd = 1 / 97.0
	rubEur = 1 / 105.0
	usdRub = 97.0
	usdEur = 1 / 1.1
	eurRub = 105.0
	eurUsd = 1.1
)

func main() {
	var (
		quantity, sum                float64
		answer, transfer1, transfer2 string
		menu                         int

		rubblesBalance = 1000.0
		dollarsBalance = 100.0
		euroBalance    = 5.0
	)

	bank := map[string]float64{
		"rub": rubblesBalance,
		"usd": dollarsBalance,
		"eur": euroBalance,
	}

	course := map[string]map[string]float64{
		"rub": {
			"rub": 1,
			"usd": rubUsd,
			"eur": rubEur,
		},
		"usd": {
			"rub": usdRub,
			"usd": 1,
			"eur": usdEur,
		},
		"eur": {
			"rub": eurRub,
			"usd": eurUsd,
			"eur": 1,
		},
	}

	fmt.Println("Здравствуйте, чем могу помочь?")

mainLoop:
	for {
		fmt.Println("Меню:\n1)Баланс\n2)Внести деньги\n3)Снять деньги\n4)Перевод валюты\n5)Выход")
		fmt.Scanln(&menu)
		switch menu {
		case 1:
			fmt.Printf("У вас на счету:\n%v рублей\n%v долларов\n%v евро\n", rubblesBalance, dollarsBalance, euroBalance)
		case 2:
		deposit:
			for {
				fmt.Println("Вы можете ввести следующие валюты: Rub, Usd, Eur. Какую валюту вы хотите внести?")
				fmt.Scanln(&transfer1)

				transfer1 = strings.ToLower(transfer1)

				if _, ok := bank[transfer1]; !ok {
					fmt.Println("Вы ввели данные не верно.")
					continue deposit
				}

				fmt.Println("Сколько вы хотите ввести?")
				fmt.Scanln(&quantity)

				if quantity < 0 {
					fmt.Println("Вы ввели неверное значение.")
					continue deposit
				}

				bank[transfer1] += quantity

				fmt.Printf("На вашем счету: %v %s.\n", float64(int(bank[transfer1]*100))/100.0, transfer1)

				fmt.Println("Вернуться в главное меню? (Да/Нет)")
				fmt.Scanln(&answer)

				answer = strings.ToLower(answer)

				if answer == "да" {
					break deposit
				}
			}
		case 3:
		depositOut:
			for {
				fmt.Println("Вы можете снять следующие валюты: Rub, Usd, Eur. Какую валюту вы хотите снять?")
				fmt.Scanln(&transfer1)

				transfer1 = strings.ToLower(transfer1)

				if _, ok := bank[transfer1]; !ok {
					fmt.Println("Вы ввели данные не верно.")
					continue depositOut
				}

				fmt.Println("Сколько вы хотите снять?")
				fmt.Scanln(&quantity)

				if quantity < 0 || quantity > bank[transfer1] {
					fmt.Println("Вы ввели неверное значение.")
					continue depositOut
				}

				bank[transfer1] -= quantity

				fmt.Printf("На вашем счету: %v %s.\n", float64(int(bank[transfer1]*100))/100.0, transfer1)

				fmt.Println("Вернуться в главное меню? (Да/Нет)")
				fmt.Scanln(&answer)

				answer = strings.ToLower(answer)

				if answer == "да" {
					break depositOut
				}
			}
		case 4:
		convert:
			for {
				fmt.Println("Вы можете совершить переводы между Rub, Usd, Eur. Что во что вы хотите перевести?")
				fmt.Scanln(&transfer1, &transfer2)

				transfer1 = strings.ToLower(transfer1)
				transfer2 = strings.ToLower(transfer2)

				_, ok := bank[transfer1]
				_, ok1 := bank[transfer2]
				if !ok || !ok1 {
					fmt.Println("Вы ввели данные не верно.")
					continue convert
				}

				fmt.Println("Сколько вы хотите перевести?")
				fmt.Scanln(&quantity)

				bank[transfer1] -= quantity
				if rubblesBalance < quantity || quantity == 0 {
					fmt.Println("Вы делаете что-то не так.")
					continue convert
				}

				sum = quantity * course[transfer1][transfer2]
				bank[transfer2] += sum

				fmt.Printf("У вас на счету:\n%v %s\n%v %s\n", bank[transfer1], transfer1, bank[transfer2], transfer2)
			}
		case 5:
			break mainLoop
		default:
			fmt.Println("Что-то пошло не так.")
		}
	}
}
