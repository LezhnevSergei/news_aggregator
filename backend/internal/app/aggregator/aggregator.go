package aggregator

import (
	"github.com/LezhnevSergei/news_aggregator/internal/app/models"
	"github.com/LezhnevSergei/news_aggregator/internal/app/store"
	"github.com/jasonlvhit/gocron"
	"github.com/mmcdole/gofeed"
	"log"
	"sync"
	"time"
)

type NewsAggregator struct {
	Mu       *sync.Mutex
	ParseURL string
}

func NewNewsAggregator(parseURL string, mu *sync.Mutex) *NewsAggregator {
	return &NewsAggregator{ParseURL: parseURL, Mu: mu}
}

func (a *NewsAggregator) SaveNews(store store.Store, intervalInSec uint64) {
	fp := gofeed.NewParser()

	gocron.Every(intervalInSec).Seconds().Do(a.saveNews, fp, store)
	<-gocron.Start()
}

func (a *NewsAggregator) saveNews(fp *gofeed.Parser, store store.Store) {
	feed, err := fp.ParseURL(a.ParseURL)
	if err != nil {
		log.Fatal(err)
	}

	newsList, err := a.getNewNews(feed, store)
	if err != nil {
		log.Fatal(err)
	}

	a.Mu.Lock()
	store.News().CreateList(newsList)
	a.Mu.Unlock()
}

func (a *NewsAggregator) getNewNews(feed *gofeed.Feed, store store.Store) (*[]models.News, error) {
	news, err := store.News().GetList()
	if err != nil {
		return nil, err
	}
	fromTime := time.Now().UTC().Add(time.Hour * 2)
	if len(*news) > 0 {
		fromTime = ((*news)[0]).CreatedAt.UTC()
	}
	var newNews []models.News

	parsedNews := a.getParsedNews(feed)
	for _, n := range *parsedNews {
		if n.CreatedAt.UTC().Add(time.Hour * 3).After(fromTime) {
			newNews = append(newNews, n)
		}
	}

	return &newNews, err
}

func (a *NewsAggregator) getParsedNews(feed *gofeed.Feed) *[]models.News {
	var newsList []models.News

	for _, newsItem := range feed.Items {
		published, err := time.Parse("Mon, 02 Jan 2006 15:04:05 -0700", newsItem.Published)
		if err != nil {
			continue
		}
		news := models.News{
			Title:     newsItem.Title,
			CreatedAt: published,
		}
		a.Mu.Lock()
		newsList = append(newsList, news)
		a.Mu.Unlock()
	}

	return &newsList
}
