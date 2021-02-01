package teststore

import (
	"fmt"
	"main/internal/model"
)

type bookRepository map[string]*model.Book

func (r bookRepository) Create(book *model.Book) error {
	if _, ok := r[book.Title]; ok {
		return fmt.Errorf("book with same title already defined")
	}
	r[book.Title] = book
	return nil
}

func (r bookRepository) FindBookByTitle(title string) (*model.Book, error) {
	if _, ok := r[title]; !ok {
		return nil, fmt.Errorf("book with title(%s) not defined", title)
	}
	return r[title], nil
}

func (r bookRepository) DeleteBookByTitle(title string) error {
	if _, ok := r[title]; ok {
		delete(r, title)
	}
	return nil
}
