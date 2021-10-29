package models

import "time"

type News struct {
	ID        int
	Title     string
	CreatedAt time.Time
}
