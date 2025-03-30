package main

type dataBase struct {
	Player
}

var playerList map[int]Player

func data() {
	playerList = map[int]Player{
		0: {
			ID:       0,
			nickName: "Yuldash13",
			level:    999,
			banned:   false,
		},
		1: {
			ID:       1,
			nickName: "Rustank",
			level:    888,
			banned:   false,
		},
		2: {
			ID:       2,
			nickName: "Читер777",
			level:    777,
			banned:   true,
		},
	}
}
