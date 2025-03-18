package main

import "fmt"

type Multiply struct {
	a Expression
	b Expression
}

func NewMultiply(first, second Expression) *Multiply {
	return &Multiply{a: first, b: second}
}

func (m *Multiply) evaluate(x int) int {
	return m.a.evaluate(x) * m.b.evaluate(x)
}

func (m *Multiply) string() string {
	return fmt.Sprintf("(%v * %v)", m.a.string(), m.b.string())
}

func (m *Multiply) equals(exp Expression) bool {
	return exp.string() == m.string()
}
