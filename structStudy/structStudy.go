package structStudy

import "fmt"

type Book struct {
	title       string
	auth        string
	bookContext string
}

func changeBookContextTest(book Book) {
	book.bookContext = "NO Context"
}

func changeBookContext(book *Book) {
	book.bookContext = "NO Context"
}

func StructStudy() {
	var theGreatGatsby Book
	theGreatGatsby.title = "the great gatsby"
	theGreatGatsby.auth = "Francis Scott Key Fitzgerald"
	theGreatGatsby.bookContext = "Too long to write"
	fmt.Printf("%v\n", theGreatGatsby)

	changeBookContextTest(theGreatGatsby)
	fmt.Printf("%v\n", theGreatGatsby)

	changeBookContext(&theGreatGatsby)
	fmt.Printf("%v\n", theGreatGatsby)
}

func init() {
	fmt.Println("structStudy.init()")
}
