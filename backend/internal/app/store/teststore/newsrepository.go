package teststore

import "github.com/LezhnevSergei/news_aggregator/internal/app/models"

type NewsRepository struct {
	store *Store
	news  map[int]*models.News
}

func (r *NewsRepository) GetList() (*[]models.News, error) {
	var news []models.News
	for _, newsItem := range r.news {
		news = append(news, *newsItem)
	}

	return &news, nil
}

func (r *NewsRepository) Create(n *models.News) error {
	n.ID = len(r.news) + 1
	r.news[n.ID] = n

	return nil
}

func (r *NewsRepository) CreateList(news *[]models.News) {
	for _, n := range *news {
		n.ID = len(r.news) + 1
		r.news[n.ID] = &n
	}
}
