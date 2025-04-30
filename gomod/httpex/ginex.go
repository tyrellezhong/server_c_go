package httpex

import (
	"github.com/gin-gonic/gin"
)

func GinServer() {
	// 创建一个默认的 Gin 路由器
	r := gin.Default()

	// 定义一个简单的 GET 路由
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// 启动 HTTP 服务器
	r.Run() // 默认监听在 0.0.0.0:8080
}
