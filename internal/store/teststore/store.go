package teststore

import (
	"main/internal/model"
	"main/internal/store"
)

type TestStore struct {
	bookRepository bookRepository
}

func NewTestStore() *TestStore {
	return &TestStore{
		bookRepository: make(bookRepository),
	}
}

func (s *TestStore) BookRepository() store.BookRepository {
	if s.bookRepository == nil {
		s.bookRepository = make(map[string]*model.Book)
	}

	return s.bookRepository
}
