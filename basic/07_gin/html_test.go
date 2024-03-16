package _7_gin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

func TestHTML(t *testing.T) {
	server := gin.Default()

	// 加载模板文件
	server.LoadHTMLFiles("index.html")
	server.GET("/index", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"title": "/index",
		})
	})

	server.Run(":8088")
}
