package main

import (
	"fmt"
	"time"
)

type Book struct {
	title   string
	author  string
	pages   int
	isSaved bool
	savedAt time.Time
}

// 1. Write Data
func (b Book) saveBookBurned() {
	b.isSaved = true
	b.savedAt = time.Now()
	// book copy burned in the func Stack, the original book is not affected
}

// to Not burn book in the func Stack, we use pointer to the Book "*Book"
func (b *Book) saveBook() {
	b.isSaved = true
	b.savedAt = time.Now()
}

// 2. Read Data
func (b Book) getBookInfo() string {
	return fmt.Sprintf("Title: %s, Author: %s, Pages: %d, IsSaved: %t, SavedAt: %s",
		b.title,
		b.author,
		b.pages,
		b.isSaved,
		b.savedAt,
	)
}
