package main

import "fmt"

type Add struct {
	a Expression
	b Expression
}

func NewAdd(first, second Expression) *Add {
	return &Add{a: first, b: second}
}

func (a *Add) evaluate(x int) int {
	return a.a.evaluate(x) + a.b.evaluate(x)
}

func (a *Add) string() string {
	return fmt.Sprintf("(%v + %v)", a.a.string(), a.b.string())
}

func (a *Add) equals(exp Expression) bool {
	add, ok := exp.(*Add)
	if !ok {
		return false
	}
	return a.a.equals(add.a) == a.b.equals(add.b)
}
