package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 创建一个默认的 Gin 路由器
	r := gin.Default()

	// 添加中间件（可选）
	r.Use(gin.Logger())   // 日志中间件
	r.Use(gin.Recovery()) // 恢复中间件

	// 定义路由和处理函数
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Gin HTTP Server",
		})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// 带参数的路由
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello " + name,
		})
	})

	// POST 请求示例
	r.POST("/user", func(c *gin.Context) {
		var user struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		}

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"status": "user created",
			"name":   user.Name,
			"email":  user.Email,
		})
	})

	// 分组路由
	api := r.Group("/api")
	{
		api.GET("/version", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"version": "1.0.0",
			})
		})
	}

	// 静态文件服务
	r.Static("/static", "./static")

	// 启动服务器，默认监听 0.0.0.0:8080
	r.Run() // 也可以指定端口 r.Run(":3000")
}
