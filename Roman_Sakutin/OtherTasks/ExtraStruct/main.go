package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Player struct {
	name   string
	hp     int
	damage int
}

func main() {
	player1 := Player{"Yuldash", 100, 30}
	player2 := Player{"Rustam", 100, 20}

	rand.Seed(time.Now().UnixNano())

	for {
		randomNum := rand.Intn(2)
		if randomNum == 0 {
			player1.punch(&player2)
		} else if randomNum == 1 {
			player2.punch(&player1)
		}

		if player1.hp < 0 {
			fmt.Println("Yuldash noob")
			break
		} else if player2.hp < 0 {
			fmt.Println("Rustam noob")
			break
		}
	}
}

func (p *Player) punch(player *Player) {
	player.hp -= p.damage
}
