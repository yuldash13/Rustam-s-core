package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Player interface {
	GetCards(cards []*Card)
	ToString() string
}

func main() {
	rand.Seed(time.Now().UnixNano())

	deck := NewDeck()
	deck.InitDeck()

	players := []Player{
		NewPlayer("Yuldash"),
		NewPlayer("Rustam"),
		NewCheater("Cheater777"),
	}

	dealer := NewDealer(deck, players)

	dealer.ShakeDeck()

	for i := 0; i < len(dealer.players); i++ {
		dealer.TakePlayer(i).GetCards(dealer.GiveCards(5))
	}

	for i := 0; i < len(dealer.players); i++ {
		fmt.Println(dealer.TakePlayer(i).ToString())
	}
	
}
