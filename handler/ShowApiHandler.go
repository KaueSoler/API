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
