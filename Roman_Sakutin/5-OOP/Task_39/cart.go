package main

import "fmt"

type Cart struct {
	places     int
	passengers []*Passenger
}

func NewCart(places int) *Cart {
	return &Cart{
		places:     places,
		passengers: nil,
	}
}

func (c *Cart) ToStringCart() string {
	return fmt.Sprintf("cart has %d places, %d passengers\n", c.places, len(c.passengers))
}
