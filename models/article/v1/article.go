package v1

import my "gin/models"

type Article struct {
	ID      uint
	Name    string
	Title   string
	Content string
	// CreatedAt time.Time
	// UpdatedAt time.Time
	CreatedAt my.MyTime // 会导致时间不会自动更新
	UpdatedAt my.MyTime `gorm:"comment:'修改时间';type:datetime;"`
}
