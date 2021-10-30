package teststore

import (
	"github.com/LezhnevSergei/news_aggregator/internal/app/models"
	"github.com/LezhnevSergei/news_aggregator/internal/app/store"
)

type Store struct {
	newsRepository *NewsRepository
}

func New() *Store {
	return &Store{}
}

func (s *Store) News() store.NewsRepository {
	if s.newsRepository != nil {
		return s.newsRepository
	}

	s.newsRepository = &NewsRepository{
		store: s,
		news:  make(map[int]*models.News),
	}

	return s.newsRepository
}
