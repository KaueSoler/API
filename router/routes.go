package router

import (
	"github.com/KaueSoler/api/handler"
	"github.com/gin-gonic/gin"
)

func InitializeRouter(router *gin.Engine) {
	// Definition
	version1 := router.Group("/api/version1")
	{
		version1.GET("/GET", handler.ShowApiHandler)

		version1.POST("/POST", handler.CreateApiHandler)

		version1.DELETE("/DELETE", handler.DeleteApiHandler)

		version1.PUT("/PUT", handler.UpdateApiHandler)
	}
}
