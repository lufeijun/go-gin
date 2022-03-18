package article

import (
	"errors"
	"gin/database/orm"
	"gin/models/v1/article"
	"gin/structs"
	"math"

	"gorm.io/gorm"
)

// 文章部分

func List(id int, name string, page int64, pagesize int64) (pagestruct structs.PageStruct) {
	var articles []article.Article
	sql := orm.MysqlOrm

	// id 查询
	if id != 0 {
		sql = sql.Where("id", id)
	}

	if name != "" {
		sql = sql.Where("name like ?", "%"+name+"%")
		// sql = sql.Where("name like %?%", name)
	}

	var total int64
	sql.Model(&article.Article{}).Count(&total)
	pagestruct.LastPage = int64(math.Ceil(float64(total) / float64(pagesize)))

	result := sql.Debug().Scopes(orm.Paginate(page, pagesize)).Order("id desc").Find(&articles)
	if result.Error != nil {
		panic(result.Error)
	}

	// 赋值
	pagestruct.Total = total
	pagestruct.Page = page
	pagestruct.Size = pagesize
	pagestruct.Data = articles

	return
}

func Add(name, title, content string, category_first_id, category_second_id int) (id uint, err error) {
	article := article.Article{
		Name:             name,
		Title:            title,
		Content:          content,
		CategoryFirstId:  category_first_id,
		CategorySecondId: category_second_id,
		UserId:           1,
	}

	result := orm.MysqlOrm.Create(&article)

	if result.Error != nil {
		err = result.Error
		return
	}

	id = article.ID

	return
}

func Update(id uint, name, title, content string, category_first_id, category_second_id int) {
	article := article.Article{}

	article.ID = id

	if name != "" {
		article.Name = name
	}

	if title != "" {
		article.Title = title
	}

	if content != "" {
		article.Content = content
	}

	if category_first_id != 0 {
		article.CategoryFirstId = category_first_id
	}

	if category_second_id != 0 {
		article.CategorySecondId = category_second_id
	}

	orm.MysqlOrm.Updates(&article)

}

func Detail(id int) (article article.Article, err error) {

	// CategoryFirst 如果没有指定选择字段时，仍然是 select * from ，

	result := orm.MysqlOrm.Debug().Preload("CategoryFirst", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "name")
	}).Preload("CategorySecond").First(&article, id)

	if result.Error != nil {
		err = errors.New("未找到对应详情")
	}
	return
}

// 类目部分
func CategoryAdd(name string, level, parentId int) (category article.Category, err error) {

	var count int64

	orm.MysqlOrm.Model(&article.Category{}).Where("name", name).Count(&count)

	if count > 0 {
		err = errors.New("类目名称不能相同")
		return
	}

	category = article.Category{
		Name:     name,
		Level:    level,
		ParentId: parentId,
	}

	result := orm.MysqlOrm.Create(&category)

	if result.Error != nil {
		err = result.Error
		return
	}

	return
}

func CategoryUpdate(id int, name string) (category article.Category, err error) {
	result := orm.MysqlOrm.First(&category, id)

	if result.Error != nil {
		err = errors.New("未找到对应的记录")
		return
	}

	var count int64
	orm.MysqlOrm.Model(&article.Category{}).Where("id <>", id).Where("name", name).Count(&count)
	if result.Error != nil {
		err = errors.New("名称不能重复")
		return
	}

	// 更新
	orm.MysqlOrm.Model(&category).Update("name", name)

	return
}

func CategoryDetail(id int) (category article.Category, err error) {
	result := orm.MysqlOrm.Preload("Childrens").First(&category, id)

	if result.Error != nil {
		err = errors.New("未找到对应的记录")
		return
	}

	// orm.MysqlOrm.Debug().Preload("Childrens").First(&category, id)

	return
}

func CategoryList() (categorys []article.Category, err error) {

	result := orm.MysqlOrm.Preload("Childrens").Where("level = ?", 1).Find(&categorys)

	if result.Error != nil {
		err = errors.New("未找到对应的记录")
		return
	}

	// var ategorys1 []article.Category
	// orm.MysqlOrm.Preload("Childrens").Where("level = ?", 1).Find(&ategorys1)

	return
}
