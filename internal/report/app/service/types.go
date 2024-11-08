package service

import "time"

type Click struct {
	ShortURL  string
	ClickedAt time.Time
}

type ClickCount struct {
	ShortURL string
	Count    int
}
