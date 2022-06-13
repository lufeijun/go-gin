package pgsql

import "gin/models"

type User struct {
	Id        uint64          `json:"id"`
	Name      string          `json:"name"`
	Age       uint16          `json:"age"`
	Address   string          `json:"address"`
	CreatedAt models.GormTime `json:"created_at"`
	UpdatedAt models.GormTime `json:"updated_at"`
}

func (User) TableName() string {
	return "users"
}
