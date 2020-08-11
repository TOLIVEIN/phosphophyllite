package routers

import (
	"phosphophyllite/controllers"

	"github.com/gin-gonic/gin"
)

// InitRouter ...
func InitRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("views/*")

	router.GET("/register", controllers.RegisterGet)
	router.POST("/register", controllers.RegisterPost)

	router.GET("/login", controllers.LoginGet)

	return router
}
