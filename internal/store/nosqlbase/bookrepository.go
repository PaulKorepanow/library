package nosqlbase

import (
	"github.com/go-redis/redis"
	"main/internal/model"
	"time"
)

const expirationTime = 5 * time.Minute

type BookRepository struct {
	rdb *redis.Client
}

func (r *BookRepository) Create(book *model.Book) error {
	if err := r.rdb.Set(book.Title, book, expirationTime).Err(); err != nil {
		return err
	}
	return nil
}

func (r *BookRepository) FindByTitle(title string) (*model.Book, error) {
	val, err := r.rdb.Get(title).Bytes()
	if err != nil {
		return nil, err
	}
	book := new(model.Book)
	if err := book.UnmarshalBinary(val); err != nil {
		return nil, err
	}
	return book, nil
}
