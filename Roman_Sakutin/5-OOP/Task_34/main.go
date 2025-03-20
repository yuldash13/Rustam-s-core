package main

import (
	"fmt"
)

type Character struct {
	name string
	Drawer
}

type Drawer struct {
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
			fmt.Println("Создайте своего персонажа:")
			fmt.Println("1)Введите имя:\n2)Выберите обозначение персонажа (#/$/&):\n3)Введите свои координаты (x,y=1-5):")
			showField(field)
			var (
				name, symbol string
				x, y         int
			)
			fmt.Scan(&name, &symbol, &x, &y)
			newDrawer := NewDrawer(symbol, x, y)
			info := saveCharacter(NewCharacter(name,*newDrawer))
			showCharacter(info, field)
		} else {
			fmt.Println("Вы ввели что-то неверно.")
		}
	}
}

func NewCharacter(name string, drawer Drawer) *Character {
	return &Character{name: name, Drawer: drawer}
}

func NewDrawer(symbol string, x, y int) *Drawer{
	return &Drawer{symbol: symbol, x: x, y: y}
}

func saveCharacter(c *Character) map[string]Character {
	info := map[string]Character{
		c.name: *c,
	}
	return info
}

func showCharacter(info map[string]Character, field [][]string) {
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
			fmt.Printf("Обозначение: %v\n", key.Drawer.symbol)
			fmt.Printf("Координата x: %v\n", key.Drawer.x)
			fmt.Printf("Координата y: %v\n", key.Drawer.y)

			field[key.Drawer.y][key.Drawer.x] = key.Drawer.symbol
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

