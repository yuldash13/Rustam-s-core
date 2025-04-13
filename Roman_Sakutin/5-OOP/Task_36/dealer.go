package main

type Dealer struct {
	deck    *Deck
	players []Player
}

func NewDealer(deck *Deck, players []Player) *Dealer {
	return &Dealer{
		deck:    deck,
		players: players,
	}
}

func (d *Dealer) GiveCards(i int) []*Card {
	cards := d.deck.GetCards(i)
	return cards
}

func (d *Dealer) ShakeDeck() {
	d.deck.RandomDeck()
}

func (d *Dealer) TakePlayer(i int) Player {
	return d.players[i]
}
