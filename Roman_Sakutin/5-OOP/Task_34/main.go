package main

import (
	"fmt"
)

type character struct {
	name   string
	symbol string
	x      int
	y      int
}

func main() {
	field := [][]string{
		{"|", "-", "-", "-", "-", "-", "|"},
		{"|", " ", " ", " ", " ", " ", "|"},
		{"|", " ", " ", " ", " ", " ", "|"},
		{"|", " ", " ", " ", " ", " ", "|"},
		{"|", " ", " ", " ", " ", " ", "|"},
		{"|", " ", " ", " ", " ", " ", "|"},
		{"|", "-", "-", "-", "-", "-", "|"},
	}
	for {
		fmt.Println("Вы хотите создать нового персонажа?\n1)Да\n2)Нет")
		var answer string
		fmt.Scan(&answer)
		if answer == "Нет" || answer == "нет" {
			break
		} else if answer == "Да" || answer == "да" {
			info := saveCharacter(createCharacter(field))
			showCharacter(info, field)
		} else {
			fmt.Println("Вы ввели что-то неверно.")
		}
	}
}

func createCharacter(field [][]string) character {
	var (
		person character
	)
	fmt.Println("Создайте своего персонажа:")
	fmt.Println("1)Введите имя:\n2)Выберите обозначение персонажа (#/$/&):\n3)Введите свои координаты (x,y=1-5):")
	showField(field)
	fmt.Scan(&person.name, &person.symbol, &person.x, &person.y)
	if field[person.y][person.x] == "#" || field[person.y][person.x] == "$" || field[person.y][person.x] == "&" {
		fmt.Println("Здесь уже стоит другой игрок.")
		createCharacter(field)
	}
	if field[person.y][person.x] == "|" || field[person.y][person.x] == "-" {
		fmt.Println("Здесь нельзя встать.")
		createCharacter(field)
	}
	return person
}

func showCharacter(info map[string]character, field [][]string) {
	for {
		fmt.Println("Ниже представлены все созданные персонажи. Кого хотите посмотреть?")
		for key := range info {
			fmt.Println(key)
		}
		fmt.Println("Никого")
		var person string
		fmt.Scan(&person)
		key, exists := info[person]
		if exists {
			fmt.Printf("Имя: %v\n", key.name)
			fmt.Printf("Обозначение: %v\n", key.symbol)
			fmt.Printf("Координата x: %v\n", key.x)
			fmt.Printf("Координата y: %v\n", key.y)

			field[key.y][key.x] = key.symbol
			showField(field)
		} else if person == "Никого" {
			break
		} else {
			fmt.Println("Такого персонажа нет")
		}
	}
}

func showField(field [][]string) {
	for i := 0; i < len(field); i++ {
		fmt.Println(field[i])
	}
}

func saveCharacter(person character) map[string]character {
	info := map[string]character{
		person.name: person,
	}
	return info
}
