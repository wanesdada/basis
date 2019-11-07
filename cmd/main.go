package main

import(
	"github.com/gin-gonic/gin"
)
func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context){
		c.JSON(200,gin.H{
			"ceshi":"hello world",
		})
	})
	r.Run(":8088")   // 强指定端口，默认8088
}
