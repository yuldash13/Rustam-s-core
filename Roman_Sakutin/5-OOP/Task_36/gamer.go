package main

import "fmt"

type Gamer struct {
	name string
	hand []*Card
}

func NewPlayer(name string) *Gamer {
	return &Gamer{
		name: name,
		hand: nil,
	}
}

func (g *Gamer) GetCards(cards []*Card) {
	g.hand = append(g.hand, cards...)
}

func (g *Gamer) ToString() string {
	var hand string
	for i := 0; i < len(g.hand); i++ {
		hand += fmt.Sprintf("%v)%v\n", i+1, g.hand[i].ToString())
	}
	return fmt.Sprintf("Gamer:\n-%v\nHand:\n%v", g.name, hand)
}
