package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type userinfo struct {
	Username string `form:username json:"username"`
	Password string `form:username json:"password"`
}

func main() {
	e := gin.Default()
	e.LoadHTMLFiles("./h.html") //载入模板
	e.GET("/user", func(c *gin.Context) {
		c.HTML(http.StatusOK, "h.html", nil) //要返回到这个模板,那么肯定先要载入模板
	})
	e.POST("/form", func(c *gin.Context) {
		var user userinfo
		err := c.ShouldBind(&user) //需要在输入中找见参数,传入user,修改其值(在其有值的前提下).取到值后赋予给相关
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("%#v\n", user)
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}
	})
	e.POST("/json", func(c *gin.Context) {
		//username:=c.Query("username")
		//password:=c.Query("password")
		//user:=userinfo{
		//	username,
		//	password,
		//}
		var user userinfo
		err := c.ShouldBind(&user) //需要在输入中找见参数,传入user,修改其值(在其有值的前提下).取到值后赋予给相关
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("%#v\n", user)
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}
	})
	e.Run(":9090")
}
