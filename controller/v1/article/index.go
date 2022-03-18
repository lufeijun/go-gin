package article

import (
	"gin/structs/response"
	"gin/ultility/v1/article"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 文章部分

// 文章列表
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

	res.Data = article.List(id, name, page, pageSize)

	c.JSON(200, res)
}

// 详情
func Detail(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	res := response.GetResponse()

	if err != nil {
		res.SetMessage("id 字段出错")
		c.JSON(200, res)
		return
	}

	article, err := article.Detail(id)

	if err != nil {
		res.SetMessage("未找到")
	} else {
		res.Data = article
	}

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

	category_first_id, err := strconv.Atoi(c.PostForm("category_first_id"))
	if err != nil {
		res.SetMessage("category_first_id 字段有误")
		c.JSON(200, res)
		return
	}

	category_second_id, err := strconv.Atoi(c.PostForm("category_second_id"))
	if err != nil {
		res.SetMessage("category_second_id 字段有误")
		c.JSON(200, res)
		return
	}

	id, err := article.Add(name, title, content, category_first_id, category_second_id)

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
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		res.SetMessage("id 字段出错")
		c.JSON(200, res)
		return
	}

	name := c.PostForm("name")
	title := c.PostForm("title")
	content := c.PostForm("content")

	category_first_id, err := strconv.Atoi(c.PostForm("category_first_id"))
	if err != nil {
		category_first_id = 0
	}

	category_second_id, err := strconv.Atoi(c.PostForm("category_second_id"))
	if err != nil {
		category_second_id = 0
	}

	article.Update(uint(id), name, title, content, category_first_id, category_second_id)

	c.JSON(200, res)
	return
}

// 文章类目

// 列表
func CategoryList(c *gin.Context) {
	res := response.GetResponse()

	category, err := article.CategoryList()

	if err != nil {
		res.SetMessage(err.Error())
	} else {
		res.SetData(category)
	}

	c.JSON(http.StatusOK, res)
	return
}

// 添加
func CategoryAdd(c *gin.Context) {
	res := response.GetResponse()

	name := c.PostForm("name")

	if name == "" {
		res.SetMessage("name 不能为空")
		c.JSON(http.StatusOK, res)
		return
	}

	level, err := strconv.Atoi(c.PostForm("level"))
	if err != nil {
		level = 1
	}

	parentId, err := strconv.Atoi(c.PostForm("parent_id"))
	if err != nil {
		parentId = 0
	}

	if level == 2 && parentId == 0 {
		res.SetMessage("parent_id 参数有误")
		c.JSON(http.StatusOK, res)
		return
	}

	category, err := article.CategoryAdd(name, level, parentId)

	if err != nil {
		res.SetMessage(err.Error())
		c.JSON(http.StatusOK, res)
		return
	}

	res.SetData(category)
	c.JSON(http.StatusOK, res)
	return
}

// 更新
func CategoryUpdate(c *gin.Context) {
	res := response.GetResponse()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		res.SetMessage("id 参数有误")
		c.JSON(http.StatusOK, res)
		return
	}

	name := c.PostForm("name")

	category, err := article.CategoryUpdate(id, name)

	if err != nil {
		res.SetMessage(err.Error())
	} else {
		res.SetData(category)
	}

	c.JSON(http.StatusOK, res)
	return

}

// 详情
func CategoryDetail(c *gin.Context) {
	res := response.GetResponse()

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		res.SetMessage("id 参数有误")
		c.JSON(http.StatusOK, res)
		return
	}

	category, err := article.CategoryDetail(id)

	if err != nil {
		res.SetMessage(err.Error())
	} else {
		res.SetData(category)
	}

	c.JSON(http.StatusOK, res)
	return

}
