package store

import "github.com/LezhnevSergei/news_aggregator/internal/app/models"

type NewsRepository struct {
	store *Store
}

func (r NewsRepository) Create(u *models.News) (*models.News, error) {
	if err := r.store.db.QueryRow(
		"INSERT INTO news (title, created_at) VALUES ($1, $2) RETURNING id",
		u.Title,
		u.CreatedAt,
	).Scan(&u.ID); err != nil {
		return u, err
	}

	return u, nil
}
