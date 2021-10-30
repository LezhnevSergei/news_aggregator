package sqlstore

import (
	"database/sql"
	"github.com/LezhnevSergei/news_aggregator/internal/app/store"

	_ "github.com/lib/pq"
)

type Store struct {
	db             *sql.DB
	newsRepository *NewsRepository
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) News() store.NewsRepository {
	if s.newsRepository != nil {
		return s.newsRepository
	}

	s.newsRepository = &NewsRepository{
		store: s,
	}

	return s.newsRepository
}
