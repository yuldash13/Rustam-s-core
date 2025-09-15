package main

import "fmt"

type Ticket struct {
	id        int
	direction *Direction
}

func NewTicket(id int, direction *Direction) *Ticket {
	return &Ticket{
		id:        0,
		direction: direction,
	}
}

func (t *Ticket) ToStringTicket() string {
	return fmt.Sprintf("%v\n", *t.direction)
}
