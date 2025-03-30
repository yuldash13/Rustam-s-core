package main

import "fmt"

type Player struct {
	ID       int
	nickName string
	level    int
	banned   bool
}

func main() {
	data()
	for {
		var menu string
		fmt.Println("Меню:\n1)Показать\n2)Забанить\n3)Разбанить\n4)Добавить\n5)Удалить\n6)Выход")
		fmt.Scan(&menu)
		switch menu {
		case "Показать":
			showDataBase()
		case "Забанить":
			var id int
			showDataBase()
			fmt.Println("Кого забанить? Введите ID:")
			fmt.Scan(&id)
			banPlayer(id)
			fmt.Println()
		case "Разбанить":
			var id int
			showDataBase()
			fmt.Println("Кого разбанить? Введите ID:")
			fmt.Scan(&id)
			unBanPlayer(id)
			fmt.Println()
		case "Добавить":
			var nick string
			fmt.Println("Введите ник игрока:")
			fmt.Scan(&nick)
			addPlayer(nick)
			fmt.Println()
		case "Удалить":
			var id int
			fmt.Println("Кого удалить? Введите ID:")
			showDataBase()
			fmt.Scan(&id)
			deletePlayer(id)
			fmt.Println()
		case "Выход":
			return
		default:
			fmt.Println("Вы ввели что-то не то")
			fmt.Println()
		}
	}

}

func addPlayer(nick string) *dataBase {
	player := Player{
		ID:       len(playerList),
		nickName: nick,
		level:    1,
		banned:   false,
	}
	playerList[len(playerList)] = player
	fmt.Println("Игрок добавлен")
	return &dataBase{playerList[len(playerList)]}
}

func deletePlayer(id int) {
	_, exist := playerList[id]
	if !exist {
		fmt.Println("Такого игрока нет")
		return
	}
	fmt.Println("Игрок №", id, "удален")
	delete(playerList, id)
}

func banPlayer(id int) *dataBase {
	player := playerList[id]
	if player.banned == true {
		fmt.Println("Этот игрок уже забанен.")
		return nil
	}
	player.banned = true
	playerList[id] = player
	fmt.Println("Игрок №", id, "ЗАБАНЕН!!!!!")
	return &dataBase{playerList[id]}
}

func unBanPlayer(id int) *dataBase {
	player := playerList[id]
	if player.banned == false {
		fmt.Println("Этот игрок не забанен.")
		return nil
	}
	player.banned = false
	playerList[id] = player
	fmt.Println("Игрок №", id, "разбанен")
	return &dataBase{playerList[id]}
}

func showDataBase() {
	for i := 0; i < len(playerList); i++ {
		fmt.Println("Игрок №", i)
		fmt.Println("ID:", playerList[i].ID)
		fmt.Println("Nick Name:", playerList[i].nickName)
		fmt.Println("Level:", playerList[i].level)
		fmt.Println("Banned:", playerList[i].banned)
		fmt.Println()
	}
}
