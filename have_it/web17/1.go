package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mattn/go-colorable"
	"net/http"
)

func main() {
	gin.ForceConsoleColor()
	// 替换默认Writer（关键步骤）
	gin.DefaultWriter = colorable.NewColorableStdout()

	r := gin.Default()
	r.GET("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "GET",
		})
	})

	r.POST("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "POST",
		})
	})

	r.DELETE("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "DELETE",
		})
	})

	r.PUT("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "PUT",
		})
	})
	//路由组,公共域名
	video:=r.Group("/video",nil)
	{ //视频首页和详情页
		video.GET("/video/index", func(c *gin.Context) {
			c.JSON(200, gin.H{"msg": "video/index"})
		})
	}


	//商城的首页和详情页
	r.GET("/shop/index", func(c *gin.Context) {
		c.JSON(200,gin.H{"msg":"shop/index"})
	})
//默认请求
r.NoRoute(func(c *gin.Context) {
	c.JSON(http.StatusOK,gin.H{"hhh":"www.baidu.com"})
})

	r.Run(":9090")
}
