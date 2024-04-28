package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitializeRouter(router *gin.Engine) {
	// Definition
	version1 := router.Group("/api/version1")
	{
		version1.GET("/GET", func(structure *gin.Context) {
			structure.JSON(http.StatusOK, gin.H{
				"message": "Success GET",
			})
		})
		version1.POST("/POST", func(structure *gin.Context) {
			structure.JSON(http.StatusOK, gin.H{
				"message": "Success POST",
			})
		})
		version1.DELETE("/DELETE", func(structure *gin.Context) {
			structure.JSON(http.StatusOK, gin.H{
				"message": "Success DELETE",
			})
		})
		version1.PUT("/PUT", func(structure *gin.Context) {
			structure.JSON(http.StatusOK, gin.H{
				"message": "Success PUT",
			})
		})
	}
}
