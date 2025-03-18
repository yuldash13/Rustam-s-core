package main

import "fmt"

type Subtract struct {
	a Expression
	b Expression
}

func NewSubtract(first, second Expression) *Subtract {
	return &Subtract{a: first, b: second}
}

func (s *Subtract) evaluate(x int) int {
	return s.a.evaluate(x) - s.b.evaluate(x)
}

func (s *Subtract) string() string {
	return fmt.Sprintf("(%v - %v)", s.a.string(), s.b.string())
}

func (s *Subtract) equals(exp Expression) bool {
	return exp.string() == s.string()
}
