package v1

import "time"

type Article struct {
	ID        uint
	Name      string
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
