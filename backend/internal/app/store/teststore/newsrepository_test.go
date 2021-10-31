package teststore_test

import (
	"github.com/LezhnevSergei/news_aggregator/internal/app/models"
	"github.com/LezhnevSergei/news_aggregator/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewsRepository_Create(t *testing.T) {
	s := teststore.New()
	n := models.TestNews()

	assert.NoError(t, s.News().Create(n))
	assert.NotNil(t, n)
}

func TestNewsRepository_Get(t *testing.T) {
	s := teststore.New()

	n := models.TestNews()
	n2 := models.TestNews()

	s.News().Create(n)
	s.News().Create(n2)

	news, err := s.News().GetList()

	assert.NoError(t, err)
	assert.NotNil(t, news)
	assert.Equal(t, len(*news), 2)
}

func TestNewsRepository_GetNewsByTitle(t *testing.T) {
	s := teststore.New()

	n := models.TestNews()

	s.News().Create(n)

	searched1, err := s.News().GetNewsByTitle("invalid")
	searched2, err := s.News().GetNewsByTitle("Test")

	assert.NoError(t, err)
	assert.Empty(t, searched1)
	assert.NotEmpty(t, searched2)
}
