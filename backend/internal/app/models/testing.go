package models

import (
	"time"
)

func TestNews() *News {
	date := time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local)

	return &News{
		Title:     "TestNews",
		CreatedAt: date,
	}
}
