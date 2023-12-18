package classStudy

import "fmt"

type Book struct {
	Title      string
	Auth       string
	BookContxt string
}

func (book *Book) GetTitle() string {
	return book.Title
}

func (book *Book) setContext(context string) {
	book.BookContxt = context
}
func (book *Book) ShowBook() {
	fmt.Println("title =", book.Title)
	fmt.Println("auth =", book.Auth)
	fmt.Println("bookContxt =", book.BookContxt)
}
func ClassStudy() {
	theGreatGatsby := Book{Title: "the great gatsby", Auth: "Francis Scott Key Fitzgerald", BookContxt: "Too long to write"}
	theGreatGatsby.setContext("DO NOT FINISH")
	theGreatGatsby.ShowBook()
}
func init() {
	fmt.Println("classStudy.init()")
}
