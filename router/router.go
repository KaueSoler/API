package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Initialize router
	router := gin.Default()
	InitializeRouter(router)
	// Api rodando, teste em postman http://localhost:8080/ping
	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
