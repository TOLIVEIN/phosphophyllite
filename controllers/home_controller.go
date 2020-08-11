package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//HomeGet ...
func HomeGet(context *gin.Context) {
	isLogin := GetSession(context)
	context.HTML(http.StatusOK, "home.html", gin.H{
		// "title":   "首页",
		"isLogin": isLogin,
	})
}
