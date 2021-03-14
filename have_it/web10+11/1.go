package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type msg struct {
	Name    string `json:"name"`
	Message string
	Age     int
}

func main() {
	//调用引擎
	e := gin.Default()
	//客户端发送+服务器返回
	e.GET("/web", func(c *gin.Context) {
		//获取客户端输入参数方案1
		//name := c.Query("upp") //输入方式是./web?关键词=用户输入
		//获取客户端输入参数方案2
		//name := c.DefaultQuery("fuck", "你妈逼")
		//获取客户端输入参数方案3
		name,ok:=c.GetQuery("fuck")
		if !ok {
			name="hkhjk"
		}
		//返回内容
		//方法1:使用map
		//data:= gin.H{"name": "小王子", "age": 18, "message": "hello"}
		//方法2:结构体,直接输入内容,对应相对位置
		data := msg{
			name,
			"hello",
			18,
		}
		//返回形式
		c.JSON(http.StatusOK, data)
	})

	//启动引擎
	e.Run(":9000")
}
