package router

import "github.com/gin-gonic/gin"

func Setup() *gin.Engine {
	e := gin.Default()

	e.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	return e
}
