package main

import "math/rand"

type Deck struct {
	cards []*Card
}

func NewDeck() *Deck {
	return &Deck{cards: []*Card{}}
}

func (d *Deck) AddCard(card *Card) {
	for i := 0; i < 4; i++ {
		d.cards = append(d.cards, card)
	}
}

func (d *Deck) GetCards(i int) []*Card {
	cards := append([]*Card{}, d.cards[0:i]...)
	d.cards = d.cards[i:]
	return cards
}

func (d *Deck) RandomDeck() {
	for i := 0; i < len(d.cards); i++ {
		j := rand.Intn(len(d.cards))
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	}
}

func (d *Deck) InitDeck() {
	d.AddCard(NewCard("six"))
	d.AddCard(NewCard("seven"))
	d.AddCard(NewCard("eight"))
	d.AddCard(NewCard("nine"))
	d.AddCard(NewCard("ten"))
	d.AddCard(NewCard("knight"))
	d.AddCard(NewCard("queen"))
	d.AddCard(NewCard("king"))
	d.AddCard(NewCard("ace"))
}
