package main

import (
	"book/dao"
	"book/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	err := dao.InitDB()
	if err != nil {
		return
	}
	r := gin.Default()
	r.LoadHTMLGlob("../templates/*")

	// /book/new页面显示
	r.GET("/book/new", func(c *gin.Context) {
		c.HTML(http.StatusOK, "new_book.html", gin.H{"Title": "新增图书"})
	})

	// 数据提交
	r.POST("/book/new", func(c *gin.Context) {
		book := model.NewBook()
		if err := c.Bind(book); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// 数据插入到数据库
		err := dao.Add(book.Title, book.Price)
		if err != nil {
			return
		}
		c.JSON(http.StatusOK, gin.H{"插入": "成功"})
	})

	// /book/list页面显示
	r.GET("/book/list", func(c *gin.Context) {
		// 获取数据库查询信息展示
		bookList, err := dao.QueryAllBook()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 1,
				"msg":  err,
			})
			return
		}
		// 返回数据
		c.HTML(http.StatusOK, "book_list.html", gin.H{
			"code": 0,
			"data": bookList,
		})
	})

	r.DELETE("/book", func(c *gin.Context) {
		// 获取数据库查询信息展示
		c.HTML(http.StatusOK, "book_list.html", gin.H{"ID": "123", "Title": "新增图书", "Price": "123"})
	})

	err = r.Run(":8080")
	if err != nil {
		return
	}
}
