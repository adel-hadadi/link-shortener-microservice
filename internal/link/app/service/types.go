package service

import "time"

type Link struct {
	ID           int
	OriginalLink string
	ShortLink    string
	CreatedAt    time.Time
}
