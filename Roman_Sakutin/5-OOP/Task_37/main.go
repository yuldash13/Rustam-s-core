package main

import (
	"fmt"
)

func main() {
	storage := NewStorage()
	storage.InitStorage()

	var menu string
	for {
		fmt.Println("Меню:\n1)Покажи\n2)Добавь\n3)Удали\n4)Поиск\n5)Выход")
		fmt.Println()
		fmt.Scan(&menu)
		switch menu {
		case "Покажи":
			fmt.Println(storage.ShowBooks())
			fmt.Println()
		case "Добавь":
			var (
				name, author string
				release      int
			)
			fmt.Println("Введите информацию о книги, которую надо добавить:\n1)Название\n2)Автор\n3)Год")
			fmt.Scan(&name, &author, &release)
			storage.AddBook(name, author, release)
			fmt.Println()
		case "Удали":
			var name string
			fmt.Scan("Покажи")
			fmt.Println("Введите название книги, которую надо удалить:")
			fmt.Scan(&name)
			storage.DeleteBook(name)
			fmt.Println()
		case "Поиск":
			var search string
			fmt.Println("Выберите фильтр поиска:\n1)Название\n2)Автор\n3)Год")
			fmt.Scan(&search)
			storage.SearchBooks(search)
			fmt.Println()
		case "Выход":
			return
		default:
			fmt.Println("Вы ввели что-то не то.")
			fmt.Println()
		}

	}
}
