package router

import "github.com/gin-gonic/gin"

func InitializeRouter(router *gin.Engine) {
	// Definition
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
