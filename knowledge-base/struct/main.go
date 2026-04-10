package main

import (
	"fmt"
	"time"
)

func main() {
	book1 := Book{
		title:   "The Great Gatsby",
		author:  "",
		pages:   180,
		isSaved: false,
		savedAt: time.Time{},
	}
	book1.saveBookBurned()
	fmt.Println(book1.getBookInfo())
	book1.saveBook()
	fmt.Println(book1.getBookInfo())
}
