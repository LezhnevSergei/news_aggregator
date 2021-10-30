package sqlstore

import "github.com/LezhnevSergei/news_aggregator/internal/app/models"

type NewsRepository struct {
	store *Store
}

func (r NewsRepository) Create(n *models.News) error {
	return r.store.db.QueryRow(
		"INSERT INTO news (title, created_at) VALUES ($1, $2) RETURNING id",
		n.Title,
		n.CreatedAt,
	).Scan(&n.ID)
}
