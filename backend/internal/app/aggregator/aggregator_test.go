package aggregator_test

import (
	"github.com/LezhnevSergei/news_aggregator/internal/app/aggregator"
	"github.com/LezhnevSergei/news_aggregator/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
	"time"
)

func TestNewsAggregator_SaveNews(t *testing.T) {
	s := teststore.New()
	na := aggregator.NewNewsAggregator("https://lenta.ru/rss/news", new(sync.Mutex))

	go na.SaveNews(s, 5)
	time.Sleep(7 * time.Second)

	na.Mu.Lock()
	news, err := s.News().GetList()
	na.Mu.Unlock()

	assert.NoError(t, err)
	assert.NotNil(t, news)
	assert.NotEqual(t, len(*news), 0)
}
