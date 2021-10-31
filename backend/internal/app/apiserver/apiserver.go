package apiserver

import (
	"database/sql"
	"github.com/LezhnevSergei/news_aggregator/internal/app/aggregator"
	"github.com/LezhnevSergei/news_aggregator/internal/app/store/sqlstore"
	"net/http"
	"sync"
)

func Start(config *Config) error {
	db, err := newDB(config.DatabaseURL)
	if err != nil {
		return err
	}

	defer db.Close()

	store := sqlstore.New(db)
	srv := NewServer(store)

	na := aggregator.NewNewsAggregator("https://lenta.ru/rss/news", new(sync.Mutex))
	go na.SaveNews(store, 2)

	return http.ListenAndServe(config.BindAddr, srv)
}

func newDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
