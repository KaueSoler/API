package router

import "github.com/gin-gonic/gin"

func Inicialize() {
	// inicialização do router utilizando configurações padrão do gin-gonic
	router := gin.Default()

	// deifnição de uma rota
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// Api rodando, teste em postman http://localhost:8080/ping
	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
