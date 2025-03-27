package main

import (
	"fmt"
	"strings"
)

const (
	z = 5
)

type Character struct {
	name   string
	symbol string
	x      int
	y      int
}

type Drawer struct {
	field [][]string
}

var info map[string]Character

func main() {
	info = make(map[string]Character, 100)
	drawer := NewDrawer()
	drawer.initField()
	for {
		fmt.Println("Вы хотите создать нового персонажа?\n1)Да\n2)Нет")
		var answer string
		fmt.Scan(&answer)
		if strings.ToLower(answer) == "нет" {
			break
		} else if strings.ToLower(answer) == "да" {
			character := createCharacter(*drawer)
			character.saveCharacter()
			fmt.Println("Вы хотите посмотреть созданных персонажей?\n1)Да\n2)Нет")
			fmt.Scan(&answer)
			if strings.ToLower(answer) == "нет" {
				break
			} else if strings.ToLower(answer) == "да" {
				character.showCharacter(info, drawer)
			}
		} else {
			fmt.Println("Вы ввели что-то неверно.")
		}
	}
}

func NewCharacter(name string, symbol string, x, y int) *Character {
	return &Character{
		name:   name,
		symbol: symbol,
		x:      x,
		y:      y,
	}
}

func NewDrawer() *Drawer {
	return &Drawer{}
}

func createCharacter(drawer Drawer) *Character {
	fmt.Println("Создайте своего персонажа:")
	fmt.Println("1)Введите имя:\n2)Выберите обозначение персонажа (#/$/&):\n3)Введите свои координаты (x,y=1-5):")
	drawer.showField()
	var (
		name, symbol string
		x, y         int
	)
	fmt.Scan(&name, &symbol, &x, &y)
	if x > 5 || y > 5 || x < 1 || y < 1 {
		fmt.Println("Значение координаты должно быть от 1 до 5.")
		createCharacter(drawer)
	}
	return NewCharacter(name, symbol, x, y)
}

func (c *Character) saveCharacter() {
	info[c.name] = *c
}

func (c *Character) showCharacter(info map[string]Character, drawer *Drawer) {
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

			drawer.saveSym(key.symbol, key.x, key.y)
			drawer.showField()
		} else if person == "Никого" {
			break
		} else {
			fmt.Println("Такого персонажа нет")
		}
	}
}

func (d *Drawer) showField() {
	for i := 0; i < len(d.field); i++ {
		fmt.Println(d.field[i])
	}
}

func (d *Drawer) initField() {
	d.field = [][]string{
		{"|", "-", "-", "-", "-", "-", "|"},
		{"|", " ", " ", " ", " ", " ", "|"},
		{"|", " ", " ", " ", " ", " ", "|"},
		{"|", " ", " ", " ", " ", " ", "|"},
		{"|", " ", " ", " ", " ", " ", "|"},
		{"|", " ", " ", " ", " ", " ", "|"},
		{"|", "-", "-", "-", "-", "-", "|"},
	}
}

func (d *Drawer) saveSym(symbol string, x, y int) {
	d.field[y][x] = symbol
}
