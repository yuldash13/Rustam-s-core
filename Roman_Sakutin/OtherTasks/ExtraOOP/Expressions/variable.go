package main

import "fmt"

type Variable struct {
	v string
}

func NewVariable(value string) *Variable {
	return &Variable{v: value}
}

func (v *Variable) evaluate(x int) int {
	return x
}

func (v *Variable) string() string {
	return fmt.Sprintf("%v", v.v)
}

func (v *Variable) equals(exp Expression) bool {
	va, ok := exp.(*Variable)
	if !ok {
		return false
	}
	return v.v == va.v
}
