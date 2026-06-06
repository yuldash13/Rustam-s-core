package main

import (
	"fmt"
)

type Book struct {
	name    string
	author  string
	release int
}

func NewBook(name string, author string, release int) *Book {
	return &Book{
		name:    name,
		author:  author,
		release: release,
	}
}

func NewBookArr(book Book) []Book {
	return []Book{book}
}

func (b *Book) ToString() string {
	return fmt.Sprintf("Name: %v\n1) Author: %v\n2) Release year: %v\n", b.name, b.author, b.release)
}
