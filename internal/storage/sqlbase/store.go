package sqlbase

import "github.com/jinzhu/gorm"

type SqlStore struct {
	db             *gorm.DB
	bookRepository *bookRepository
}

func NewSqlStore(db *gorm.DB) *SqlStore {
	return &SqlStore{
		db: db,
	}
}

func (s *SqlStore) Book() *bookRepository {
	if s.bookRepository == nil {
		s.bookRepository = &bookRepository{db: s.db}
	}

	return s.bookRepository
}
