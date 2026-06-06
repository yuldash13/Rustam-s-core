package main

type Direction struct {
	name string
}

func NewDirection(name string) *Direction {
	return &Direction{name: name}
}
