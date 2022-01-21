package v1

import "gin/models"

type Manager struct {
	ID        uint            `json:"id"`
	Name      string          `json:"name"`
	Email     string          `json:"email"`
	Phone     string          `json:"phone"`
	Pwd       string          `json:"-"`
	CreatedAt models.GormTime `json:"created_at"`
	UpdatedAt models.GormTime `json:"updated_at"`
}
