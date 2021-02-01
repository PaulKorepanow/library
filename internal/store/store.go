package store

type Store interface {
	BookRepository() BookRepository
}
