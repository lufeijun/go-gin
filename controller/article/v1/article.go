package v1

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"

	response "gin/structs"
	articleUltility "gin/ultility/article/v1"
)

// 列表
func List(c *gin.Context) {
	id, err := strconv.Atoi(c.PostForm("id"))

	if err != nil {
		id = 0
	}

	// page
	page, err := strconv.ParseInt(c.PostForm("page"), 10, 64)
	if err != nil {
		page = 1
	}

	// pagesize
	pageSize, err := strconv.ParseInt(c.PostForm("page_size"), 10, 64)
	if err != nil {
		pageSize = 10
	}

	name := c.PostForm("name")

	list := articleUltility.ArticleList(id, name, page, pageSize)

	c.JSON(200, list)
}

// 添加
func Add(c *gin.Context) {
	res := response.Response{}

	name := c.PostForm("name")
	title := c.PostForm("title")
	content := c.PostForm("content")

	if name == "" {
		res.Status = 1
		res.Message = "name 字段不能为空"
		c.JSON(200, res)
		return
	}

	if title == "" {
		res.Status = 1
		res.Message = "title 字段不能为空"
		c.JSON(200, res)
		return
	}

	id, err := articleUltility.ArticleAdd(name, title, content)

	if err != nil {
		res.Status = 1
		res.Message = "插入失败：" + err.Error()
		c.JSON(200, res)
		return
	}

	res.Status = 0
	res.Message = "success"
	res.Data = id

	c.JSON(200, res)

}

// 更新
func Update(c *gin.Context) {
	res := response.Response{}
	id, err := strconv.Atoi(c.PostForm("id"))

	if err != nil {
		res.Status = 1
		res.Message = "id 字段出错"
		c.JSON(200, res)
		return
	}

	name := c.PostForm("name")
	title := c.PostForm("title")
	content := c.PostForm("content")

	// if name == "" {
	// 	res.Status = 1
	// 	res.Message = "name 字段不能为空"
	// 	c.JSON(200, res)
	// 	return
	// }

	// if title == "" {
	// 	res.Status = 1
	// 	res.Message = "title 字段不能为空"
	// 	c.JSON(200, res)
	// 	return
	// }

	articleUltility.ArticleUpdate(uint(id), name, title, content)

	res.Status = 1
	res.Message = "更新成功"

	c.JSON(200, res)

}

func Detail(c *gin.Context) {

	id, err := strconv.Atoi(c.Query("id"))

	res := response.Response{}

	if err != nil {
		res.Status = 1
		res.Message = "id 字段出错"
		c.JSON(200, res)
		return
	}

	article := articleUltility.ArticleDetail(id)

	// a := article.CreatedAt.Format("2006-01-02 15:04:05")

	fmt.Println(article.CreatedAt.String())

	res1 := response.ToClientData(0, "success", article)

	c.JSON(200, res1)
}
