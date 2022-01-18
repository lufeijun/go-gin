package v1

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"gin/structs/response"
	articleUltility "gin/ultility/article/v1"
)

// 列表
func List(c *gin.Context) {
	res := response.GetResponse()
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

	res.Data = articleUltility.ArticleList(id, name, page, pageSize)

	c.JSON(200, res)
}

// 添加
func Add(c *gin.Context) {
	res := response.GetResponse()

	name := c.PostForm("name")
	title := c.PostForm("title")
	content := c.PostForm("content")

	if name == "" {
		res.SetMessage("name 字段不能为空")
		c.JSON(200, res)
		return
	}

	if title == "" {
		res.SetMessage("title 字段不能为空")
		c.JSON(200, res)
		return
	}

	id, err := articleUltility.ArticleAdd(name, title, content)

	if err != nil {
		res.SetMessage("插入失败：" + err.Error())
		c.JSON(200, res)
		return
	}

	res.Data = id
	c.JSON(200, res)

}

// 更新
func Update(c *gin.Context) {
	res := response.GetResponse()
	id, err := strconv.Atoi(c.PostForm("id"))

	if err != nil {
		res.SetMessage("id 字段出错")
		c.JSON(200, res)
		return
	}

	name := c.PostForm("name")
	title := c.PostForm("title")
	content := c.PostForm("content")

	articleUltility.ArticleUpdate(uint(id), name, title, content)

	c.JSON(200, res)

}

func Detail(c *gin.Context) {

	id, err := strconv.Atoi(c.PostForm("id"))

	res := response.GetResponse()

	if err != nil {
		res.SetMessage("id 字段出错")
		c.JSON(200, res)
		return
	}

	article, err := articleUltility.ArticleDetail(id)

	if err != nil {
		res.SetMessage("未找到")
	} else {
		res.Data = article
	}

	c.JSON(200, res)
}
