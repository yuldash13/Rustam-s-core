package main

import "fmt"

type character struct {
	name   string
	gender string
	age    int
	height int
	class  struct {
		class     string
		weapon    string
		bonusItem string
	}
}

func main() {
	info := saveCharacter(createCharacter())
	showCharacter(info)
}

func createCharacter() character {
	var person character
	fmt.Println("Создайте своего персонажа:")
	fmt.Println("1)Введите имя:\n2)Введите пол (мужской/женский):\n3)Введите возраст (18-99):\n4)Введите рост (150-250 см):")
	fmt.Println("5)Выберите класс (воин/маг/разбойник):\n6)Выберите оружие (меч/посох/кинжалы):\n7)Выберите бонусный предмет (яблоко/ключ/ничего):")
	fmt.Scan(&person.name, &person.gender, &person.age, &person.height, &person.class.class, &person.class.weapon, &person.class.bonusItem)
	return person
}

func saveCharacter(person character) map[string]character {
	info := map[string]character{
		person.name: person,
	}
	return info
}

func showCharacter(info map[string]character) {
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
			fmt.Printf("Пол: %v\n", key.gender)
			fmt.Printf("Возраст: %v\n", key.age)
			fmt.Printf("Рост: %v\n", key.height)
			fmt.Printf("Класс персонажа: %v\n", key.class.class)
			fmt.Printf("Оружие: %v\n", key.class.weapon)
			fmt.Printf("Бонусный пердмет: %v\n", key.class.bonusItem)
		} else if person == "Никого" {
			break
		} else {
			fmt.Println("Такого персонажа нет")
		}
	}
}
