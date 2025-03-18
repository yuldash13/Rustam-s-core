package main

import "fmt"

type Const struct {
	c int
}

func NewConst(value int) *Const {
	return &Const{c: value}
}

func (c *Const) evaluate(_ int) int {
	return c.c
}

func (c *Const) string() string {
	return fmt.Sprintf("%v", c.c)
}

func (c *Const) equals(_ Expression) bool {
	return c.c == c.c
}
