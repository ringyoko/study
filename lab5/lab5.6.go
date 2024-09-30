package main

import (
	"fmt"
)

type Stringer interface {
	String() string
}

type Book struct {
	Title  string
	Author string
}

func (b Book) String() string {
	return fmt.Sprintf("Название книги: %s, автор: %s", b.Title, b.Author)
}

func main() {

	book1 := Book{Title: "Преступление и наказание", Author: "Ф.М.Достоевский"}
	book2 := Book{Title: "Люди как боги", Author: "С.А.Снегов"}

	fmt.Println(book1)
	fmt.Println(book2)
}
