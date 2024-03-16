package _7_gin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

func Test_Hello(t *testing.T) {
	g := gin.Default()

	g.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "ok",
		})
	})

	v1 := g.Group("/v1")
	{
		v1.GET("/user", func(ctx *gin.Context) {

		})
	}

	_ = g.Run(":8088")
}
