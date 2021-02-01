package nosqlbase

import "github.com/go-redis/redis"

type NoSqlStore struct {
	db             *redis.Client
	bookRepository *BookRepository
}

func NewNoSqlStore() *NoSqlStore {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return &NoSqlStore{
		db: rdb,
	}
}

func (s *NoSqlStore) BookRepository() *BookRepository {
	if s.bookRepository != nil {
		return s.bookRepository
	}
	rep := &BookRepository{
		rdb: s.db,
	}
	s.bookRepository = rep
	return s.bookRepository
}
