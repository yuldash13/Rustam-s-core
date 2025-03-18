package main

import "fmt"

type Divide struct {
	a Expression
	b Expression
}

func NewDivide(first, second Expression) *Divide {
	return &Divide{a: first, b: second}
}

func (d *Divide) evaluate(x int) int {
	return d.a.evaluate(x) / d.b.evaluate(x)
}

func (d *Divide) string() string {
	return fmt.Sprintf("(%v / %v)", d.a.string(), d.b.string())
}

func (d *Divide) equals(exp Expression) bool {
	return exp.string() == d.string()
}
