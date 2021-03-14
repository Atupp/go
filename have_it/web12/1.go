package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func get(c *gin.Context) { //因为要来回修改其内容,所以用地址
	//返回内容.模板
	content := "login.html"
	//返回内容.数据---空
	//返回形式
	c.HTML(http.StatusOK, content, nil)
}

func post(c *gin.Context) {
	//方法1
	//username:=c.PostForm("username")//取的是login.HTML中name="XXX"
	//password:=c.PostForm("password")
	//方法2
	//username:=c.DefaultPostForm("username","sb")
	//password:=c.DefaultPostForm("password","6666")
	//方法3
	username,ok:=c.GetPostForm("username")
		if !ok {
			username="sb"
		}
	password,_:=c.GetPostForm("password")
		if !ok {
			password="123"
		}
	//方法3

	//返回内容.模板
	con:="post.html"
	//返回内容.数据
	data:=gin.H{
		"Name":username,
		"Password":password,
	}
	//返回形式
	c.HTML(http.StatusOK,con,data)
}
func main() {
	//调用引擎
	e := gin.Default()
	//加载模板文件
	e.LoadHTMLGlob("templates/*")
	//请求+发送
	e.GET("/get", get)
	e.POST("/get", post)
	//启动引擎
	e.Run(":9090")
}
