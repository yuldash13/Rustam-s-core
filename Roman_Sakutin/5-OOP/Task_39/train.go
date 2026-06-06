package main

import "fmt"

type Train struct {
	carts    []*Cart
	quantity int
	tickets  []*Ticket
}

func NewTrain(carts []*Cart, quantity int, tickets []*Ticket) *Train {
	return &Train{
		carts:    carts,
		quantity: quantity,
		tickets:  tickets,
	}
}

func (t *Train) ToStringTrain() string {
	var str string
	str = t.tickets[0].ToStringTicket()
	for i := 0; i < len(t.carts); i++ {
		str += fmt.Sprintf(t.carts[i].ToStringCart())
	}
	return str
}
