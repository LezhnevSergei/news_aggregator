package teststore

import "github.com/LezhnevSergei/news_aggregator/internal/app/models"

type NewsRepository struct {
	store *Store
	news  map[int]*models.News
}

func (r *NewsRepository) Create(n *models.News) error {
	n.ID = len(r.news) + 1
	r.news[n.ID] = n

	return nil
}
