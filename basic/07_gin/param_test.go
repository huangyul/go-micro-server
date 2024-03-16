package _7_gin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

func Test_params(t *testing.T) {
	g := gin.Default()
	// 静态路由
	g.GET("/user", func(ctx *gin.Context) {
		name := ctx.Query("name")
		ctx.String(http.StatusOK, "name: %s", name)
	})
	// 参数路由
	g.GET("/user/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		ctx.String(http.StatusOK, fmt.Sprintf(`id:%s`, id))
	})
	// 获取form上的参数
	g.POST("/user1", func(ctx *gin.Context) {
		person := ctx.PostForm("person")
		ctx.String(http.StatusOK, "persion: %s", person)
	})

	_ = g.Run(":8088")
}
