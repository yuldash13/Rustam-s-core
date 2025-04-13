package main

import "fmt"

type Card struct {
	Name string
}

func NewCard(name string) *Card {
	return &Card{Name: name}
}

func (c *Card) ToString() string {
	return fmt.Sprintf("%v", c.Name)
}
