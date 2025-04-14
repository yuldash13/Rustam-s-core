package main

import (
	"fmt"
)

type Book struct {
	name    string
	author  string
	release int
}

func NewBook(name string, author string, data int) *Book {
	return &Book{
		name:    name,
		author:  author,
		release: data,
	}
}

func (b *Book) ToString() string {
	return fmt.Sprintf("Name: %v\n1) Author: %v\n2) Release year: %v\n", b.name, b.author, b.release)
}
