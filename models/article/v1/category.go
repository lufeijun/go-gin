package v1

import "gin/models"

type Category struct {
	ID        uint               `json:"id"`
	Name      string             `json:"name"`
	Level     int                `json:"level"`
	ParentId  int                `json:"parent_id"`
	CreatedAt models.GormTime    `json:"created_at"`
	UpdatedAt models.GormTime    `json:"updated_at"`
	Childrens []CategoryChildren `gorm:"foreignKey:ParentId;" json:"childrens"`
}

func (Category) TableName() string {
	return "articles_category"
}

// 自己关联自己
type CategoryChildren struct {
	ID        uint            `json:"id"`
	Name      string          `json:"name"`
	Level     int             `json:"level"`
	ParentId  int             `json:"parent_id"`
	CreatedAt models.GormTime `json:"created_at"`
	UpdatedAt models.GormTime `json:"updated_at"`
}

func (CategoryChildren) TableName() string {
	return "articles_category"
	// return "articles_category_bak"
}
