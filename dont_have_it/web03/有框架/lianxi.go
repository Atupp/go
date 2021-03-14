package main

import "github.com/gin-gonic/gin"
func main() {
r:=gin.Default()
r.GET("/UPP", func(c *gin.Context) {
	c.JSON(200, gin.H{
		"001": 002,
	})
})
r.Run(":9090")
}
