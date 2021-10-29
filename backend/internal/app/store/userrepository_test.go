package store_test

import (
	"github.com/LezhnevSergei/news_aggregator/internal/app/models"
	"github.com/LezhnevSergei/news_aggregator/internal/app/store"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestUserRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("news")

	date := time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local)

	n, err := s.News().Create(&models.News{
		Title:     "TestNews",
		CreatedAt: date,
	})

	assert.NoError(t, err)
	assert.NotNil(t, n)
}
