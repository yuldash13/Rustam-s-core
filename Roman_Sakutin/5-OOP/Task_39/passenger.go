package main

type Passenger struct {
	name   string
	ticket *Ticket
}

func NewPassenger(name string, ticket *Ticket) *Passenger {
	return &Passenger{
		name:   name,
		ticket: ticket,
	}
}
