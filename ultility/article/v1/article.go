package v1

import (
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

	// article.UpdatedAt = my.MyTimeInit()

	orm.GormDB.Updates(&article)

}

func ArticleDetail(id int) (article articleModel.Article) {
	orm.GormDB.First(&article, id)
	return
}
