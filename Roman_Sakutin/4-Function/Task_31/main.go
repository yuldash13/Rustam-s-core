package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

type Person struct {
	surname  string
	name     string
	lastname string
	job      string
}

func main() {
	var menu string
	person1 := Person{
		surname:  "Назаров",
		name:     "Юлдаш",
		lastname: "Русланович",
		job:      "Инженер",
	}
	person2 := Person{
		surname:  "Назаров",
		name:     "Рустам",
		lastname: "Русланович",
		job:      "Программист",
	}
	person3 := Person{
		surname:  "Атамуратова",
		name:     "Ширин",
		lastname: "Розымбаевна",
		job:      "Мама",
	}
	person4 := Person{
		surname:  "Абдалов",
		name:     "Артур",
		lastname: "Умарович",
		job:      "Врач",
	}
	people := []Person{person1, person2, person3, person4}
	for {
		fmt.Println("Меню:\n1)Список\n2)Добавить\n3)Удалить\n4)Поиск\n5)Выход")
		fmt.Scan(&menu)
		clearConsol()
		switch menu {
		case "Список":
			clearConsol()
			showList(&people)
		case "Добавить":
			clearConsol()
			addWorker(&people)
		case "Удалить":
			clearConsol()
			deleteWorker(&people)
		case "Поиск":
			clearConsol()
			searchPerson(&people)
		case "Выход":
			clearConsol()
			return
		default:
			clearConsol()
			fmt.Println("Вы ввели команду неверно.")
		}
	}
}

func searchPerson(people *[]Person) *Person {
	var exPerson Person
	fmt.Println("Введите фамилию сотрудника:")
	fmt.Scan(&exPerson.surname)
	var n int
	for i, person := range *people {
		if person.surname == exPerson.surname {
			fmt.Printf("%d)%s %s %s - %s\n", i+1, person.surname, person.name, person.lastname, person.job)
			n++
		}
	}
	if n > 0 {
		return &Person{}
	} else if n == 0 {
		fmt.Println("Такого сотрудника не существует")
	}
	return &Person{}
}

func addWorker(people *[]Person) *Person {
	var newPerson Person
	showList(people)
	fmt.Println("Введите ФИО нового сотрудника через пробел:")
	fmt.Scan(&newPerson.surname, &newPerson.name, &newPerson.lastname)
	fmt.Println("Введите должность нового сотрудника:")
	fmt.Scan(&newPerson.job)
	for _, person := range *people {
		if person == newPerson {
			fmt.Println("Такой сотрудник уже существует")
			return &Person{}
		}
	}
	*people = append(*people, newPerson)
	return &Person{}
}

func deleteWorker(people *[]Person) *Person {
	var oldPerson Person
	showList(people)
	fmt.Println("Введите ФИО сотрудника которого надо удалить:")
	fmt.Scan(&oldPerson.surname, &oldPerson.name, &oldPerson.lastname)
	fmt.Println("Введите должность сотрудника:")
	fmt.Scan(&oldPerson.job)
	var n int
	for i, person := range *people {
		if person == oldPerson {
			*people = append((*people)[:i], (*people)[i+1:]...)
			n++
		}
	}
	if n > 0 {
		return &Person{}
	} else if n == 0 {
		fmt.Println("Такого сотрудника не существует")
	}
	return &Person{}
}

func showList(people *[]Person) {
	fmt.Println("Список сотрудников:")
	for i, person := range *people {
		fmt.Printf("%d)%s %s %s - %s\n", i+1, person.surname, person.name, person.lastname, person.job)
	}
}

func clearConsol() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
