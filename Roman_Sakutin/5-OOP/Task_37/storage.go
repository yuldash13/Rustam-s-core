package main

import (
	"fmt"
)

type Storage struct {
	storage map[string]*Book
}

func NewStorage() *Storage {
	return &Storage{storage: map[string]*Book{}}
}

func (s *Storage) InitStorage() {
	book1 := NewBook("MyWar", "ErenEger", 2014)
	book2 := NewBook("FairyTail", "NatsuDragnil", 2003)
	book3 := NewBook("Naruto", "NarutoUzumaki", 2001)

	s.storage = map[string]*Book{
		book1.name: book1,
		book2.name: book2,
		book3.name: book3,
	}
}

func (s *Storage) AddBook(name string, author string, data int) {
	for key, _ := range s.storage {
		if s.storage[name] == s.storage[key] {
			fmt.Println("Книга уже есть.")
			return
		}
	}
	s.storage[name] = NewBook(name, author, data)
	fmt.Println("Книга добавлена.")
}

func (s *Storage) DeleteBook(name string) {
	for key, _ := range s.storage {
		if s.storage[name] == s.storage[key] {
			delete(s.storage, name)
			fmt.Println("Книга удалена.")
			return
		}
	}
	fmt.Println("Такой книги нет.")
}

func (s *Storage) ShowBooks() {
	for key, _ := range s.storage {
		fmt.Println(s.storage[key].ToString())
	}
}

func (s *Storage) SearchBooks(search string) {
	if search == "Название" {
		var name string
		fmt.Println("Введите название:")
		fmt.Scan(&name)
		s.SearchBooksName(name)
		return
	} else if search == "Автор" {
		var author string
		fmt.Println("Введите автора:")
		fmt.Scan(&author)
		s.SearchBooksAuthor(author)
		return
	} else if search == "Год" {
		var release int
		fmt.Println("Введите год:")
		fmt.Scan(&release)
		s.SearchBooksRelease(release)
		return
	}
}

func (s *Storage) SearchBooksName(name string) {
	for key, _ := range s.storage {
		if name == key {
			fmt.Println(s.storage[key].ToString())
			fmt.Println()
		}
	}
}

func (s *Storage) SearchBooksAuthor(author string) {
	for key, _ := range s.storage {
		if author == s.storage[key].author {
			fmt.Println(s.storage[key].ToString())
			fmt.Println()
		}
	}
}

func (s *Storage) SearchBooksRelease(release int) {
	for key, _ := range s.storage {
		if release == s.storage[key].release {
			fmt.Println(s.storage[key].ToString())
			fmt.Println()
		}
	}
}
