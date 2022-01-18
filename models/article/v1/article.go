package v1

import (
	"gin/models"
)

type Article struct {
	ID        uint            `json:"id"`
	UserId    uint            `json:"user_id"`
	Name      string          `json:"name"`
	Title     string          `json:"title"`
	Content   string          `json:"content"`
	CreatedAt models.GormTime `json:"created_at"`
	UpdatedAt models.GormTime `json:"updated_at"`
}
