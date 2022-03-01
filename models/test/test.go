package test

import (
	"gin/models"
	"time"

	orm "gin/database"
)

type Test0301 struct {
	ID   uint            `json:"id"`
	Name string          `json:"name"`
	Time models.GormTime `json:"created_at"`
}

func (Test0301) TableName() string {
	return "test0301"
}

func UpdateTest0301Cron() {
	orm.GormDB.Model(&Test0301{}).Where("id", 1).Update("time", time.Now().Format("2006-01-02 15:04:05"))
}
