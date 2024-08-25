package main

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//加载静态文件
	r.Static("/xxx", "./statics")
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	// r.LoadHTMLFiles("./posts/about.html",
	// 	"./posts/blog-single.html",
	// 	"./posts/blog.html",
	// 	"./posts/contact.html",
	// 	"./posts/home.html",
	// 	"./posts/portfolio.html",
	// 	"./posts/services.html",
	// 	"./posts/team.html")

	r.LoadHTMLGlob("./posts/*")

	r.GET("/posts/index", func(c *gin.Context) {
		//HTTP请求
		c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
			"title": "liwenzhou.com111111",
		})
	})

	r.GET("/users/index", func(c *gin.Context) {
		//HTTP请求
		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
			"title1": "<a href='https://movie.douban.com/'>豆瓣电影</a>",
			"title2": "<a href='https://movie.douban.com/tv/'>豆瓣电视</a>",
		})
	})

	//返回从网上下载的模板
	r.GET("/index.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.GET("/services.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "services.html", nil)

	})
	r.GET("/about.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "about.html", nil)

	})
	r.GET("/blog-single.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "blog-single.html", nil)

	})
	r.GET("/blog.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "blog.html", nil)

	})
	r.GET("/contact.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "contact.html", nil)

	})
	r.GET("/portfolio.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "portfolio.html", nil)

	})
	r.GET("/team.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "team.html", nil)

	})

	r.Run(":9060")

}
