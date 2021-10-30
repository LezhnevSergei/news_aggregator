package sqlstore_test

import (
	"github.com/LezhnevSergei/news_aggregator/internal/app/models"
	"github.com/LezhnevSergei/news_aggregator/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewsRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("news")

	s := sqlstore.New(db)
	n := models.TestNews()

	assert.NoError(t, s.News().Create(n))
	assert.NotNil(t, n)
}
