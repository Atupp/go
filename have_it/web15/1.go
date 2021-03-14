package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"path"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./index.html")
	r.POST("/upload", func(c *gin.Context) {
		//从请求中读取文件
		file,err:=c.FormFile("f1")
		if err!=nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
			return
		}
		//将读取的文件存本地
		//dst:=fmt.Sprintf("./%s",file.Filename)
		log.Println(file.Filename)
		dst:=path.Join("./",file.Filename)
		c.SaveUploadedFile(file,dst)
		c.JSON(200,gin.H{
			"status":"ok",
		})

	})
	r.GET("/index", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})
	r.Run(":9070")
}
//点了上传,没反应