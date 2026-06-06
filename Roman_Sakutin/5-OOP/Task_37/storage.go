package main

import (
	"fmt"
)

type Storage struct {
	storage1 map[string][]Book
	storage2 map[string][]Book
	storage3 map[int][]Book
}

func NewStorage() *Storage {
	return &Storage{
		storage1: map[string][]Book{},
		storage2: map[string][]Book{},
		storage3: map[int][]Book{},
	}
}

func (s *Storage) InitStorage() {
	book1 := NewBook("MyWar", "ErenEger", 2014)
	book2 := NewBook("FairyTail", "NatsuDragnil", 2003)
	book3 := NewBook("Naruto", "NarutoUzumaki", 2001)

	bookArr1 := NewBookArr(*book1)
	bookArr2 := NewBookArr(*book2)
	bookArr3 := NewBookArr(*book3)

	s.storage1 = map[string][]Book{
		book1.name: bookArr1,
		book2.name: bookArr2,
		book3.name: bookArr3,
	}
	s.storage2 = map[string][]Book{
		book1.author: bookArr1,
		book2.author: bookArr2,
		book3.author: bookArr3,
	}
	s.storage3 = map[int][]Book{
		book1.release: bookArr1,
		book2.release: bookArr2,
		book3.release: bookArr3,
	}
}

func (s *Storage) AddBook(name string, author string, release int) {
	book := NewBook(name, author, release)
	if _, ok := s.storage1[name]; ok {
		fmt.Println("Книга уже есть.")
		return
	}
	s.storage1[name] = append(s.storage1[name], *book)
	s.storage2[author] = append(s.storage2[author], *book)
	s.storage3[release] = append(s.storage3[release], *book)
	fmt.Println("Книга добавлена.")
}

func (s *Storage) DeleteBook(name string) {
	if _, ok := s.storage1[name]; !ok {
		fmt.Println("Такой книги нет.")
		return
	}
	book := s.storage1[name][0]
	delete(s.storage1, name)
	for i := 0; i < len(s.storage2[book.author]); i++ {
		if s.storage2[book.author][i].name == book.name {
			s.storage2[book.author] = append(s.storage2[book.author][:i], s.storage2[book.author][i+1:]...)
		}
	}
	for i := 0; i < len(s.storage3[book.release]); i++ {
		if s.storage3[book.release][i].name == book.name {
			s.storage3[book.release] = append(s.storage3[book.release][:i], s.storage3[book.release][i+1:]...)
		}
	}
	fmt.Println("Книга удалена.")
}

func (s *Storage) ShowBooks() string {
	var str string
	for key, _ := range s.storage1 {
		str += fmt.Sprintln(s.storage1[key][0].ToString())
	}
	return str
}

func (s *Storage) SearchBooks(search string) {
	switch search {
	case "Название":
		var name string
		fmt.Println("Введите название:")
		fmt.Scan(&name)
		s.SearchBooksName(name)
		return
	case "Автор":
		var author string
		fmt.Println("Введите автора:")
		fmt.Scan(&author)
		s.SearchBooksAuthor(author)
		return
	case "Год":
		var release int
		fmt.Println("Введите год:")
		fmt.Scan(&release)
		s.SearchBooksRelease(release)
		return
	}
}

func (s *Storage) SearchBooksName(name string) {
	if _, ok := s.storage1[name]; ok {
		fmt.Println(s.storage1[name][0].ToString())
		fmt.Println()
	}
}

func (s *Storage) SearchBooksAuthor(author string) {
	if _, ok := s.storage2[author]; ok {
		fmt.Println(s.ToStringStr(author))
		fmt.Println()
	}
}

func (s *Storage) SearchBooksRelease(release int) {
	if _, ok := s.storage3[release]; ok {
		fmt.Println(s.ToStringInt(release))
		fmt.Println()
	}
}

func (s *Storage) ToStringStr(stroke string) string {
	var str string
	for i := 0; i < len(s.storage2[stroke]); i++ {
		str += s.storage2[stroke][i].ToString()
		str += fmt.Sprintln()
	}
	return str
}

func (s *Storage) ToStringInt(release int) string {
	var str string
	for i := 0; i < len(s.storage3[release]); i++ {
		str += s.storage3[release][i].ToString()
		str += fmt.Sprintln()
	}
	return str
}
