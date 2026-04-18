package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Book struct {
	title   string    `json:"title"`
	author  string    `json:"author"`
	pages   int       `json:"pages"`
	isSaved bool      `json:"isSaved"`
	savedAt time.Time `json:"savedAt"`
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

func (b Book) serializeBook() string {
	jsonData, _ := json.Marshal(b)
	return string(jsonData)
}
