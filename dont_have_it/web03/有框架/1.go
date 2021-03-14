package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
func sayhello(c *gin.Context){
c.JSON(200,gin.H{
	"message":"hello golang",
})
}


func main (){
r:=gin.Default()//返回的是个引擎,用于管理访问什么地址\执行什么函数,类似管理员
r.GET("/hello",sayhello)//RouterGroup结构体元素之是*Engine,那么大的结构体的方法,其元素也可以调用
r.PUT("/hello", func(c *gin.Context) {
	c.JSON(200,"hhhhhh")
})
r.POST("/hello", func(c *gin.Context) {
	c.JSON(200,gin.H{
		"111":"888",
	})
})
r.DELETE("/hello", func(c *gin.Context) {
	c.JSON(http.StatusOK,gin.H{
		"111":"777",
	})
})
//启动服务
r.Run(":9090")
}