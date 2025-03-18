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

func (v *Variable) equals(_ Expression) bool {
	return v.v == v.v
}
