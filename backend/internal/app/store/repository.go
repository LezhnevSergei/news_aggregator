package store

import "github.com/LezhnevSergei/news_aggregator/internal/app/models"

type NewsRepository interface {
	Create(news *models.News) error
}
