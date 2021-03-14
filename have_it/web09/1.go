package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)
//静态文件处理 .css  js文件  图片

func index(w *gin.Context) {
	//模板渲染
	w.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "liwenzhou.com",
	})
}

func users(w *gin.Context) {
	w.HTML(http.StatusOK, "666", gin.H{
		"title": "<a href='https://liwenzhou.com'>李文周的博客</a>",
	})
}
func safe(str string)template.HTML{
	return template.HTML(str)
}
func main() {
	e := gin.Default()
	//添加自定义函数
    a:=template.FuncMap{
    	"safe":safe,
	}
	e.SetFuncMap(a)
    //加载静态文件
    e.Static("/xxx","./statics")//XXX代表的是路径缩写,是./statics到最终文件路径的缩写.
	//模板解析
	e.LoadHTMLGlob("./templates/**/*")//**表示目录,*表示文件
	//e.LoadHTMLFiles("./templates/index/users.tmpl","./templates/users/home.html")
	//模板渲染

	e.GET("/users",users)
	e.GET("/index", index)
    //返回从网上下载的模板
    e.GET("/home",func(c *gin.Context){
    	c.HTML(http.StatusOK,"home.html",nil)
	})
	//启动server
	e.Run(":9090")
}