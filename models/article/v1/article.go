package v1

import (
	"gin/models"
)

type Article struct {
	ID        uint `json: "id"`
	UserId    uint `gorm:"column:user_id"`
	Name      string
	Title     string
	Content   string
	Date      models.GormTime
	CreatedAt models.GormTime
	UpdatedAt models.GormTime
}
