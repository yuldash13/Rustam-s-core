package main

import "fmt"

type DB struct {
	playerList map[int]*Player
}

func NewDB() *DB {
	return &DB{playerList: map[int]*Player{}}
}

func (db *DB) InitDB() {
	db.playerList = map[int]*Player{
		0: NewPlayer(0, "Yuldash", 999, false),
		1: NewPlayer(1, "Rustank", 888, false),
		2: NewPlayer(2, "Chiter777", 777, true),
	}
}

func (db *DB) AddPlayer(nick string) {
	player := NewPlayer(len(db.playerList), nick, 1, false)
	db.playerList[len(db.playerList)] = player
	fmt.Println("Игрок добавлен")
}

func (db *DB) DeletePlayer(id int) {
	_, exist := db.playerList[id]
	if !exist {
		fmt.Println("Такого игрока нет")
		return
	}
	delete(db.playerList, id)
	fmt.Println("Игрок №", id, "удален")
}

func (db *DB) BanPlayer(id int) {
	if db.ShowPlayer(id).ShowBan() == true {
		fmt.Println("Этот игрок уже забанен.")
		return
	}
	db.ShowPlayer(id).BanPlayer()
	fmt.Println("Игрок №", id, "ЗАБАНЕН!!!!!")
}

func (db *DB) UnBanPlayer(id int) {
	if db.ShowPlayer(id).ShowBan() == false {
		fmt.Println("Этот игрок не забанен.")
		return
	}
	db.ShowPlayer(id).UnPlayer()
	fmt.Println("Игрок №", id, "разбанен")
}

func (db *DB) ShowDataBase() {
	for i := 0; i < len(db.playerList); i++ {
		fmt.Println(db.playerList[i].ToString())
	}
}

func (db *DB) ShowPlayer(id int) *Player {
	return db.playerList[id]
}
