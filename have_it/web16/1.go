package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/mattn/go-colorable"
)

func main() {
	gin.ForceConsoleColor()
	// 替换默认Writer（关键步骤）
	gin.DefaultWriter = colorable.NewColorableStdout()



	r:=gin.Default()
	r.GET("/index", func(c *gin.Context) {
		//c.JSON(http.StatusOK,gin.H{
		//	"status":"ok",
		//})
		c.Redirect(http.StatusMovedPermanently,"http://www.baidu.com")
	})
	r.GET("/a", func(c *gin.Context) {
		//访问A,跳转到B的路由
		c.Request.URL.Path="/b"//请求的URL的路径
		r.HandleContext(c)//继续后续处理
	})
	r.GET("/b", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"message":"b",
		})
	})
	r.Run(":9090")
}
