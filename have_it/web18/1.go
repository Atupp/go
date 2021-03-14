package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mattn/go-colorable"
	"time"
)

func m1(c *gin.Context) { //这个就是中间件,统计耗时
	//记录时间
	fmt.Println("m1 in")
	start := time.Now()
	c.Next() //调用后续函数
	cost := time.Since(start)
	fmt.Printf("cost:%v\n", cost)
	fmt.Println("m1 out")
	//c.Abort()//阻止调用后面的函数
}

func m2(c *gin.Context) {
	//记录时间
	fmt.Println("m2 in")
	//假如在并发中修改C的值,只能用C的拷贝,因为并发不安全
	//go XXX c.Copy()  只能用只读
	c.Set("name","AT")//这里输入了内容,相当于登录界面,后续要获取你输入的东西,跨中间件存值
	//c.Abort()//跳过后续函数
	fmt.Println("m2 out")
	//c.Abort()
}

func main() {
	gin.ForceConsoleColor()
	// 替换默认Writer（关键步骤）
	gin.DefaultWriter = colorable.NewColorableStdout()

	r := gin.Default()
	r.Use(m1,m2) //全局注册中间件

	r.GET("/index", func(c *gin.Context) {
		name,ok:=c.Get("name")//这是要取哪的KEY和VALUE哈
		if !ok {
			name="匿名用户"
		}
		c.JSON(200, gin.H{"55": name})
	})
	r.Run(":9000")
}
