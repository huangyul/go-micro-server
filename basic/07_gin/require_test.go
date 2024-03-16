package _7_gin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

type Account struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password"`
}

func TestRequire(t *testing.T) {
	g := gin.Default()
	g.POST("/register", func(ctx *gin.Context) {
		var account Account
		if err := ctx.ShouldBindJSON(&account); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
			return
		}
		ctx.String(http.StatusOK, "登录成功")
	})

	_ = g.Run(":8088")
}
