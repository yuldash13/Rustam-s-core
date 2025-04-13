package main

import "fmt"

type Cheater struct {
	name string
	hand []*Card
}

func NewCheater(name string) *Cheater {
	return &Cheater{
		name: name,
		hand: nil,
	}
}

func (c *Cheater) GetCards(cards []*Card) {
	c.hand = append(c.hand, cards...)
	c.SwapCards()
}

func (c *Cheater) SwapCards() {
	cards := []*Card{
		NewCard("ace"),
		NewCard("ace"),
		NewCard("ace"),
		NewCard("ace"),
		NewCard("ace"),
	}
	c.hand = cards
}

func (c *Cheater) ToString() string {
	var hand string
	for i := 0; i < len(c.hand); i++ {
		hand += fmt.Sprintf("%v)%v\n", i+1, c.hand[i].ToString())
	}
	return fmt.Sprintf("Gamer:\n-%v\nHand:\n%v", c.name, hand)
}
