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

	return router
}
