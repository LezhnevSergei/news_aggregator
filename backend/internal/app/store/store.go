package store

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Store struct {
	config         *Config
	db             *sql.DB
	newsRepository *NewsRepository
}

func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.config.DataBaseURL)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db

	return nil
}

func (s *Store) Close() {
	s.db.Close()
}

func (s *Store) News() *NewsRepository {
	if s.newsRepository != nil {
		return s.newsRepository
	}

	s.newsRepository = &NewsRepository{
		store: s,
	}

	return s.newsRepository
}
