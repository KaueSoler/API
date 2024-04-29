package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateApiHandler(structure *gin.Context) {
	structure.JSON(http.StatusOK, gin.H{
		"message": "Success PUT",
	})

}
