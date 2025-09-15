package main

import (
	"math/rand"
	"strconv"
)

type Dispatcher struct {
	name   string
	trains []*Train
}

func NewDispatcher(name string) *Dispatcher {
	return &Dispatcher{
		name:   name,
		trains: nil,
	}
}

func (d *Dispatcher) MakeTrain(name string) string {
	direction := d.MakeDirection(name)
	quantity, tickets := d.MakeTicket(direction)
	passengers := d.SellTickets(quantity, tickets)
	carts := d.MakeCart(passengers)
	train := NewTrain(carts, quantity, tickets)
	d.trains = append(d.trains, train)
	return train.ToStringTrain()
}

func (d *Dispatcher) MakeDirection(name string) *Direction {
	direction := NewDirection(name)
	return direction
}

func (d *Dispatcher) SellTickets(quantity int, tickets []*Ticket) []*Passenger {
	passengers := make([]*Passenger, quantity)
	for i := 0; i < len(passengers); i++ {
		str := "Passenger" + strconv.Itoa(i+1)
		passengers[i] = NewPassenger(str, tickets[i])
	}
	return passengers
}

func (d *Dispatcher) MakeTicket(direction *Direction) (int, []*Ticket) {
	quantity := rand.Intn(21) + 20
	var tickets []*Ticket
	for i := 0; i < quantity; i++ {
		ticket := NewTicket(i+1, direction)
		tickets = append(tickets, ticket)
	}
	return quantity, tickets
}

func (d *Dispatcher) SaveCarts(passengers []*Passenger) []*Cart {
	var carts []*Cart
	carts = append(carts, d.MakeCart(passengers)...)
	return carts
}

func (d *Dispatcher) MakeCart(passengers []*Passenger) []*Cart {
	places := rand.Intn(11) + 10
	var carts []*Cart
	if len(passengers) > places {
		cart := NewCart(places)
		cart.passengers = append(cart.passengers, passengers[:places]...)
		carts = append(carts, cart)
		passengers = passengers[places:]
		carts = append(carts, d.MakeCart(passengers)...)
		return carts
	}
	cart := NewCart(places)
	cart.passengers = append(cart.passengers, passengers...)
	carts = append(carts, cart)
	return carts
}
