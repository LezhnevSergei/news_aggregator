package store

type Store interface {
	News() NewsRepository
}
