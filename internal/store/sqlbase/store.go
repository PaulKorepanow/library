package sqlbase

import (
	"github.com/jinzhu/gorm"
	"main/internal/store"
)

type SqlStore struct {
	db             *gorm.DB
	bookRepository *bookRepository
}

func NewSqlStore(db *gorm.DB) *SqlStore {
	return &SqlStore{
		db: db,
	}
}

func (s *SqlStore) BookRepository() store.BookRepository {
	if s.bookRepository == nil {
		s.bookRepository = &bookRepository{db: s.db}
	}

	return s.bookRepository
}
