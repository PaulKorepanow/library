package storage

type Store interface {
	NewStore(datbaseURL string) Store
}
