package sqlstore

import (
	"github.com/LezhnevSergei/news_aggregator/internal/app/models"
)

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

// CreateList I tried to do it with one request, but I couldn't :(
func (r NewsRepository) CreateList(news *[]models.News) {
	for _, n := range *news {
		if err := r.Create(&n); err != nil {
			continue
		}
	}
}

func (r NewsRepository) GetList() (*[]models.News, error) {
	var newsList []models.News
	rows, err := r.store.db.Query("SELECT * FROM news ORDER BY created_at DESC")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var news models.News
		if err := rows.Scan(&news.ID, &news.Title, &news.CreatedAt); err != nil {
			return nil, err
		}

		newsList = append(newsList, news)
	}

	return &newsList, nil
}
