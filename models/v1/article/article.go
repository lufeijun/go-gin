package article

import "gin/models"

type Article struct {
	ID               uint            `json:"id"`
	UserId           uint            `json:"user_id"`
	Name             string          `json:"name"`
	Title            string          `json:"title"`
	Content          string          `json:"content"`
	CategoryFirstId  int             `json:"category_first_id"`
	CategorySecondId int             `json:"category_second_id"`
	CreatedAt        models.GormTime `json:"created_at"`
	UpdatedAt        models.GormTime `json:"updated_at"`
	CategoryFirst    CategorySmart   `gorm:"foreignKey:CategoryFirstId;" json:"category_first"`
	CategorySecond   CategorySmart   `gorm:"foreignKey:CategorySecondId;" json:"category_second"`
}
