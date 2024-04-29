package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowApiHandler(structure *gin.Context) {
	structure.JSON(http.StatusOK, gin.H{
		"message": "Success GET",
	})

}

func CreateApiHandler(structure *gin.Context) {
	structure.JSON(http.StatusOK, gin.H{
		"message": "Success POST",
	})

}

func DeleteApiHandler(structure *gin.Context) {
	structure.JSON(http.StatusOK, gin.H{
		"message": "Success DELETE",
	})

}

func UpdateApiHandler(structure *gin.Context) {
	structure.JSON(http.StatusOK, gin.H{
		"message": "Success PUT",
	})

}
