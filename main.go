package main

import (
	"phosphophyllite/database"
	"phosphophyllite/routers"
)

func main() {
	// router := gin.Default()

	// router.LoadHTMLGlob("templates/*")

	// router.GET("/admin", func(context *gin.Context) {
	// 	context.HTML(http.StatusOK, "admin.html", gin.H{
	// 		"title": "Admin Page...",
	// 	})
	// })

	// router.GET("/", func(context *gin.Context) {
	// 	context.String(http.StatusOK, "hello gin")
	// })

	// router.GET("/ping", func(context *gin.Context) {
	// 	context.JSON(200, gin.H{
	// 		"message": "pong...",
	// 	})
	// })

	database.InitMysql()
	router := routers.InitRouter()

	router.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
