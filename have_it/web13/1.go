package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
func get(c *gin.Context){
	//获取路径参数
	name:=c.Param("name")//取到了name和age对应的输入
	age:=c.Param("age")
	//内容
	//格式
	c.JSON(http.StatusOK,gin.H{
		"name":name,//输入赋予这边//内容的拓展是模板和html
		"age":age,
	})
	}
func main() {
	//调用引擎
e:=gin.Default()
//请求+发送
e.GET("/:name/:age",get)//URL参数,路径参数,注意路由匹配时候是否会冲突
//启动
e.Run(":9090")
}
