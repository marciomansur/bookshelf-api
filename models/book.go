package models

type Book struct {
	ID string `json:"id,omitempty"`

	Title string `json:"title,omitempty"`

	Description string `json:"description,omitempty"`

	Author *Author `json:"author,omitempty"`
}

type Author struct {
	Name string `json:"name,omitempty"`
}

// Memory bookshelf
var bookshelf []Book

// Insert a book to bookshelf (memory)
func Add(book Book) {
	bookshelf = append(bookshelf, book)
}

// Get all books
func Get() []Book {
	return bookshelf
}
