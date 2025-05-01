package main

type Seller struct {
	name        string
	superMarket *SuperMarket
}

func NewSeller(name string, superMarket *SuperMarket) *Seller {
	return &Seller{
		name:        name,
		superMarket: superMarket,
	}
}

func (s *Seller) NextClient() {
	client := s.superMarket.queue.Dequeue()
	s.superMarket.money += client.PayProducts()
}
