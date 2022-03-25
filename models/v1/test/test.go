package test

import "gin/models"

type TestOne struct {
	ID   uint            `json:"id"`
	Date models.GormTime `json:"date"`
	// CreatedAt models.GormTime `json:"created_at"`
	// UpdatedAt models.GormTime `json:"updated_at"`
}

func (TestOne) TableName() string {
	return "test_one"
}
