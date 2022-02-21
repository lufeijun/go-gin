package v1

import (
	"errors"
	"fmt"
	articleModel "gin/models/article/v1"
	structs "gin/structs"

	orm "gin/database"
)

func ArticleList(id int, name string, page int64, pagesize int64) (pagestruct structs.PageStruct) {
	var articles []articleModel.Article
	sql := orm.GormDB

	// id 查询
	if id != 0 {
		sql = sql.Where("id", id)
	}

	if name != "" {
		sql = sql.Where("name like ?", "%"+name+"%")
	}

	var total int64
	sql.Model(&articleModel.Article{}).Count(&total)
	pagestruct.LastPage = total/pagesize + 1

	sql.Scopes(orm.Paginate(page, pagesize)).Find(&articles)

	// 赋值
	pagestruct.Total = total
	pagestruct.Page = page
	pagestruct.Size = pagesize
	pagestruct.Data = articles

	return
}

func ArticleAdd(name, title, content string) (id uint, err error) {
	article := articleModel.Article{
		Name:    name,
		Title:   title,
		Content: content,
	}

	result := orm.GormDB.Create(&article)

	if result.Error != nil {
		err = result.Error
		return
	}

	id = article.ID

	return
}

func ArticleUpdate(id uint, name, title, content string) {
	article := articleModel.Article{}

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

	orm.GormDB.Updates(&article)

}

func ArticleDetail(id int) (article articleModel.Article, err error) {
	result := orm.GormDB.First(&article, id)

	if result.Error != nil {
		err = errors.New("未找到对应详情")
	}
	return
}

// 类目部分
func CategoryAdd(name string, level, parentId int) (category articleModel.Category, err error) {
	category = articleModel.Category{
		Name:     name,
		Level:    level,
		ParentId: parentId,
	}

	result := orm.GormDB.Create(&category)

	if result.Error != nil {
		err = result.Error
		return
	}

	return
}

func CategoryUpdate(id int, name string) (category articleModel.Category, err error) {
	result := orm.GormDB.First(&category, id)

	if result.Error != nil {
		err = errors.New("未找到对应的记录")
		return
	}

	// 更新
	orm.GormDB.Model(&category).Update("name", name)

	return
}

func CategoryDetail(id int) (category articleModel.Category, err error) {
	result := orm.GormDB.Preload("Childrens").First(&category, id)

	if result.Error != nil {
		err = errors.New("未找到对应的记录")
		return
	}

	orm.GormDB.Debug().Preload("Childrens").First(&category, id)

	return
}

func CategoryList() (categorys []articleModel.Category, err error) {

	result := orm.GormDB.Preload("Childrens").Where("level = ?", 1).Find(&categorys)

	if result.Error != nil {
		err = errors.New("未找到对应的记录")
		return
	}

	var ategorys1 []articleModel.Category
	orm.GormDB.Debug().Preload("Childrens", "id < ?", 3).Where("level = ?", 1).Find(&ategorys1)

	return
}

func CategoryTest() {
	// test1()

	test2()
}

// 自动创建，更新关联信息
func test1() {
	category := articleModel.Category{
		Name:     "一级类目-test",
		Level:    1,
		ParentId: 0,
		Childrens: []articleModel.CategoryChildren{
			{Name: "二级类目-test-1", Level: 2},
			{Name: "二级类目-test-2", Level: 2},
			{Name: "二级类目-test-3", Level: 2},
		},
	}

	result := orm.GormDB.Create(&category)

	fmt.Println(result.Error)
}

func test2() {
	var category articleModel.Category
	// var categorys []articleModel.Category
	// result := orm.GormDB.Model(&category).Association("CategoryChildren").Find(&category)
	// orm.GormDB.Debug().Association("CategoryChildren").Find(&category)
	// orm.GormDB.Debug().Find(&category)

	// result := orm.GormDB.Model(&category).Association("Childrens")

	// if result.Error != nil {
	// 	panic(result.Error)
	// }

	var categoryChildren articleModel.CategoryChildren

	orm.GormDB.Debug().Model(&category).Association("Childrens").Find(&categoryChildren)

	// .Find(&categoryChildren)
	fmt.Println(categoryChildren)

}
