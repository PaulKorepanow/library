package store

import "main/internal/model"

type BookRepository interface {
	Create(book *model.Book) error
	FindBookByTitle(title string) (*model.Book, error)
	DeleteBookByTitle(title string) error
}
