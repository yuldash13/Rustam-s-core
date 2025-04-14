package main

import "fmt"

func main() {
	dataBase := NewDB()
	dataBase.InitDB()

	for {
		var menu string
		fmt.Println("Меню:\n1)Показать\n2)Забанить\n3)Разбанить\n4)Добавить\n5)Удалить\n6)Выход")
		fmt.Scan(&menu)
		switch menu {
		case "Показать":
			dataBase.ShowDataBase()
		case "Забанить":
			var id int
			dataBase.ShowDataBase()
			fmt.Println("Кого забанить? Введите ID:")
			fmt.Scan(&id)
			dataBase.BanPlayer(id)
			fmt.Println()
		case "Разбанить":
			var id int
			dataBase.ShowDataBase()
			fmt.Println("Кого разбанить? Введите ID:")
			fmt.Scan(&id)
			dataBase.UnBanPlayer(id)
			fmt.Println()
		case "Добавить":
			var nick string
			fmt.Println("Введите ник игрока:")
			fmt.Scan(&nick)
			dataBase.AddPlayer(nick)
			fmt.Println()
		case "Удалить":
			var id int
			fmt.Println("Кого удалить? Введите ID:")
			dataBase.ShowDataBase()
			fmt.Scan(&id)
			dataBase.DeletePlayer(id)
			fmt.Println()
		case "Выход":
			return
		default:
			fmt.Println("Вы ввели что-то не то")
			fmt.Println()
		}
	}

}
