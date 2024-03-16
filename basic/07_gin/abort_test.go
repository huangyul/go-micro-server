package _7_gin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"testing"
)

func TestAbort(t *testing.T) {
	server := gin.Default()

	server.GET("/user", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello, world")
	})

	go func() {
		_ = server.Run(":8088")
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("server关闭中")
}
