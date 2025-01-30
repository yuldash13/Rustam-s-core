package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func main() {
	var menu string
	post := []string{"Инженер", "Программист", "Мама", "Врач"}
	list := []string{"Назаров Юлдаш Русланович", "Назаров Рустам Русланович", "Атамуратова Ширин Розымбаевна", "Абдалов Артур Умарович"}
	for {
		fmt.Println("Меню:\n1)Список\n2)Добавить\n3)Удалить\n4)Поиск\n5)Выход")
		fmt.Scan(&menu)
		clearConsol()
		switch menu {
		case "Список":
			showList(list, post)
		case "Добавить":
		add:
			for {
				var lastName, name, lastLastName, newFullName, newPosition string
				showList(list, post)
				fmt.Println("Введите ФИО нового сотрудника через пробел:")
				fmt.Scan(&lastName, &name, &lastLastName)
				newFullName = lastName + " " + name + " " + lastLastName
			checking:
				for i := 0; ; i++ {
					if list[i] == newFullName {
						clearConsol()
						fmt.Println("Такой сотрудник уже существует")
						break checking
					}
					if i < len(list) {
						list = addWorker(list, newFullName)
						fmt.Println("Введите должность нового сотрудника:")
						fmt.Scan(&newPosition)
						post = addWorker(post, newPosition)
						clearConsol()
						break add
					}
				}
			}
		case "Удалить":
		delete:
			for {
				var lastName, name, lastLastName, fullName string
				showList(list, post)
				fmt.Println("Введите ФИО сотрудника которого надо удалить:")
				fmt.Scan(&lastName, &name, &lastLastName)
				clearConsol()
				fullName = lastName + " " + name + " " + lastLastName
				for i := 0; i < len(list); i++ {
					if list[i] == fullName {
						list, post = deleteWorker(list, post, fullName)
						break delete
					}
				}
				fmt.Println("Такого сотрудника не существует.")
			}
		case "Поиск":
			clearConsol()
			var lastName string
			fmt.Println("Введите фамилию сотрудника:")
			fmt.Scan(&lastName)
			newList, newPost := searchLastName(list, post, lastName)
			showList(newList, newPost)
		case "Выход":
			clearConsol()
			return
		default:
			clearConsol()
			fmt.Println("Вы ввели команду неверно.")
		}
	}
}

func searchLastName(list, post []string, lastName string) ([]string, []string) {
	var (
		newList, newPost []string
	)
	for i := 0; i < len(list); i++ {
		str := strings.Split(list[i], " ")
		if str[0] == lastName {
			newList = append(newList, list[i])
			newPost = append(newPost, post[i])
		}
	}
	return newList, newPost
}

func addWorker(array []string, str string) []string {
	var (
		newArray = make([]string, 0, len(array)+1)
	)
	slice := array[:(len(array))]
	newArray = append(newArray, slice...)
	newArray = append(newArray, str)
	return newArray
}

func deleteWorker(list, post []string, str string) ([]string, []string) {
	var (
		newList = make([]string, 0, len(list)-1)
		newPost = make([]string, 0, len(post)-1)
	)
	for i := 0; i < len(list); i++ {
		if list[i] == str {
			continue
		}
		newList = append(newList, list[i])
		newPost = append(newPost, post[i])
	}

	return newList, newPost
}

func showList(list, post []string) {
	fmt.Println("Список сотрудников:")
	for i := 0; i < len(list); i++ {
		fmt.Printf("%d)%s - %s\n", i+1, list[i], post[i])
	}
}
func clearConsol() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
